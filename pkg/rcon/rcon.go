package rcon

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socket"
	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

const (
	jobChannelSize = 100
)

var (
	fallbackTimeout = time.Duration(30) * time.Second
	sleepTimeout    = time.Second

	timeoutErr = errors.New("timeout error")
)

type ServerConfig struct {
	Host     string
	Port     string
	Password string
}

type CommandData struct {
	Command string
	Format  config.ResponseFormat
}

type RconJob struct {
	Data     CommandData
	Response chan []string
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
	sc, err := socket.NewConnection(cfg.Host, cfg.Port, cfg.Password)
	if err != nil {
		return &Rcon{}, errors.New("Invalid credentials provided")
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

func (r *Rcon) GetServerConfig() ServerConfig {
	return r.config
}

func (r *Rcon) worker() {
	sc, _ := socket.NewConnection(r.config.Host, r.config.Port, r.config.Password)
	defer sc.Close()
	defer r.waitGroup.Done()

	for {
		select {
		case job := <-r.jobChannel:
			resp, err := sc.Execute(job.Data.Command, job.Data.Format)
			if err == nil {
				if job.Response != nil {
					job.Response <- resp
				}
				continue
			}

			if errors.Is(err, socket.InvalidRconCommand) {
				continue
			}

			logger.Warn("worker: recreating connection")
			logger.Warn(job.Data.Command[:min(30, len(job.Data.Command)-1)], err)
			time.Sleep(sleepTimeout)
			err = sc.Reconnect()
			if err != nil {
				logger.Warn("worker: creating new connection failed", err)
			}
			time.Sleep(sleepTimeout)

			resp, err = sc.Execute(job.Data.Command, job.Data.Format)
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

func (r *Rcon) runBasicCommand(command string) (string, error) {
	data, err := r.RunCommand(command, config.RF_DIRECT)
	if err != nil || len(data) == 0 {
		return "", err
	}
	return data[0], nil
}

func (r *Rcon) runListCommand(command string) ([]string, error) {
	return r.RunCommand(command, config.RF_INDEXEDLIST)
}

func (r *Rcon) runUnindexedListCommand(command string) ([]string, error) {
	return r.RunCommand(command, config.RF_UNINDEXEDLIST)
}

func (r *Rcon) RunCommand(command string, format config.ResponseFormat) ([]string, error) {
	recvChan := make(chan []string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), fallbackTimeout)
	defer cancel()

	go func() {
		r.QueueJob(RconJob{CommandData{command, format}, recvChan})
	}()

	select {
	case data := <-recvChan:
		return data, nil
	case <-ctx.Done():
		command = strings.ReplaceAll(command, config.NEWLINE, config.ESCAPED_NEWLINE)
		logger.Info("command: " + command + " timed out")
		return nil, timeoutErr
	}
}

func boolToToggleStr(enabled bool) string {
	if enabled {
		return "on"
	}
	return "off"
}

func runSetCommand(r *Rcon, cmd string) error {
	_, err := r.runBasicCommand(cmd)
	return err
}

func getNumVal(r *Rcon, cmd string) (int, error) {
	resp, err := r.runBasicCommand(cmd)
	if err != nil {
		return 0, err
	}
	return util.ToInt(resp), nil
}

func getBoolVal(r *Rcon, cmd string) (bool, error) {
	resp, err := r.runBasicCommand(cmd)
	if err != nil {
		return false, err
	}
	if resp == "1" || strings.Contains(resp, "TRUE") {
		return true, nil
	}
	return false, nil
}
