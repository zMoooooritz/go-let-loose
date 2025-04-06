package socket

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/config"
)

const (
	BUF_SIZE = 32768

	KEY_LENGTH = 4

	MAX_ITERATION_COUNT = 100

	RESP_SUCCESS = "SUCCESS"
	RESP_FAIL    = "FAIL"
	RESP_EMPTY   = "EMPTY"

	INIT_TIMEOUT        = time.Duration(5 * time.Second)
	WRITE_TIMEOUT       = time.Duration(3 * time.Second)
	SLOW_READ_TIMEOUT   = time.Duration(2 * time.Second)
	NORMAL_READ_TIMEOUT = time.Duration(1 * time.Second)
	FAST_READ_TIMEOUT   = time.Duration(300 * time.Millisecond)
)

type ServerConnection struct {
	ip       string
	port     string
	password string
	conn     net.Conn
	key      []byte

	fastUnsafeLogsFetching bool
}

var ErrInvalidRconCommand = errors.New("invalid rcon command")

func NewConnection(ip, port, password string) (*ServerConnection, error) {
	sc := ServerConnection{ip: ip, port: port, password: password}
	err := sc.connect()
	return &sc, err
}

func (sc *ServerConnection) Reconnect() error {
	sc.Close()
	return sc.connect()
}

func (sc *ServerConnection) Close() {
	if sc.IsActive() {
		_ = sc.conn.Close()
	}
}

func (sc *ServerConnection) Execute(command string, format config.ResponseFormat) ([]string, error) {
	sc.setTimeout(WRITE_TIMEOUT)
	err := sc.write(command)
	sc.resetTimeout()
	if err != nil {
		return []string{}, err
	}

	var data string
	sc.setTimeout(SLOW_READ_TIMEOUT)
	resp, err := sc.read()
	sc.resetTimeout()
	if err != nil {
		return []string{}, err
	}
	data = string(resp)
	if isInvalidCommand(command, data) {
		return []string{}, ErrInvalidRconCommand
	}
	if data == RESP_EMPTY || data == RESP_SUCCESS {
		return []string{}, nil
	}

	timeout := FAST_READ_TIMEOUT
	switch format {
	case config.RF_INDEXEDLIST:
		timeout = NORMAL_READ_TIMEOUT
	case config.RF_UNINDEXEDLIST:
		timeout = SLOW_READ_TIMEOUT
	}
	iterations := 0
	for {
		if format == config.RF_INDEXEDLIST {
			complete, err := isListComplete(data)
			if err != nil {
				return []string{}, err
			}
			if complete {
				break
			}
		}
		if format == config.RF_UNINDEXEDLIST && sc.fastUnsafeLogsFetching && isLastLineHeuristic(data) { // WARN: pretty dangerous, since the data will be left on the socket and be read by the next operation (handling this is realised in the logs fetcher worker since this is only relevant for the showlog command)
			break
		}

		iterations++
		sc.setTimeout(timeout)
		resp, err := sc.read()
		sc.resetTimeout()
		if errors.Is(err, os.ErrDeadlineExceeded) {
			break
		} else if err != nil {
			return []string{}, err
		}
		data += string(resp)
	}
	var lines []string
	if format == config.RF_UNINDEXEDLIST {
		data = strings.Trim(data, config.NEWLINE)

		// make it so every log entry is in one line
		timestampPattern := regexp.MustCompile(`^\[.+? \(\d+\)\]`)
		for _, line := range strings.Split(data, config.NEWLINE) {
			if timestampPattern.MatchString(line) {
				lines = append(lines, line)
			} else {
				line = strings.ReplaceAll(line, config.NEWLINE, config.ESCAPED_NEWLINE)
				if len(lines) == 0 {
					continue
				}
				lines[len(lines)-1] += line
			}
		}
	} else if format == config.RF_INDEXEDLIST {
		lines = strings.Split(data, config.LIST_DELIMITER)
		lines = lines[1 : len(lines)-1]
	} else {
		lines = append(lines, strings.TrimSuffix(data, config.NEWLINE))
	}
	return lines, nil
}

func isInvalidCommand(command, resp string) bool {
	if resp == RESP_FAIL {
		return true
	}
	if resp == "Cannot execute command for this gamemode." {
		return true
	}

	return false
}

