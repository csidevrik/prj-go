package monitor

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/csidevrik/prj-go/wadogmikotik/internal/config"
	"github.com/csidevrik/prj-go/wadogmikotik/internal/network"
	"github.com/csidevrik/prj-go/wadogmikotik/internal/recovery"
)

type Service interface {
	Start(ctx context.Context) error
	Stop()
}

type service struct {
	cfg              *config.Config
	logger           *logrus.Logger
	recoveryService  recovery.Service
	netClient        network.Pinger
	ctx              context.Context
	cancel           context.CancelFunc
	isRunning        bool
}

func New(cfg *config.Config, logger *logrus.Logger, recoveryService recovery.Service) Service {
	return &service{
		cfg:             cfg,
		logger:          logger,
		recoveryService: recoveryService,
		netClient:       network.New(),
	}
}

func (s *service) Start(ctx context.Context) error {
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.isRunning = true

	s.logger.Info("Iniciando servicio de monitoreo de conectividad")
	go s.monitorLoop()

	return nil
}

func (s *service) Stop() {
	if s.isRunning && s.cancel != nil {
		s.cancel()
		s.isRunning = false
		s.logger.Info("Servicio de monitoreo detenido")
	}
}

func (s *service) monitorLoop() {
	ticker := time.NewTicker(s.cfg.MonitorInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			s.logger.Info("Ciclo de monitoreo terminado")
			return
		case <-ticker.C:
			s.checkConnectivity()
		}
	}
}

func (s *service) checkConnectivity() {
	hasInternet := s.hasInternetConnectivity()

	if hasInternet {
		s.logger.Debug("Conectividad a Internet verificada")
		return
	}

	s.logger.Warn("Pérdida de conectividad a Internet detectada")

	if err := s.recoveryService.DiagnoseAndRecover(); err != nil {
		s.logger.Warnf("Error durante diagnóstico/recuperación: %v", err)
	}
}

func (s *service) hasInternetConnectivity() bool {
	for _, target := range s.cfg.InternetTargets {
		ok, err := s.netClient.Ping(target, s.cfg.PingTimeout, s.cfg.RetryAttempts, s.cfg.RetryDelay)
		if err != nil {
			s.logger.Debugf("Error haciendo ping a %s: %v", target, err)
			continue
		}
		if ok {
			return true
		}
	}
	return false
}
