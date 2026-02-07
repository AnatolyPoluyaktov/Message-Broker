package broker

import (
	"sync"
)

type QueueName string
type QueueGetter interface {
	Get(name QueueName) *Queue
}

type QueueProvider struct {
	mu     *sync.RWMutex
	queues map[QueueName]*Queue
}

func NewQueueProvider() *QueueProvider {
	queues := make(map[QueueName]*Queue)
	return &QueueProvider{
		mu:     &sync.RWMutex{},
		queues: queues,
	}
}
func (qp *QueueProvider) Get(name QueueName) *Queue {
	qp.mu.RLock()
	defer qp.mu.RUnlock()
	return qp.queues[name]
}

func (qp *QueueProvider) AddQueue(name QueueName, maxMessages int64) {
	qp.mu.Lock()
	defer qp.mu.Unlock()
	qp.queues[name] = NewQueue(maxMessages)
}
