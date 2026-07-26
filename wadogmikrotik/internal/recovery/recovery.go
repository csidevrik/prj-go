package recovery

import (
    "fmt"
    "sync"
    "time"

    "github.com/sirupsen/logrus"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/config"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/network"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/ssh"
)

type Service interface {
    DiagnoseAndRecover() error
    CanRecover() bool
}

type service struct {
    cfg        *config.Config
    logger     *logrus.Logger
    netClient  network.Pinger
    sshClient  ssh.Client
    mutex      sync.Mutex
    lastRecovery time.Time
}

func New(cfg *config.Config, logger *logrus.Logger) Service {
    return &service{
        cfg:       cfg,
        logger:    logger,
        netClient: network.New(),
        sshClient: ssh.New(),
    }
}

func (s *service) CanRecover() bool {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return time.Since(s.lastRecovery) >= s.cfg.RecoveryCooldown
}

func (s *service) setRecoveryTime() {
    s.mutex.Lock()
    s.lastRecovery = time.Now()
    s.mutex.Unlock()
}

func (s *service) DiagnoseAndRecover() error {
    if !s.CanRecover() {
        s.logger.Infof("Esperando cooldown antes del nuevo intento de recuperación")
        return nil
    }

    s.logger.Info("Iniciando diagnóstico de antenas MikroTik")

    okA, err := s.netClient.Ping(s.cfg.SideB.IP, s.cfg.PingTimeout, s.cfg.RetryAttempts, s.cfg.RetryDelay)
    if err != nil {
        return fmt.Errorf("error al hacer ping a Lado B: %w", err)
    }
    s.logger.WithFields(logrus.Fields{"side": "B", "ip": s.cfg.SideB.IP, "reachable": okA}).Info("Resultado del ping de Lado B")
    if !okA {
        return fmt.Errorf("Lado B no responde al ping")
    }

    okB, err := s.netClient.Ping(s.cfg.SideA.IP, s.cfg.PingTimeout, s.cfg.RetryAttempts, s.cfg.RetryDelay)
    if err != nil {
        return fmt.Errorf("error al hacer ping a Lado A: %w", err)
    }
    s.logger.WithFields(logrus.Fields{"side": "A", "ip": s.cfg.SideA.IP, "reachable": okB}).Info("Resultado del ping de Lado A")
    if !okB {
        return fmt.Errorf("Lado A no responde al ping")
    }

    s.logger.Info("Ambas antenas responden, ejecutando recuperación remota")
    if err := s.rebootSide(s.cfg.SideA.IP, s.cfg.SideA.SSH, "Lado A"); err != nil {
        return fmt.Errorf("falló reinicio Lado A: %w", err)
    }

    if err := s.rebootSide(s.cfg.SideB.IP, s.cfg.SideB.SSH, "Lado B"); err != nil {
        return fmt.Errorf("falló reinicio Lado B: %w", err)
    }

    s.setRecoveryTime()
    s.logger.Info("Recuperación realizada y cooldown activado")

    return nil
}

func (s *service) rebootSide(ip string, sshCfg config.SSHConfig, label string) error {
    s.logger.WithFields(logrus.Fields{"side": label, "ip": ip}).Info("Enviando comando de reinicio")
    if err := s.sshClient.RunCommand(ip, sshCfg, "/system reboot"); err != nil {
        return err
    }
    s.logger.WithFields(logrus.Fields{"side": label, "ip": ip}).Info("Comando /system reboot enviado")
    return nil
}
