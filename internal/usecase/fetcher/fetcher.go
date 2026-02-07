package fetcher

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/google/uuid"
)

type MessageFetcherUC struct {
	queueProvider broker.QueueGetter
}

func NewMessageFetcherUC(queueProvider broker.QueueGetter) (*MessageFetcherUC, error) {
	return &MessageFetcherUC{
		queueProvider: queueProvider,
	}, nil
}

func (mfuc *MessageFetcherUC) FetchMessages(qName broker.QueueName, subID uuid.UUID) ([]broker.Message, error) {
	queue := mfuc.queueProvider.Get(qName)

	return queue.GetMessages(subID)

}
