# Script de instalación del servicio WadoG MikroTik Recovery
# Ejecutar como Administrador

param(
    [Parameter(Mandatory=$false)]
    [ValidateSet("install", "uninstall", "start", "stop", "restart", "status")]
    [string]$Action = "install"
)

# Verificar que se ejecuta como administrador
if (-not ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Error "Este script debe ser ejecutado como Administrador"
    exit 1
}

$ServiceName = "WadoGmikrotik"
$DisplayName = "WadoG MikroTik Recovery Service"
$ExePath = Join-Path $PSScriptRoot "service.exe"
$ConfigPath = Join-Path $PSScriptRoot "configs/config.yaml"

if (-not (Test-Path $ExePath)) {
    Write-Error "Ejecutable no encontrado: $ExePath"
    Write-Host "Primero debes compilar el proyecto:"
    Write-Host "  go build -o service.exe ./cmd/service"
    exit 1
}

if (-not (Test-Path $ConfigPath)) {
    Write-Error "Archivo de configuración no encontrado: $ConfigPath"
    Write-Host "Copia configs/config.example.yaml a configs/config.yaml y ajusta los valores"
    exit 1
}

Write-Host "========================================" -ForegroundColor Green
Write-Host "WadoG MikroTik Recovery Service Manager" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Green

switch ($Action) {
    "install" {
        Write-Host "`nInstalando servicio '$DisplayName'..."
        & "$ExePath" install
        if ($LASTEXITCODE -eq 0) {
            Write-Host "✓ Servicio instalado exitosamente" -ForegroundColor Green
            Write-Host "  Próximo paso: .\install-service.ps1 -Action start"
        } else {
            Write-Host "✗ Error al instalar el servicio" -ForegroundColor Red
            exit 1
        }
    }

    "uninstall" {
        Write-Host "`nDesinstalando servicio '$DisplayName'..."
        & "$ExePath" uninstall
        if ($LASTEXITCODE -eq 0) {
            Write-Host "✓ Servicio desinstalado exitosamente" -ForegroundColor Green
        } else {
            Write-Host "✗ Error al desinstalar el servicio" -ForegroundColor Red
            exit 1
        }
    }

    "start" {
        Write-Host "`nIniciando servicio '$DisplayName'..."
        & "$ExePath" start
        if ($LASTEXITCODE -eq 0) {
            Write-Host "✓ Servicio iniciado exitosamente" -ForegroundColor Green
            Start-Sleep 2
            & "$ExePath" status
        } else {
            Write-Host "✗ Error al iniciar el servicio" -ForegroundColor Red
            exit 1
        }
    }

    "stop" {
        Write-Host "`nDeteniendo servicio '$DisplayName'..."
        & "$ExePath" stop
        if ($LASTEXITCODE -eq 0) {
            Write-Host "✓ Servicio detenido exitosamente" -ForegroundColor Green
        } else {
            Write-Host "✗ Error al detener el servicio" -ForegroundColor Red
            exit 1
        }
    }

    "restart" {
        Write-Host "`nReiniciando servicio '$DisplayName'..."
        & "$ExePath" stop
        Start-Sleep 1
        & "$ExePath" start
        if ($LASTEXITCODE -eq 0) {
            Write-Host "✓ Servicio reiniciado exitosamente" -ForegroundColor Green
        } else {
            Write-Host "✗ Error al reiniciar el servicio" -ForegroundColor Red
            exit 1
        }
    }

    "status" {
        Write-Host "`nObteniendo estado del servicio..."
        $service = Get-Service -Name $ServiceName -ErrorAction SilentlyContinue
        if ($service) {
            Write-Host "Nombre:      $($service.Name)"
            Write-Host "Estado:      $($service.Status)"
            Write-Host "Inicio:      $($service.StartType)"
            if ($service.Status -eq "Running") {
                Write-Host "✓ El servicio está en ejecución" -ForegroundColor Green
            } else {
                Write-Host "✗ El servicio no está en ejecución" -ForegroundColor Red
            }
        } else {
            Write-Host "✗ El servicio no está instalado" -ForegroundColor Red
            exit 1
        }
    }

    default {
        Write-Host "Acción no válida: $Action"
        exit 1
    }
}

Write-Host ""
