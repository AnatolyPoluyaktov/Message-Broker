package producer

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
)

type ProducerUC struct {
	queueProvider broker.QueueGetter
}

func NewProducerUC(queueProvider broker.QueueGetter) (*ProducerUC, error) {
	return &ProducerUC{
		queueProvider: queueProvider,
	}, nil
}

func (puc *ProducerUC) ProduceMessage(queueName broker.QueueName, msg broker.Message) error {
	queue := puc.queueProvider.Get(queueName)
	queue.Publish(msg)
	return nil

}
