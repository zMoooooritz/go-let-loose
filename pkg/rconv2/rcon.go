package rconv2

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

const (
	jobChannelSize = 100

	RCON_VERSION = 2
)

var (
	fallbackTimeout = 10 * time.Second
	sleepTimeout    = time.Second

	errTimeout = errors.New("timeout error")
)

type ServerConfig struct {
	Host     string
	Port     string
	Password string
}

type CommandData struct {
	Command string
	Body    string
}

type RconJob struct {
	Data     CommandData
	Response chan string
}

type Rcon struct {
	config     ServerConfig
	jobChannel chan RconJob
	waitGroup  *sync.WaitGroup
	context    context.Context
	cancel     context.CancelFunc
}

func NewRcon(cfg ServerConfig, workerCount int) (*Rcon, error) {
	// test for correct credentials
	sc, err := socketv2.NewConnection(cfg.Host, cfg.Port, cfg.Password, RCON_VERSION)
	if err != nil {
		return &Rcon{}, errors.New("invalid credentials provided")
	}
	sc.Close()

	waitGroup := sync.WaitGroup{}
	context, cancel := context.WithCancel(context.Background())

	rcon := Rcon{
		config:     cfg,
		jobChannel: make(chan RconJob, jobChannelSize),
		waitGroup:  &waitGroup,
		context:    context,
		cancel:     cancel,
	}

	for range workerCount {
		waitGroup.Add(1)
		go rcon.worker()
	}

	return &rcon, nil
}

func (r *Rcon) Close() {
	r.cancel()
	r.waitGroup.Wait()
}

func (r *Rcon) worker() {
	sc, _ := socketv2.NewConnection(r.config.Host, r.config.Port, r.config.Password, RCON_VERSION)
	defer sc.Close()
	defer r.waitGroup.Done()

	for {
		select {
		case job := <-r.jobChannel:
			resp, err := sc.Execute(job.Data.Command, job.Data.Body)
			if err == nil {
				if job.Response != nil {
					job.Response <- resp
				}
				continue
			}

			logger.Warn("worker: recreating connection")
			time.Sleep(sleepTimeout)
			err = sc.Reconnect()
			if err != nil {
				logger.Warn("worker: creating new connection failed", err)
			}
			time.Sleep(sleepTimeout)

			resp, err = sc.Execute(job.Data.Command, job.Data.Body)
			if err == nil && job.Response != nil {
				job.Response <- resp
			}
		case <-r.context.Done():
			return
		}
	}
}

func (r *Rcon) QueueJob(job RconJob) {
	r.jobChannel <- job
}

func runCommand[T, U any](rcn *Rcon, req T) (*U, error) {
	request := socketv2.RconRequest[T]{Body: req}
	cmd, body := request.ToArgs()

	recvChan := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), fallbackTimeout)
	defer cancel()

	go func() {
		rcn.QueueJob(RconJob{CommandData{cmd, body}, recvChan})
	}()

	select {
	case data := <-recvChan:
		var result U
		if _, ok := any(result).(string); ok {
			result = any(data).(U)
			return &result, nil
		}

		err := json.Unmarshal([]byte(data), &result)
		return &result, err
	case <-ctx.Done():
		var result U
		logger.Warn("runCommand: timeout occurred", "cmd:", cmd, "body:", body)
		return &result, errTimeout
	}
}
