package mediator

import (
	"github.com/google/uuid"
	"log"
	"sync"
)

type CallbackId string
type callbackFn func(data any) error

type callback struct {
	id CallbackId
	fn callbackFn
}

type PubSub struct {
	handles map[string][]callback
	mu      sync.RWMutex
}

func InitPubSub() *PubSub {
	return &PubSub{
		handles: make(map[string][]callback),
	}
}

func (ps *PubSub) Subscribe(event string, fn callbackFn) (CallbackId, func()) {
	id := CallbackId(uuid.NewString())
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.handles[event] = append(ps.handles[event], callback{
		id: id,
		fn: fn,
	})

	return id, func() {
		ps.unsubscribe(event, id)
	}

}

func (ps *PubSub) unsubscribe(event string, id CallbackId) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	callbacks, ok := ps.handles[event]
	if !ok {
		return
	}

	updated := callbacks[:0]
	for _, cb := range callbacks {
		if cb.id != id {
			updated = append(updated, cb)
		}
	}

	ps.handles[event] = updated
}

func (ps *PubSub) Publish(event string, data any) {
	ps.mu.RLock()
	callbacks := ps.handles[event]
	ps.mu.RUnlock()

	var wg sync.WaitGroup

	for _, cb := range callbacks {
		wg.Add(1)
		go func(cb callback) {
			defer wg.Done()
			if err := cb.fn(data); err != nil {
				log.Printf("PubSub error on event '%s' (ID: %s): %v", event, cb.id, err)
			}
		}(cb)
	}

	go wg.Wait()
}
