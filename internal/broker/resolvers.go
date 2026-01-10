package broker

import (
	"github.com/samber/do/v2"
)

func resolveQueueProvider(_ do.Injector) (*QueueProvider, error) {
	return NewQueueProvider()
}

func InitBroker(injector do.Injector) {
	do.Package(
		do.Lazy(resolveQueueProvider),
	)(injector)
}
