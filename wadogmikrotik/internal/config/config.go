package config

import (
    "fmt"
    "os"
    "time"

    "gopkg.in/yaml.v3"
)

type SSHConfig struct {
    User          string        `yaml:"user"`
    Password      string        `yaml:"password,omitempty"`
    PrivateKey    string        `yaml:"private_key,omitempty"`
    KeyPassphrase string        `yaml:"key_passphrase,omitempty"`
    Timeout       time.Duration `yaml:"timeout,omitempty"`
}

type SideConfig struct {
    IP  string    `yaml:"ip"`
    SSH SSHConfig `yaml:"ssh"`
}

type Config struct {
    SideA           SideConfig      `yaml:"side_a"`
    SideB           SideConfig      `yaml:"side_b"`
    InternetTargets []string        `yaml:"internet_targets"`
    MonitorInterval time.Duration   `yaml:"monitor_interval"`
    RecoveryCooldown time.Duration  `yaml:"recovery_cooldown"`
    PingTimeout     time.Duration   `yaml:"ping_timeout"`
    RetryAttempts   int             `yaml:"retry_attempts"`
    RetryDelay      time.Duration   `yaml:"retry_delay"`
    LogFile         string          `yaml:"log_file"`
    LogMaxSize      int             `yaml:"log_max_size"`
    LogMaxBackups   int             `yaml:"log_max_backups"`
    LogMaxAge       int             `yaml:"log_max_age"`
}

func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("no se puede leer config %s: %w", path, err)
    }

    var cfg Config
    if err = yaml.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("error parseando YAML: %w", err)
    }

    return validate(&cfg)
}

func validate(cfg *Config) (*Config, error) {
    if cfg.SideA.IP == "" {
        return nil, fmt.Errorf("side_a.ip no puede estar vacío")
    }
    if cfg.SideA.SSH.User == "" {
        return nil, fmt.Errorf("side_a.ssh.user no puede estar vacío")
    }
    if cfg.SideA.SSH.Password == "" && cfg.SideA.SSH.PrivateKey == "" {
        return nil, fmt.Errorf("side_a.ssh.password o side_a.ssh.private_key debe proporcionarse")
    }
    if cfg.SideB.IP == "" {
        return nil, fmt.Errorf("side_b.ip no puede estar vacío")
    }
    if cfg.SideB.SSH.User == "" {
        return nil, fmt.Errorf("side_b.ssh.user no puede estar vacío")
    }
    if cfg.SideB.SSH.Password == "" && cfg.SideB.SSH.PrivateKey == "" {
        return nil, fmt.Errorf("side_b.ssh.password o side_b.ssh.private_key debe proporcionarse")
    }
    if len(cfg.InternetTargets) == 0 {
        return nil, fmt.Errorf("internet_targets no puede estar vacío")
    }
    if cfg.MonitorInterval <= 0 {
        cfg.MonitorInterval = 5 * time.Second
    }
    if cfg.RecoveryCooldown <= 0 {
        cfg.RecoveryCooldown = 300 * time.Second
    }
    if cfg.PingTimeout <= 0 {
        cfg.PingTimeout = 5 * time.Second
    }
    if cfg.RetryAttempts < 1 {
        cfg.RetryAttempts = 1
    }
    if cfg.RetryDelay <= 0 {
        cfg.RetryDelay = 2 * time.Second
    }
    if cfg.LogFile == "" {
        cfg.LogFile = "logs/wadogmikrotik.log"
    }
    if cfg.LogMaxSize <= 0 {
        cfg.LogMaxSize = 10
    }
    if cfg.LogMaxBackups < 0 {
        cfg.LogMaxBackups = 5
    }
    if cfg.LogMaxAge < 0 {
        cfg.LogMaxAge = 30
    }

    cfg.SideA.SSH.Timeout = 10 * time.Second
    cfg.SideB.SSH.Timeout = 10 * time.Second

    return cfg, nil
}
