# Guía de Instalación y Deployment

## Requisitos Previos

- Windows XP SP3 o superior (Windows 10/11 recomendado)
- Go 1.23 o superior (solo si compilas desde el código)
- Conectividad SSH a las antenas MikroTik
- Usuario SSH con permisos para ejecutar `/system reboot`
- PowerShell 3.0 o superior (para usar los scripts)

## Configuración de MikroTik

Antes de instalar el servicio, configura un usuario dedicado en cada antena MikroTik:

### En MikroTik (Lado A y Lado B):

```routeros
# Crear usuario
/user add name=wadogservice password=your_secure_password group=full

# Verificar que SSH está habilitado
/ip service enable ssh
```

> **Nota de Seguridad**: Idealmente, crea un usuario con permisos limitados solo a `/system reboot`.

## Compilación

### Opción 1: Usar el script (Recomendado)

Desde Windows, ejecuta:

```batch
build.bat
```

O desde PowerShell:

```powershell
.\build.bat
```

### Opción 2: Compilar manualmente

```powershell
# Descargar dependencias
go mod download
go mod tidy

# Compilar
go build -o service.exe ./cmd/service
```

El ejecutable `service.exe` estará en el directorio raíz del proyecto.

## Configuración del Servicio

### 1. Crear archivo de configuración

```bash
# Copiar el archivo de ejemplo
Copy-Item configs/config.example.yaml configs/config.yaml
```

### 2. Editar `configs/config.yaml`

Abre el archivo con tu editor de texto favorito y configura:

```yaml
side_a:
  ip: "192.168.1.1"      # IP real de tu antena A
  ssh:
    user: "wadogservice"
    password: "tu_password_segura"

side_b:
  ip: "192.168.1.2"      # IP real de tu antena B
  ssh:
    user: "wadogservice"
    password: "tu_password_segura"

internet_targets:
  - "8.8.8.8"            # Google DNS (prueba primero)
  - "1.1.1.1"            # Cloudflare DNS
  - "208.67.222.222"     # OpenDNS (opcional)

monitor_interval: "5s"       # Revisar cada 5 segundos
recovery_cooldown: "300s"    # Esperar 5 min después de recuperar
ping_timeout: "5s"
retry_attempts: 2
retry_delay: "2s"
```

### 3. Verificar conectividad

Antes de instalar el servicio, verifica que puedas hacer ping a las antenas:

```powershell
# Desde Windows, verifica conectividad SSH
Test-NetConnection -ComputerName 192.168.1.1 -Port 22
Test-NetConnection -ComputerName 192.168.1.2 -Port 22
```

## Instalación Como Servicio de Windows

### Opción 1: Usar el script PowerShell (Recomendado)

Abre PowerShell **como Administrador** y ejecuta:

```powershell
# Permitir ejecución de scripts (si es necesario)
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope CurrentUser

# Instalar
.\install-service.ps1 -Action install

# Iniciar
.\install-service.ps1 -Action start

# Ver estado
.\install-service.ps1 -Action status
```

### Opción 2: Línea de comandos directa

Abre **CMD o PowerShell como Administrador** en el directorio del proyecto:

```powershell
# Instalar
.\service.exe install

# Iniciar
.\service.exe start

# Ver estado
.\service.exe status

# Detener
.\service.exe stop

# Desinstalar
.\service.exe uninstall
```

## Verificación de la Instalación

### Verificar que el servicio está instalado

Abre **Administrador de Servicios**:

```powershell
# Desde PowerShell
Get-Service WadoGmikrotik
```

O en la GUI:
1. Presiona `Win + R`
2. Escribe `services.msc`
3. Busca "WadoG MikroTik Recovery Service"

El servicio debe aparecer con estado **Running**.

### Verificar los logs

Los logs se escriben en `logs/wadogmikrotik.log`:

```powershell
# Ver últimas 50 líneas
Get-Content logs/wadogmikrotik.log -Tail 50

# Ver en tiempo real (PowerShell 3.0+)
Get-Content logs/wadogmikrotik.log -Wait -Tail 50
```

Busca líneas como:
- `Iniciando servicio de monitoreo` - El servicio está funcionando
- `Conectividad a Internet verificada` - Internet OK
- `Pérdida de conectividad a Internet detectada` - Detectó pérdida
- `Ambas antenas responden, ejecutando recuperación remota` - Recuperación en curso

## Pruebas

### Prueba 1: Modo interactivo

Ejecuta el servicio en modo interactivo para ver los logs en tiempo real:

