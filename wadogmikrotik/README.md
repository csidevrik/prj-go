# WadoG MikroTik Recovery Service

Servicio de Windows en Go que monitoriza continuamente la conectividad a Internet y ejecuta de forma automática un proceso de recuperación cuando detecta una pérdida de conexión causada por un bloqueo del enlace inalámbrico.

## Características

- ✅ Monitoreo continuo de conectividad a Internet
- ✅ Diagnóstico automático de antenas MikroTik por ICMP (Ping)
- ✅ Recuperación automática mediante SSH
- ✅ Prevención de ciclos de reinicio infinito
- ✅ Registro detallado y rotativo
- ✅ Configuración externa (YAML)
- ✅ Bajo consumo de recursos

## Arquitectura

```
cmd/service/             - Punto de entrada principal
internal/
  ├── config/           - Carga y validación de configuración
  ├── logger/           - Sistema de logging rotativo
  ├── monitor/          - Ciclo continuo de monitoreo
  ├── network/          - Lógica de ICMP ping
  ├── ssh/              - Conexiones SSH a MikroTik
  ├── recovery/         - Diagnóstico y proceso de recuperación
  └── windows/          - Integración con servicios de Windows
configs/                - Archivos de configuración (YAML)
logs/                   - Registros rotativos
```

## Requisitos

- Go 1.23 o superior
- Windows XP SP3 o superior (para ejecutar como servicio)
- Conectividad SSH a las antenas MikroTik
- Usuario SSH con permiso para ejecutar `/system reboot`

## Instalación

### 1. Clonar o descargar el proyecto

```bash
cd wadogmikrotik
```

### 2. Descargar dependencias

```bash
go mod download
```

### 3. Compilar

```bash
go build -o wadogmikrotik.exe ./cmd/service
```

### 4. Configurar

Copia `configs/config.example.yaml` a `configs/config.yaml`:

```bash
cp configs/config.example.yaml configs/config.yaml
```

Edita `configs/config.yaml` con las IPs y credenciales de tus antenas MikroTik.

## Uso

### Ejecutar en modo interactivo (pruebas):

```powershell
.\wadogmikrotik.exe
```

### Instalar como servicio de Windows:

```powershell
.\wadogmikrotik.exe install
.\wadogmikrotik.exe start
```

### Ver estado del servicio:

```powershell
.\wadogmikrotik.exe status
```

### Detener el servicio:

```powershell
.\wadogmikrotik.exe stop
```

### Desinstalar el servicio:

```powershell
.\wadogmikrotik.exe uninstall
```

## Configuración (config.yaml)

```yaml
side_a:
  ip: "192.168.1.1"
  ssh:
    user: "wadogservice"
    password: "contraseña"

side_b:
  ip: "192.168.1.2"
  ssh:
    user: "wadogservice"
    password: "contraseña"

internet_targets:
  - "8.8.8.8"
  - "1.1.1.1"

monitor_interval: "5s"        # Verificar conectividad cada 5 segundos
recovery_cooldown: "300s"     # Esperar 5 min después de recuperación
ping_timeout: "5s"            # Timeout para cada ping
retry_attempts: 2             # Reintentos por ping
retry_delay: "2s"             # Delay entre reintentos
```

## Flujo de Operación

1. **Monitoreo**: Cada `monitor_interval`, verifica conectividad a `internet_targets`
2. **Pérdida detectada**: Si no hay respuesta en ningún destino
3. **Diagnóstico**:
   - Ping al Lado B → Si no responde, abortar
   - Ping al Lado A → Si no responde, abortar
4. **Recuperación** (si ambas responden):
   - Conectar SSH al Lado A → Enviar `/system reboot`
   - Conectar SSH al Lado B → Enviar `/system reboot`
   - Activar `recovery_cooldown`
5. **Retorno a monitoreo**

## Logging

Los logs se escriben en `logs/wadogmikrotik.log` con rotación automática.

Ejemplos de eventos registrados:
- Inicio/parada del servicio
- Estado de conectividad
- Resultados de pings
- Recuperaciones ejecutadas
- Errores SSH/autenticación

## Seguridad

- ✅ No se codifican credenciales en el código
- ✅ Usuario SSH dedicado con permisos mínimos
- ✅ Timeout en conexiones SSH
- ✅ Cooldown configurable evita ciclos infinitos
- ✅ Logs de auditoría completos
