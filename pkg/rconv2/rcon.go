package rconv2

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

const (
	jobChannelSize = 100

	RCON_VERSION = 2
)

var (
	fallbackTimeout = 30 * time.Second
	sleepTimeout    = time.Second

	errTimeout = errors.New("timeout error")
)

type ServerConfig struct {
	Host     string
	Port     string
	Password string
}

type commandData struct {
	Command string
	Body    string
}

type rconJob struct {
	Data     commandData
	Response chan string
	Error    chan error
}

func newRconJob(cmd, body string) rconJob {
	return rconJob{
		Data: commandData{
			Command: cmd,
			Body:    body,
		},
		Response: make(chan string, 1),
		Error:    make(chan error, 1),
	}
}

type Rcon struct {
	cache      *rconCache
	Worker     *WorkerManager
	jobChannel chan rconJob
}

func NewRcon(cfg ServerConfig, workerCount int, opts ...RconOption) (*Rcon, error) {
	// test for correct credentials
	sc, err := socketv2.NewConnection(cfg.Host, cfg.Port, cfg.Password, RCON_VERSION)
	if err != nil {
		return &Rcon{}, errors.New("invalid credentials provided")
	}
	sc.Close()

	jobChannel := make(chan rconJob, jobChannelSize)

	workerManager := newWorkerManager(cfg, jobChannel)

	rcon := Rcon{
		cache:      &rconCache{},
		Worker:     workerManager,
		jobChannel: jobChannel,
	}

	for _, opt := range opts {
		opt(&rcon)
	}

	rcon.Worker.Start(workerCount)

	return &rcon, nil
}

func (rcon *Rcon) Close() {
	rcon.cache.data.Stop()
	rcon.Worker.Close()
	close(rcon.jobChannel)
}

func runCommand[T, U any](rcn *Rcon, req T) (*U, error) {
	request := socketv2.RconRequest[T]{Body: req}
	cmd, body := request.ToArgs()

	cacheKey := cmd + "|" + body

	val, err := rcn.cache.get(cacheKey)
	if err == nil {
		cached := val.(*U)
		return cached, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), fallbackTimeout)
	rconJob := newRconJob(cmd, body)
	defer cancel()

	go func() {
		rcn.jobChannel <- rconJob
	}()

	select {
	case response := <-rconJob.Response:
		var result U
		if _, ok := any(result).(string); ok {
			result = any(response).(U)
			rcn.cache.set(cacheKey, &result)
			return &result, nil
		}

		err := json.Unmarshal([]byte(response), &result)

		if err == nil {
			rcn.cache.set(cacheKey, &result)
		}

		return &result, err
	case err := <-rconJob.Error:
		var result U
		logger.Warn("runCommand: error occurred", "cmd:", cmd, "body:", body, "err:", err)
		return &result, err
	case <-ctx.Done():
		var result U
		logger.Warn("runCommand: fallback timeout occurred", "cmd:", cmd, "body:", body)
		return &result, errTimeout
	}
}
