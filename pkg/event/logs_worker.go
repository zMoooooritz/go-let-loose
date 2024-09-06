package event

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/socket"
	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

func LogsFetcher(cfg rcon.ServerConfig, events chan<- Event, ctx context.Context, wg *sync.WaitGroup) {
	initialRun := true
	lastSeenTime := 0
	processedLogs := make(map[string]bool)

	defer wg.Done()

	sc, _ := socket.NewConnection(cfg.Host, cfg.Port, cfg.Password)
	defer sc.Close()
	sc.EnableFastUnsafeLogsFetching()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			loglines, err := sc.Execute("showlog 1", config.RF_UNINDEXEDLIST)
			if errors.Is(err, socket.InvalidRconCommand) {
				logger.Error("showlog 1", err, "this should not happen")
				continue
			}

			if err != nil {
				logger.Error("showlog 1", err, "recreating connection")
				time.Sleep(time.Second)
				err = sc.Reconnect()
				if err != nil {
					logger.Error("showlog: creating new connection failed", err)
				}
				time.Sleep(time.Second)
				continue
			}

			for _, line := range loglines {
				match := logPattern.FindStringSubmatch(line)
				if len(match) < 3 {
					continue
				}
				timestamp := util.ToInt(match[1])
				currentLine := match[2]

				if timestamp < lastSeenTime {
					continue
				}

				if timestamp == lastSeenTime && processedLogs[currentLine] {
					continue
				}

				if !initialRun { // ignore past events on startup
					for _, event := range logToEvents(line) {
						events <- event
					}
				}

				if timestamp > lastSeenTime {
					lastSeenTime = timestamp
					processedLogs = make(map[string]bool)
				}

				processedLogs[currentLine] = true
			}

			initialRun = false

			time.Sleep(400 * time.Millisecond)
		}
	}
}