```powershell
.\service.exe
```

Verás un output como:

```
INFO[2024-07-26T14:30:00Z] Iniciando servicio de monitoreo de conectividad
INFO[2024-07-26T14:30:05Z] Conectividad a Internet verificada
INFO[2024-07-26T14:30:10Z] Conectividad a Internet verificada
```

Presiona `Ctrl+C` para detener.

### Prueba 2: Simular pérdida de Internet

1. Desconecta temporalmente tu internet
2. Monitorea los logs:

```powershell
Get-Content logs/wadogmikrotik.log -Tail 50
```

Deberías ver:
```
WARN[...] Pérdida de conectividad a Internet detectada
INFO[...] Iniciando diagnóstico de antenas MikroTik
INFO[...] Resultado del ping de Lado B
INFO[...] Resultado del ping de Lado A
INFO[...] Ambas antenas responden, ejecutando recuperación remota
```

### Prueba 3: Verificar cooldown

1. Simula otra pérdida de Internet
2. El servicio no debe ejecutar recuperación si está dentro del `recovery_cooldown`
3. Verás: `Esperando cooldown antes del nuevo intento de recuperación`

## Troubleshooting

### El ejecutable no se crea

```powershell
# Verifica que Go está instalado
go version

# Verifica que estás en el directorio correcto
cd "ruta\a\wadogmikrotik"

# Intenta compilar manualmente
go build -o service.exe ./cmd/service
```

### Error: "Archivo de configuración no encontrado"

```powershell
# Verifica que exists el archivo
Test-Path configs/config.yaml

# Si no existe, cópialo
Copy-Item configs/config.example.yaml configs/config.yaml
```

### El servicio se detiene inmediatamente

Revisa los logs:

```powershell
Get-Content logs/wadogmikrotik.log
```

Errores comunes:
- `error conectando por SSH` - Verifica IP y conectividad
- `error autenticando` - Verifica usuario/password
- `error parseando YAML` - Verifica formato de config.yaml

### El servicio no llama a `/system reboot`

Verifica:
1. Usuario SSH tiene permisos de reboot
2. El formato del comando es exacto: `/system reboot`
3. Los logs muestran `Ambas antenas responden, ejecutando recuperación remota`

### El servicio usa mucha CPU

Aumenta `monitor_interval`:

```yaml
monitor_interval: "30s"  # Cambiar de 5s a 30s
```

Luego reinicia el servicio:

```powershell
.\install-service.ps1 -Action restart
```

## Mantenimiento

### Limpiar logs antiguos

Los logs se rotan automáticamente según la configuración, pero puedes limpiar manualmente:

```powershell
# Detener el servicio primero
.\service.exe stop

# Eliminar logs
Remove-Item logs/*.log

# Reiniciar
.\service.exe start
```

### Actualizar configuración

1. Edita `configs/config.yaml`
2. Reinicia el servicio:

```powershell
.\install-service.ps1 -Action restart
```

### Desinstalar completamente

```powershell
# Detener
.\install-service.ps1 -Action stop

# Desinstalar
.\install-service.ps1 -Action uninstall

# Eliminar archivos (opcional)
Remove-Item service.exe
Remove-Item logs
```

## Monitoreo Remoto

Para monitorear el servicio desde otro equipo:

```powershell
# Ver estado remoto
Invoke-Command -ComputerName nombre_equipo -ScriptBlock {
    Get-Service WadoGmikrotik
}

# Ver logs remotos
Invoke-Command -ComputerName nombre_equipo -ScriptBlock {
    Get-Content "C:\ruta\a\wadogmikrotik\logs\wadogmikrotik.log" -Tail 50
}
```

## Seguridad

✓ **Buenas prácticas aplicadas:**
- No se codifican credenciales en el código
- Usuario SSH dedicado con permisos mínimos
- Timeouts en todas las conexiones SSH
- Logs detallados para auditoría
- Cooldown previene ciclos de reinicio
- Verificación de conectividad antes de reboot

⚠ **Consideraciones adicionales:**
- Usa claves SSH en lugar de contraseñas si es posible
- Restringe permisos del archivo `configs/config.yaml` a solo lectura
- Ejecuta el servicio con cuenta de usuario, no SYSTEM
- Revisa logs regularmente para detectar patrones anormales

## Soporte

Para reportar problemas:
1. Revisa los logs en `logs/wadogmikrotik.log`
2. Verifica la conectividad SSH manualmente
3. Prueba en modo interactivo: `.\service.exe`
4. Consulta la sección Troubleshooting arriba
