package event

import (
	"context"
	"sync"

	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

type EventListener struct {
	*EventNotifier

	context   context.Context
	cancel    context.CancelFunc
	waitGroup *sync.WaitGroup
}

func NewEventListener(rcn *rcon.Rcon, cach *Cache) *EventListener {
	eventChannel := make(chan Event, channel_size)

	waitGroup := &sync.WaitGroup{}
	context, cancel := context.WithCancel(context.Background())
	eventNotifier := NewEventNotifier()

	waitGroup.Add(3)
	go EventHandler(eventChannel, eventNotifier, context, waitGroup)
	go LogsFetcher(rcn.GetServerConfig(), eventChannel, context, waitGroup)
	go ServerInfoFetcher(rcn, cach, eventChannel, context, waitGroup)

	return &EventListener{
		eventNotifier,
		context,
		cancel,
		waitGroup,
	}
}

func (e *EventListener) Close() {
	e.cancel()
	e.waitGroup.Wait()
}
