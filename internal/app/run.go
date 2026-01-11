package app

import (
	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/config"
	"github.com/AnatolyPoluyaktov/msgbroker/pkg/logger"
	"github.com/samber/do/v2"
)

type Service struct {
	cfg  *config.Config
	name string
}

func NewService(cfg *config.Config, name string) *Service {
	return &Service{
		cfg:  cfg,
		name: name,
	}
}

func (s *Service) Run() {
	logger.Info("Service starting....", "name", s.name)
	registerServices(s.cfg)
	s.registerQueues()
	logger.Info("Service started successfully", "name", s.name)
}

func (s *Service) registerQueues() {
	queueProvider := do.MustInvoke[*broker.QueueProvider](diContainer)
	for qname, qcfg := range s.cfg.Queues {
		queueProvider.AddQueue(qname, qcfg.MaxItems)
	}
}
