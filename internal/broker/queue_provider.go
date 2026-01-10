package broker

import (
	"sync"

	"github.com/google/uuid"
)

type QueueProvider struct {
	mu     *sync.RWMutex
	queues map[uuid.UUID]*Queue
	names  map[uuid.UUID]string
}

func NewQueueProvider() (*QueueProvider, error) {
	mu := &sync.RWMutex{}
	queues := make(map[uuid.UUID]*Queue)
	names := make(map[uuid.UUID]string)
	return &QueueProvider{
		mu:     mu,
		queues: queues,
		names:  names,
	}, nil
}
func (qp *QueueProvider) Get(qID uuid.UUID) *Queue {
	qp.mu.RLock()
	defer qp.mu.RUnlock()
	return qp.queues[qID]
}

func (qp *QueueProvider) AddQueue(name string, maxMessages int64) {
	qp.mu.Lock()
	defer qp.mu.Unlock()
	i, _ := uuid.NewUUID()
	qp.names[i] = name
	qp.queues[i] = NewQueue(maxMessages)
}
