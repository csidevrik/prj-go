package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/kardianos/service"

    "github.com/csidevrik/prj-go/wadogmikotik/internal/config"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/logger"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/monitor"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/recovery"
    "github.com/csidevrik/prj-go/wadogmikotik/internal/windows"
)

func main() {
    configPath := flag.String("config", "configs/config.yaml", "Ruta al archivo de configuración YAML")
    flag.Parse()

    cfg, err := config.Load(*configPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error cargando configuración: %v\n", err)
        os.Exit(1)
    }

    loggerInstance, err := logger.New(cfg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error inicializando logger: %v\n", err)
        os.Exit(1)
    }

    recoveryService := recovery.New(cfg, loggerInstance)
    monitorService := monitor.New(cfg, loggerInstance, recoveryService)

    svcConfig := &service.Config{
        Name:        "WadoGmikrotik",
        DisplayName: "WadoG MikroTik Recovery Service",
        Description: "Monitorea la conectividad a Internet y reinicia antenas MikroTik cuando el enlace se bloquea.",
    }

    program := &windows.Program{Monitor: monitorService, Logger: loggerInstance}
    err = windows.Run(program, svcConfig, loggerInstance)
    if err != nil {
        loggerInstance.Errorf("Servicio terminado con error: %v", err)
        os.Exit(1)
    }
}
