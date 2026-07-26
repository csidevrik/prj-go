@echo off
REM Script de compilación para WadoG MikroTik Recovery Service

setlocal enabledelayedexpansion

echo.
echo ========================================
echo WadoG MikroTik Recovery Service - Build
echo ========================================
echo.

REM Verificar que Go está instalado
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Error: Go no está instalado o no está en el PATH
    pause
    exit /b 1
)

echo [1/3] Descargando dependencias...
call go mod download
if %ERRORLEVEL% NEQ 0 (
    echo Error al descargar dependencias
    pause
    exit /b 1
)

echo [2/3] Actualizando módulos...
call go mod tidy
if %ERRORLEVEL% NEQ 0 (
    echo Error al actualizar módulos
    pause
    exit /b 1
)

echo [3/3] Compilando...
call go build -o service.exe ./cmd/service
if %ERRORLEVEL% NEQ 0 (
    echo Error al compilar
    pause
    exit /b 1
)

echo.
echo ✓ Compilación exitosa!
echo.
echo El ejecutable está en: service.exe
echo.
echo Próximos pasos:
echo 1. Copia configs/config.example.yaml a configs/config.yaml
echo 2. Edita configs/config.yaml con tus IPs y credenciales
echo 3. Ejecuta: powershell -ExecutionPolicy Bypass -File install-service.ps1
echo.
echo O usa directamente:
echo   service.exe install  - Instalar como servicio
echo   service.exe start    - Iniciar servicio
echo   service.exe stop     - Detener servicio
echo   service.exe status   - Ver estado
echo.
pause
