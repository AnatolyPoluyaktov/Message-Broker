package usecase

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/usecase/fetcher"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/usecase/producer"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/usecase/subscription"
	"github.com/samber/do/v2"
)

func resolveFetcher(injector do.Injector) (*fetcher.MessageFetcherUC, error) {
	qProvider := do.MustInvokeAs[broker.QueueGetter](injector)
	return fetcher.NewMessageFetcherUC(qProvider)
}

func resolveProducer(injector do.Injector) (*producer.ProducerUC, error) {
	qProvider := do.MustInvokeAs[broker.QueueGetter](injector)
	return producer.NewProducerUC(qProvider)
}
func resolveSubscriber(injector do.Injector) (*subscription.SubscriberUC, error) {
	qProvider := do.MustInvokeAs[broker.QueueGetter](injector)
	return subscription.NewSubsuscriberUC(qProvider)
}

func InitUseCases(injector do.Injector) {
	do.Package(
		do.Lazy(resolveProducer),
		do.Lazy(resolveFetcher),
		do.Lazy(resolveSubscriber),
	)(injector)
}
