package broker

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"
)

const MaxSubscriberMessages = 1000

type Message any
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

func (q *Queue) GetMessages(subID uuid.UUID) ([]Message, error) {

	q.mu.Lock()
	defer q.mu.Unlock()
	msgCh, ok := q.subs[subID]
	if !ok {
		return nil, errors.Errorf("Not found Subscriber by ID ( %s )", subID.String())
	}
	msgBuf := make([]Message, 0, len(msgCh))
	for {
		select {
		case msg := <-msgCh:
			_ = append(msgBuf, msg)

		default:
			return msgBuf, nil
		}
	}

}

func (q *Queue) AddSubscriber() (uuid.UUID, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	subID, err := uuid.NewUUID()

	if err != nil {
		return uuid.UUID{}, nil
	}
	subCh := make(chan Message, MaxSubscriberMessages)
	q.subs[subID] = subCh
	return subID, nil
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
