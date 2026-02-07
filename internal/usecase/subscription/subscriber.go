package subscription

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/google/uuid"
)

type SubscriberUC struct {
	queueProvider broker.QueueGetter
}

func NewSubsuscriberUC(queueProvider broker.QueueGetter) (*SubscriberUC, error) {
	return &SubscriberUC{
		queueProvider: queueProvider,
	}, nil
}

func (msuc *SubscriberUC) Subscribe(queueName broker.QueueName) (uuid.UUID, error) {
	queue := msuc.queueProvider.Get(queueName)
	return queue.AddSubscriber()
}
