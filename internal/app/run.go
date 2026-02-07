package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AnatolyPoluyaktov/msgbroker/internal/broker"
	"github.com/AnatolyPoluyaktov/msgbroker/internal/config"
	httpHandler "github.com/AnatolyPoluyaktov/msgbroker/internal/controller/http"
	"github.com/AnatolyPoluyaktov/msgbroker/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/samber/do/v2"
)

type Service struct {
	cfg    *config.Config
	name   string
	stopFN func()
}

func NewService(cfg *config.Config, name string) *Service {
	return &Service{
		cfg:  cfg,
		name: name,
	}
}

func (s *Service) Run() {
	logger.Info("Service starting....", "name", s.name)
	services := registerServices(s.cfg)
	s.registerQueues()
	routers:= s.registerRouters()

	logger.Info("Service started successfully", "name", s.name)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Info("Service stopping", "name", s.name)

}

func (s *Service) registerRouters() *mux.Router {

	msgHandler := InvokeServiceAs[*httpHandler.MessageHandler]()
	return httpHandler.NewRouter(msgHandler)

}


unc (s *Service) RunHTTPServer(
	ctx context.Context,
	r *mux.Router,
) (func(context.Context) error, error) {

	srv := &http.Server{
		Addr:         s.cfg.Server.Address,
		Handler:      r,
		ReadTimeout:  s.cfg.Server.Timeout,
		WriteTimeout: s.cfg.Server.Timeout,
	}

	go func() {
		logger.Info("HTTP server started", "addr", srv.Addr)

		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			logger.Error("HTTP server failed", "err", err)
		}
	}()

	shutdown := func(ctx context.Context) error {
		logger.Info("HTTP server shutting down")
		return srv.Shutdown(ctx)
	}

	return shutdown, nil
}


func (s *Service) registerQueues() {

	queueProvider := do.MustInvoke[*broker.QueueProvider](diContainer)
	for qname, qcfg := range s.cfg.Queues {
		queueProvider.AddQueue(broker.QueueName(qname), qcfg.MaxItems)
	}
}
