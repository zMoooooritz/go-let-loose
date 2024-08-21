package event

import (
	"context"
	"sync"
)

func EventHandler(events <-chan Event, eventNotifier *EventNotifier, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-events:
			if !ok {
				return
			}
			eventNotifier.notify(event)
		}
	}
}
