package broker

import (
	"sync"

	"github.com/google/uuid"
)

type QueueProvider struct {
	mu     sync.RWMutex
	queues map[uuid.UUID]Queue
	names  map[uuid.UUID]string
}

func (qp *QueueProvider) Get(qID uuid.UUID) Queue {
	qp.mu.RLock()
	defer qp.mu.RUnlock()
	return qp.Get(qID)
}

func (qp *QueueProvider) AddQueue(name string) {
	qp.mu.Lock()
	defer qp.mu.Unlock()
	i, _ := uuid.NewUUID()
	qp.names[i] = name
}
