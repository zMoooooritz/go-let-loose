package socketv2

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"time"
)

const (
	LEGACY_V1_XOR_KEY_LENGTH = 4
	DO_READ_V1_XOR_KEY       = true

	INIT_TIMEOUT        = time.Duration(5 * time.Second)
	WRITE_TIMEOUT       = time.Duration(3 * time.Second)
	SLOW_READ_TIMEOUT   = time.Duration(2 * time.Second)
	NORMAL_READ_TIMEOUT = time.Duration(1 * time.Second)
	FAST_READ_TIMEOUT   = time.Duration(300 * time.Millisecond)
)

var (
	errConnectionNotActive = errors.New("connection not active")
)

type ServerConnection struct {
	ip        string
	port      string
	password  string
	version   int
	conn      net.Conn
	key       []byte
	authToken string
}

func NewConnection(ip, port, password string, version int) (*ServerConnection, error) {
	sc := ServerConnection{ip: ip, port: port, password: password, version: version, key: nil}
	err := sc.reconnect()
	return &sc, err
}

func (sc *ServerConnection) Execute(command, body string) (string, error) {
	sc.setTimeout(WRITE_TIMEOUT)
	request := NewRawRequest(sc.authToken, sc.version, command, body)
	err := sc.write(request.Pack())
	sc.resetTimeout()
	if err != nil {
		return "", err
	}

	sc.setTimeout(NORMAL_READ_TIMEOUT)
	resp, err := sc.read()
	sc.resetTimeout()
	if err != nil {
		return "", err
	}
	resData := RconResponse{}
	err = json.Unmarshal(resp, &resData)
	if err != nil {
		return "", err
	}
	if resData.StatusCode != StatusOk {
		return "", NewRconError(resData.StatusCode, resData.StatusMessage)
	}
	return string(resData.ContentBody), err
}

func (sc *ServerConnection) Reconnect() error {
	sc.Close()
	return sc.reconnect()
}

func (sc *ServerConnection) Close() {
	if sc.IsActive() {
		sc.conn.Close()
	}
}

func (sc *ServerConnection) IsActive() bool {
	return sc.conn != nil
}

func (sc *ServerConnection) reconnect() error {
	err := sc.initialize()
	if err != nil {
		return err
	}
	err = sc.connect()
	if err != nil {
		return err
	}
	return sc.login()
}

func (sc *ServerConnection) initialize() error {
	address := sc.ip + ":" + sc.port
	conn, err := net.DialTimeout("tcp", address, INIT_TIMEOUT)
	sc.conn = conn
	if err != nil {
		return err
	}

	// RCONv1 XOR key, can be ignored
	if DO_READ_V1_XOR_KEY {
		buf := make([]byte, LEGACY_V1_XOR_KEY_LENGTH)
		sc.setTimeout(INIT_TIMEOUT)
		nbytes, err := sc.conn.Read(buf)
		if err != nil || nbytes != LEGACY_V1_XOR_KEY_LENGTH {
			sc.conn = nil
			return err
		}
	}

	return nil
}

func (sc *ServerConnection) connect() error {
	resp, err := sc.Execute("ServerConnect", "")
	if err != nil {
		return err
	}
	sc.key, err = base64.StdEncoding.DecodeString(resp)
	return err
}

func (sc *ServerConnection) login() error {
	resp, err := sc.Execute("Login", sc.password)
	if err != nil {
		return err
	}
	sc.authToken = resp
	return err
}

func (sc *ServerConnection) write(data []byte) error {
	if !sc.IsActive() {
		return errConnectionNotActive
	}

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
	if !sc.IsActive() {
		return nil, errConnectionNotActive
	}

	// Read the 8-byte header
	header := make([]byte, 8)
	_, err := io.ReadFull(sc.conn, header)
	if err != nil {
		return nil, err
	}

	respId := int32(binary.LittleEndian.Uint32(header[:4]))
	_ = respId
	length := int32(binary.LittleEndian.Uint32(header[4:]))

	answer := make([]byte, length)
	_, err = io.ReadFull(sc.conn, answer)
	if err != nil {
		return nil, err
	}

	return sc.xor(answer), nil
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

func (sc *ServerConnection) xor(data []byte) []byte {
	if sc.key == nil {
		return data
	}

	for idx := range data {
		data[idx] = data[idx] ^ sc.key[idx%len(sc.key)]
	}
	return data
}