func (sc *ServerConnection) EnableFastUnsafeLogsFetching() {
	sc.fastUnsafeLogsFetching = true
}

func isLastLineHeuristic(data string) bool {
	if len(data) == 0 {
		return false
	}
	if data[len(data)-1] == []byte(config.NEWLINE)[0] {
		lines := strings.Split(data, config.NEWLINE)
		if len(lines) < 2 {
			return false
		}
		lastFilledLine := lines[len(lines)-2]
		r := regexp.MustCompile(`^\[.+(ms|sec).+\]`)
		return r.Match([]byte(lastFilledLine))
	}
	return false
}

func (sc *ServerConnection) IsActive() bool {
	return sc.conn != nil
}

func (sc *ServerConnection) ExecuteNormalCommand(command string) (string, error) {
	resp, err := sc.Execute(command, config.RF_DIRECT)
	if err != nil {
		return "", err
	}
	return resp[0], nil
}

func (sc *ServerConnection) ExecuteListCommand(command string) ([]string, error) {
	return sc.Execute(command, config.RF_INDEXEDLIST)
}

func (sc *ServerConnection) connect() error {
	err := sc.initConnection()
	if err != nil {
		return err
	}
	return sc.authenticate()
}

func (sc *ServerConnection) initConnection() error {
	address := sc.ip + ":" + sc.port
	conn, err := net.DialTimeout("tcp", address, INIT_TIMEOUT)
	sc.conn = conn
	if err != nil {
		return err
	}
	if !sc.IsActive() {
		return errors.New("connection not active")
	}

	buf := make([]byte, KEY_LENGTH)
	sc.setTimeout(INIT_TIMEOUT)
	nbytes, err := sc.conn.Read(buf)
	sc.resetTimeout()
	if err != nil || nbytes != KEY_LENGTH {
		sc.conn = nil
		return err
	}
	sc.key = buf[:nbytes]

	return nil
}

func (sc *ServerConnection) authenticate() error {
	sc.setTimeout(WRITE_TIMEOUT)
	err := sc.write("login " + sc.password)
	sc.resetTimeout()
	if err != nil {
		return err
	}
	sc.setTimeout(SLOW_READ_TIMEOUT)
	resp, err := sc.read()
	sc.resetTimeout()
	if err != nil || string(resp) != RESP_SUCCESS {
		return errors.New("login failed")
	}
	return nil
}

func (sc *ServerConnection) write(msg string) error {
	if !sc.IsActive() {
		return errors.New("connection not active")
	}

	data := []byte(msg)
	data = sc.xor(data)
	n, err := sc.conn.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return errors.New("not all data send")
	}
	return err
}

func (sc *ServerConnection) read() ([]byte, error) {
	var resp []byte
	if !sc.IsActive() {
		return resp, errors.New("connection not active")
	}

	buf := make([]byte, BUF_SIZE)
	for {
		n, err := sc.conn.Read(buf)
		if err != nil {
			return resp, err
		}
		if n == 0 {
			break
		}

		dec := sc.xor(buf[:n])
		resp = append(resp, dec...)

		if n < BUF_SIZE {
			break
		}
	}

	return resp, nil
}

func (sc *ServerConnection) setTimeout(timeout time.Duration) {
	if sc.IsActive() {
		_ = sc.conn.SetDeadline(time.Now().Add(timeout))
	}
}

func (sc *ServerConnection) resetTimeout() {
	if sc.IsActive() {
		_ = sc.conn.SetDeadline(time.Time{})
	}
}

func isListComplete(data string) (bool, error) {
	split := strings.Split(data, config.LIST_DELIMITER)
	listLength, err := strconv.Atoi(split[0])
	if err != nil {
		return false, err
	}
	if listLength == 0 {
		return true, nil
	}
	if len(split) <= 2 {
		txt := fmt.Sprintf("invalid data %s", data)
		return false, errors.New(txt)
	}
	split = split[1 : len(split)-1]
	if listLength != len(split) {
		return false, nil
	}
	return true, nil
}

func (sc *ServerConnection) xor(data []byte) []byte {
	for idx := range data {
		data[idx] = data[idx] ^ sc.key[idx%KEY_LENGTH]
	}
	return data
}
