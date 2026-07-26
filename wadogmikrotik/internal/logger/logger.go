package logger

import (
    "os"

    "github.com/natefinch/lumberjack"
    "github.com/sirupsen/logrus"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/config"
)

func New(cfg *config.Config) (*logrus.Logger, error) {
    if err := os.MkdirAll("logs", 0o755); err != nil {
        return nil, err
    }

    output := &lumberjack.Logger{
        Filename:   cfg.LogFile,
        MaxSize:    cfg.LogMaxSize,
        MaxBackups: cfg.LogMaxBackups,
        MaxAge:     cfg.LogMaxAge,
        Compress:   false,
    }

    logger := logrus.New()
    logger.SetOutput(output)
    logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    logger.SetLevel(logrus.InfoLevel)

    return logger, nil
}
