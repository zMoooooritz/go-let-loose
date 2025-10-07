package rcon

import (
	"context"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socket"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type WorkerManager struct {
	config            ServerConfig
	jobChannel        chan rconJob
	workers           int
	workerLock        sync.Mutex
	waitGroup         *sync.WaitGroup
	stopWorkerChannel chan struct{}
	context           context.Context
	cancel            context.CancelFunc
}

func newWorkerManager(cfg ServerConfig, jobChannel chan rconJob) *WorkerManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerManager{
		config:            cfg,
		jobChannel:        jobChannel,
		workers:           0,
		workerLock:        sync.Mutex{},
		waitGroup:         &sync.WaitGroup{},
		stopWorkerChannel: make(chan struct{}),
		context:           ctx,
		cancel:            cancel,
	}
}

func (wm *WorkerManager) Start(count int) {
	wm.waitGroup.Add(count)
	for range count {
		go wm.worker()
	}
}

func (wm *WorkerManager) Stop(count int) {
	wc := wm.workerCount()
	if count > wc {
		count = wc
	}
	for range count {
		wm.stopWorkerChannel <- struct{}{}
	}
}

func (wm *WorkerManager) Count() int {
	return wm.workerCount()
}

func (wm *WorkerManager) Close() {
	wm.cancel()
	wm.waitGroup.Wait()
}

func (wm *WorkerManager) modifyWorkerCount(delta int) {
	wm.workerLock.Lock()
	defer wm.workerLock.Unlock()
	wm.workers += delta
	if wm.workers < 0 {
		wm.workers = 0
	}
}

func (wm *WorkerManager) workerCount() int {
	wm.workerLock.Lock()
	defer wm.workerLock.Unlock()
	return wm.workers
}

func (wm *WorkerManager) worker() {
	sc, err := socket.NewConnection(wm.config.Host, wm.config.Port, wm.config.Password, RCON_VERSION)
	if err != nil {
		logger.Warn("worker: failed to create connection", err)
		sc.Close()
		wm.waitGroup.Done()
		return
	}
	defer sc.Close()
	defer wm.waitGroup.Done()

	wm.modifyWorkerCount(1)

	for {
		select {
		case job := <-wm.jobChannel:
			ctx, cancel := context.WithTimeout(context.Background(), socket.CMD_TIMEOUT)
			resp, err := sc.Execute(ctx, job.Data.Command, job.Data.Body)
			cancel()

			if err == nil {
				if job.Response != nil {
					job.Response <- resp
				}
				continue
			}

			logger.Info("worker: recreating connection", err)
			time.Sleep(sleepTimeout)

			err = sc.Reconnect()
			if err != nil {
				logger.Warn("worker: creating new connection failed", err)
				if job.Error != nil {
					job.Error <- err
				}
				continue
			}

			ctx, cancel = context.WithTimeout(context.Background(), socket.CMD_TIMEOUT)
			resp, err = sc.Execute(ctx, job.Data.Command, job.Data.Body)
			cancel()

			if err == nil && job.Response != nil {
				job.Response <- resp
			} else if err != nil && job.Error != nil {
				job.Error <- err
			}
		case <-wm.stopWorkerChannel:
			// logger.Debug("worker: received stop signal")
			wm.modifyWorkerCount(-1)
			return
		case <-wm.context.Done():
			// logger.Debug("worker: received global stop signal")
			wm.modifyWorkerCount(-1)
			return
		}
	}
}
