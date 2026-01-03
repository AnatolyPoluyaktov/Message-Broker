package broker

import (
	"sync"

	"github.com/google/uuid"
)

type Queue struct {
	msgs chan Message
	mu   *sync.Mutex
	cond *sync.Cond
	subs map[uuid.UUID]chan Message
}

func (q *Queue) Publish(msg Message) {
	q.msgs <- msg
	q.cond.Signal()
}

func NewQueue(maxMessages int64) *Queue {
	msgCh := make(chan Message, maxMessages)
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	subs := make(map[uuid.UUID]chan Message)
	q := &Queue{
		msgs: msgCh,
		mu:   mu,
		cond: cond,
		subs: subs,
	}
	go q.handleMessages()
	return q
}

func (q *Queue) handleMessages() {
	for msg := range q.msgs {
		if len(q.subs) == 0 {
			q.cond.Wait()
		}

		for _, v := range q.subs {
			v <- msg
		}
	}
}
