package windows

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"github.com/csidevrik/prj-go/wadogmikotik/internal/monitor"
)

type Program struct {
	Monitor monitor.Service
	Logger  *logrus.Logger
}

func (p *Program) Start(s service.Service) error {
	p.Logger.Info("Servicio iniciado")
	ctx := context.Background()
	return p.Monitor.Start(ctx)
}

func (p *Program) Stop(s service.Service) error {
	p.Logger.Info("Deteniendo servicio...")
	p.Monitor.Stop()
	return nil
}

func Run(program *Program, svcConfig *service.Config, logger *logrus.Logger) error {
	prg := program

	s, err := service.New(prg, svcConfig)
	if err != nil {
		return err
	}

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		err := service.Control(s, cmd)
		if err != nil {
			logger.Warnf("Comando '%s' no soportado o error: %v", cmd, err)
			fmt.Printf("Comandos disponibles: install, uninstall, start, stop, restart, status\n")
		}
		return nil
	}

	go handleSignals(program, logger)

	err = s.Run()
	if err != nil {
		logger.Errorf("Error ejecutando servicio: %v", err)
		return err
	}

	return nil
}

func handleSignals(program *Program, logger *logrus.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	logger.Info("Señal de terminación recibida")
	program.Monitor.Stop()
}
