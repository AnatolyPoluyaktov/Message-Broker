package app

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/config"
	"github.com/samber/do/v2"
)

var diContainer *do.RootScope

func initDI() {
	diContainer = do.New()
}

func registerServices(cfg *config.Config) {
	initDI()
	do.ProvideValue(diContainer, cfg)

}
