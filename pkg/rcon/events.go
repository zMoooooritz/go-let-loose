package rcon

import (
	"context"
	"sync"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

const (
	channel_size = 100
)

type rconEvents struct {
	enabled bool
	eventSystem
}

func WithEvents() RconOption {
	return func(r *Rcon) {
		r.Events = &rconEvents{
			enabled:     true,
			eventSystem: *newEventSystem(r),
		}
	}
}

func WithoutEvents() RconOption {
	return func(r *Rcon) {
		r.Events = &rconEvents{
			enabled: false,
		}
	}
}

type eventSystem struct {
	*eventNotifier

	context   context.Context
	cancel    context.CancelFunc
	waitGroup *sync.WaitGroup
}

func newEventSystem(rcn *Rcon) *eventSystem {
	eventChannel := make(chan hll.Event, channel_size)

	waitGroup := &sync.WaitGroup{}
	context, cancel := context.WithCancel(context.Background())
	eventNotifier := newEventNotifier()

	waitGroup.Add(3)
	go eventHandlerRoutine(eventChannel, eventNotifier, context, waitGroup)
	go logsFetcherRoutine(rcn, eventChannel, context, waitGroup)
	go serverInfoFetcherRoutine(rcn, eventChannel, context, waitGroup)

	return &eventSystem{
		eventNotifier,
		context,
		cancel,
		waitGroup,
	}
}

func (e *eventSystem) close() {
	e.cancel()
	e.waitGroup.Wait()
}

type eventObserver interface {
	Notify(hll.Event)
}

type eventNotifier struct {
	observers        map[eventObserver]struct{}
	eventOberservers map[hll.EventType][]eventObserver
}

func newEventNotifier() *eventNotifier {
	return &eventNotifier{
		observers:        make(map[eventObserver]struct{}),
		eventOberservers: make(map[hll.EventType][]eventObserver),
	}
}

func (n *eventNotifier) Register(o eventObserver) {
	n.observers[o] = struct{}{}
}

func (n *eventNotifier) Unregister(o eventObserver) {
	delete(n.observers, o)
}

func (n *eventNotifier) registerEvent(event hll.EventType, o eventObserver) {
	if n.eventOberservers[event] == nil {
		n.eventOberservers[event] = make([]eventObserver, 0)
	}
	n.eventOberservers[event] = append(n.eventOberservers[event], o)
}

func (n *eventNotifier) notify(e hll.Event) {
	for observer := range n.observers {
		observer.Notify(e)
	}
	if n.eventOberservers[e.Type()] != nil {
		for _, observer := range n.eventOberservers[e.Type()] {
			observer.Notify(e)
		}
	}
}
