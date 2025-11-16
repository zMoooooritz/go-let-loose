package socket

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

const (
	INIT_TIMEOUT = time.Duration(5 * time.Second)
	CMD_TIMEOUT  = time.Duration(3 * time.Second)
)

var (
	requestId uint32 = 1

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
	sc := ServerConnection{
		ip:       ip,
		port:     port,
		password: password,
		version:  version,
		conn:     nil,
		key:      nil,
	}
	err := sc.reconnect()
	return &sc, err
}

func (sc *ServerConnection) Execute(ctx context.Context, command, body string) (string, error) {
	if !sc.IsActive() {
		return "", errConnectionNotActive
	}

	if deadline, ok := ctx.Deadline(); ok {
		_ = sc.conn.SetDeadline(deadline)
	} else {
		_ = sc.conn.SetDeadline(time.Time{})
	}

	rconRequest := NewRawRequest(sc.authToken, sc.version, command, body)
	logger.Debug("Request: " + rconRequest.String())
	err := sc.write(rconRequest.Pack())
	if err != nil {
		return "", err
	}
	resp, err := sc.read()
	if err != nil {
		return "", err
	}
	rconResponse := RconResponse{}
	err = json.Unmarshal(resp, &rconResponse)
	if err != nil {
		return "", err
	}
	if rconResponse.StatusCode != StatusOk {
		return "", NewRconError(rconResponse.StatusCode, rconResponse.StatusMessage)
	}
	logger.Debug("Response: " + rconResponse.String())

	return string(rconResponse.ContentBody), nil
}

func (sc *ServerConnection) Reconnect() error {
	sc.Close()
	return sc.reconnect()
}

func (sc *ServerConnection) Close() {
	if sc.IsActive() {
		_ = sc.conn.Close()
	}
	sc.reset()
}

func (sc *ServerConnection) reset() {
	sc.conn = nil
	sc.key = nil
	sc.authToken = ""
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
	return err
}

func (sc *ServerConnection) connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), INIT_TIMEOUT)
	defer cancel()

	resp, err := sc.Execute(ctx, "ServerConnect", "")
	if err != nil {
		return err
	}
	sc.key, err = base64.StdEncoding.DecodeString(resp)
	return err
}

func (sc *ServerConnection) login() error {
	ctx, cancel := context.WithTimeout(context.Background(), INIT_TIMEOUT)
	defer cancel()

	resp, err := sc.Execute(ctx, "Login", sc.password)
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

	sc.xor(data)

	header := make([]byte, 8)
	binary.LittleEndian.PutUint32(header[0:4], requestId)
	binary.LittleEndian.PutUint32(header[4:8], uint32(len(data)))

	requestId++

	fullData := append(header, data...)

	n, err := sc.conn.Write(fullData)
	if err != nil {
		return err
	}
	if n != len(fullData) {
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

	sc.xor(answer)

	return answer, nil
}

func (sc *ServerConnection) xor(data []byte) {
	if sc.key == nil {
		return
	}

	for idx := range data {
		data[idx] = data[idx] ^ sc.key[idx%len(sc.key)]
	}
}
