package app

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/config"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/usecase"

	"github.com/samber/do/v2"
)

var diContainer *do.RootScope

func initDI() {
	diContainer = do.New()
}

func InvokeServiceAs[T any]() T {
	return do.MustInvokeAs[T](diContainer)
}

func registerServices(cfg *config.Config) func() {
	initDI()
	do.ProvideValue(diContainer, cfg)
	broker.InitBroker(diContainer)
	usecase.InitUseCases(diContainer)
	return func() {
		diContainer.Shutdown()
	}
}
