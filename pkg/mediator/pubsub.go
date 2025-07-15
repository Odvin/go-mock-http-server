package mediator

import (
	"sync"
)

var instance *PubSub
var once sync.Once

type topicTitle string
type TopicChanel chan any

type PubSub struct {
	topics map[topicTitle][]TopicChanel
	mu     sync.RWMutex
}

func GetPubSub() *PubSub {
	once.Do(func() {
		instance = &PubSub{
			topics: make(map[topicTitle][]TopicChanel),
		}
	})

	return instance
}

func (ps *PubSub) Subscribe(topic topicTitle) (TopicChanel, func()) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	chanel := TopicChanel(make(chan any))

	ps.topics[topic] = append(ps.topics[topic], chanel)

	return chanel, func() {
		ps.unsubscribe(topic, chanel)
	}
}

func (ps *PubSub) unsubscribe(topic topicTitle, chanel TopicChanel) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	channels, ok := ps.topics[topic]
	if !ok {
		return
	}

	updated := channels[:0]
	for _, ch := range channels {
		if ch != chanel {
			updated = append(updated, ch)
		} else {
			close(ch)
		}
	}

	ps.topics[topic] = updated
}

func (ps *PubSub) Publish(topic topicTitle, data any) {
	ps.mu.RLock()
	channels := ps.topics[topic]
	ps.mu.RUnlock()

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch TopicChanel) {
			defer wg.Done()
			ch <- data
		}(ch)
	}

	go wg.Wait()
}
