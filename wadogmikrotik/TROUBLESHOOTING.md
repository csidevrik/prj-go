# Guía de Troubleshooting

## Problemas Comunes durante Instalación

### Problema: "Go no está instalado"

**Síntoma**: `go: command not found` o similar

**Solución**:
1. Descarga Go desde https://golang.org/dl/
2. Instala la versión 1.23 o superior
3. Reinicia PowerShell/CMD
4. Verifica: `go version`

---

### Problema: "Archivo de configuración no encontrado"

**Síntoma**: 
```
Error cargando configuración: no se puede leer config configs/config.yaml
```

**Solución**:
1. Verifica que estás en el directorio correcto:
   ```powershell
   pwd  # o cd si necesitas cambiar
   ```
2. Copia el archivo de ejemplo:
   ```powershell
   Copy-Item configs/config.example.yaml configs/config.yaml
   ```
3. Edita `configs/config.yaml` con tus valores

---

### Problema: Error al compilar con "missing go.sum entry"

**Síntoma**:
```
missing go.sum entry for module providing package...
```

**Solución**:
```powershell
go mod tidy
go mod download
```

---

## Problemas después de Instalar el Servicio

### Problema: Servicio se detiene inmediatamente

**Síntomas**:
- El servicio aparece como "Stopped" en Servicios
- Logs vacíos o muy cortos

**Diagnóstico**:
```powershell
# Ver logs
Get-Content logs/wadogmikrotik.log

# Correr en modo interactivo para ver errores
.\service.exe
```

**Causas comunes y soluciones**:

1. **Archivo YAML inválido**
   - Verifica sintaxis YAML: alineación correcta, sin tabs
   - Prueba en un validador online: https://www.yamllint.com/
   - Ejemplo correcto:
   ```yaml
   side_a:
     ip: "192.168.1.1"
     ssh:
       user: "wadogservice"
       password: "tu_password"
   ```

2. **Configuración incompleta**
   - Verifica que todas las IPs estén presentes
   - Verifica que `internet_targets` no está vacío
   - Revisa el ejemplo en `configs/config.example.yaml`

3. **Permisos del archivo config.yaml**
   - Asegúrate de que el servicio puede leer el archivo
   - Verifica permisos: clic derecho → Propiedades → Seguridad

---

### Problema: Error "Unable to authenticate"

**Síntoma**:
```
ssh: unable to authenticate, attempted methods [password], no supported methods remain
```

**Causas y soluciones**:

1. **Usuario o contraseña incorrectos**
   ```powershell
   # Verifica credenciales en config.yaml
   # Conecta manualmente para probar:
   ssh wadogservice@192.168.1.1
   ```

2. **Usuario no existe en MikroTik**
   ```routeros
   # En MikroTik, verifica:
   /user print
   
   # Si no existe, crea:
   /user add name=wadogservice password=mi_password group=full
   ```

3. **SSH deshabilitado**
   ```routeros
   # En MikroTik:
   /ip service enable ssh
   /ip service print  # Verifica que SSH está enabled
   ```

4. **Firewall bloqueando SSH**
   - Verifica que puerto 22 está abierto desde el servidor
   - En MikroTik: `/ip firewall filter print` y revisa reglas

---

### Problema: Error "No connection could be made"

**Síntoma**:
```
dial tcp 192.168.1.1:22: connectex: No connection could be made
```

**Causas y soluciones**:

1. **IP incorrecta**
   - Verifica que la IP en config.yaml es correcta
   - Prueba pingear: `ping 192.168.1.1`

2. **Antena no accesible**
   ```powershell
   # Verifica conectividad
   Test-NetConnection -ComputerName 192.168.1.1 -Port 22
   ```

3. **SSH no está escuchando en puerto 22**
   ```routeros
   /ip service print  # Verifica puerto 22
   /ip service set ssh disabled=no port=22
   ```

4. **Firewall de Windows bloqueando**
   - Windows Defender/Firewall puede estar bloqueando
   - Abre una excepción para el ejecutable

---

### Problema: Ping fallando pero SSH funciona

**Síntoma**:
```
Error haciendo ping a 8.8.8.8: sendto udp4 0.0.0.0:0->8.8.8.8:0: i/o timeout
```

**Posibles causas**:
1. Firewall del ISP bloqueando ICMP
2. Rutas de red incorrectas
3. DNS resuelto pero red sin acceso

**Solución**:
- Cambiar `internet_targets` a puertos HTTP:
  ```yaml
  internet_targets:
    - "8.8.8.8"      # DNS Google
    - "1.1.1.1"      # DNS Cloudflare
    - "208.67.222.222" # OpenDNS
  ```
- O aumentar `ping_timeout`:
  ```yaml
  ping_timeout: "10s"  # Cambiar de 5s a 10s
  ```

---

### Problema: El servicio se reinicia demasiado o muy poco

**Si se reinicia demasiado**:
- Aumenta `recovery_cooldown`:
  ```yaml
  recovery_cooldown: "600s"  # 10 minutos en lugar de 5
  ```
- Revisa si realmente hay pérdida de Internet

**Si no se reinicia cuando debería**:
- Verifica que `internet_targets` realmente no responden
- Reduce `ping_timeout` para detección más rápida:
  ```yaml
  ping_timeout: "3s"  # Más sensible
  ```
- Aumenta `retry_attempts` para ser más persistente:
  ```yaml
  retry_attempts: 3
  ```

---

### Problema: Alto consumo de CPU

**Síntoma**: El proceso `service.exe` usa mucho CPU constantemente

**Soluciones**:
1. Aumenta `monitor_interval`:
   ```yaml
   monitor_interval: "30s"  # de 5s a 30s
   ```

2. Reduce `retry_attempts`:
   ```yaml
   retry_attempts: 1  # de 2 a 1
   ```

3. Aumenta `retry_delay`:
   ```yaml
   retry_delay: "5s"  # de 2s a 5s
   ```

Después de cambiar config.yaml, reinicia:
```powershell
.\install-service.ps1 -Action restart
```

---

### Problema: Archivos de log muy grandes

**Síntoma**: `logs/wadogmikrotik.log` crece constantemente

**Soluciones**:

1. **Los logs rotan automáticamente**, pero puedes ajustar:
   ```yaml
   log_max_size: 5        # Reducir de 10 MB
   log_max_backups: 3     # Reducir de 5
   log_max_age: 7         # Reducir de 30 días
   ```

2. **Reduce nivel de detalle** (cambiar DEBUG a INFO):
   - Actualmente está en INFO, el cual es normal
   - DEBUG logs son por defecto deshabilitados

3. **Limpiar manualmente**:
   ```powershell
   # Detener servicio
   .\install-service.ps1 -Action stop
   
   # Eliminar logs viejos
   Remove-Item logs\wadogmikrotik.log*
   
   # Reiniciar
   .\install-service.ps1 -Action start
   ```

---

## Problemas en Tiempo de Ejecución

### Problema: La recuperación no se ejecuta

**Síntomas**:
- Ve "Pérdida de conectividad detectada" en logs
- Pero no ve "Comando /system reboot enviado"

**Diagnóstico**:
```powershell
# Ver logs
Get-Content logs/wadogmikrotik.log | Select-String "Pérdida", "Diagnóstico", "Lado A", "Lado B"
```

**Causas comunes**:

1. **Antena no responde al ping**
   ```
   INFO: Resultado del ping de Lado A ... reachable=false
   WARN: Lado A no responde al ping
   ```
   - Verifica que la antena está activa
   - Prueba ping manual: `ping 192.168.1.1`

2. **En período de cooldown**
   ```
   INFO: Esperando cooldown antes del nuevo intento
   ```
   - Normal, espera a que termine cooldown
   - Reduce `recovery_cooldown` si es demasiado largo

3. **Error SSH que no se muestra**
   - Aumenta verbosidad ejecutando en modo interactivo:
   ```powershell
   .\service.exe
   ```

---

### Problema: Antenas reinician demasiadas veces

**Síntomas**: Las antenas se reinician cada pocos segundos

**Causas**:
1. El servicio está reiniiciando porque Internet sigue caído
2. `recovery_cooldown` es demasiado corto

**Solución**:
1. Aumenta `recovery_cooldown`:
   ```yaml
   recovery_cooldown: "900s"  # 15 minutos
   ```

2. Identifica la causa real:
   - ¿Es un problema del proveedor de Internet?
   - ¿Es el enlace inalámbrico entre antenas?
   - Revisa logs de las antenas MikroTik directamente

---

### Problema: Antenas se reinician incluso sin pérdida de Internet

**Síntomas**: Ve "Comando /system reboot enviado" sin haber detectado pérdida

**Posibles causas**:
1. `internet_targets` son inaccesibles (firewall, DNS bloqueado)
   - Prueba con direcciones IP en lugar de nombres
   - Usa múltiples targets (Google, Cloudflare, OpenDNS)

2. Problema intermitente de conectividad
   - Aumenta `retry_attempts` y `retry_delay`

**Diagnóstico**:
```powershell
# Prueba conectividad manual
ping 8.8.8.8
ping 1.1.1.1

# Si ninguno responde pero hay Internet:
# Usa direcciones IP diferentes o agrega puertos HTTP
```

---

## Problemas de Permiso

### Problema: "Access denied" al instalar servicio

**Síntoma**:
```
Error: Access denied
```

**Solución**:
1. Abre PowerShell **como Administrador**
   - Clic derecho en PowerShell → Ejecutar como administrador
   - O: Win + X, luego selecciona PowerShell (Admin)

2. Comprueba permisos:
   ```powershell
   # Esto debe devolver "Administrator"
   whoami /groups | Select-String "Administrator"
   ```

---

### Problema: "Access denied" al leer config.yaml

**Síntoma**: El servicio no puede leer configuración

**Solución**:
1. Verifica permisos del archivo:
   ```powershell
   Get-Acl configs/config.yaml
   ```

2. Si es necesario, resetea permisos:
   ```powershell
   # Otorga permiso al usuario actual
   $acl = Get-Acl configs/config.yaml
   $acl.SetAccessRuleProtection($false, $false)
   $acl | Set-Acl configs/config.yaml
   ```

---

## Debugging Avanzado

### Ejecutar en modo interactivo con salida detallada

```powershell
# Detiene el servicio primero
.\install-service.ps1 -Action stop

# Ejecuta en modo interactivo (verás todos los logs en la consola)
.\service.exe

# Presiona Ctrl+C para detener
```

### Capturar salida de error completa

```powershell
# Ejecutar y guardar en archivo
.\service.exe 2>&1 | Tee-Object debug.log

# Después examina el archivo
Get-Content debug.log
```

### Monitorear logs en tiempo real

```powershell
# Ver logs a medida que se escriben
Get-Content logs/wadogmikrotik.log -Wait -Tail 50

# O usando PowerShell:
Get-Content logs/wadogmikrotik.log -Tail 50 -Wait
```

---

## Verificación de Conectividad

### Script para verificar acceso SSH

```powershell
# Crear archivo test-ssh.ps1
$servers = @("192.168.1.1", "192.168.1.2")
$user = "wadogservice"
$password = "your_password"

foreach ($server in $servers) {
    Write-Host "Probando $server..."
    
    # Verificar ping
    $ping = Test-Connection -ComputerName $server -Count 1 -Quiet
    Write-Host "  Ping: $(if($ping) {'OK'} else {'FAILED'})"
    
    # Verificar SSH
    $ssh = Test-NetConnection -ComputerName $server -Port 22
    Write-Host "  SSH: $(if($ssh.TcpTestSucceeded) {'OK'} else {'FAILED'})"
}
```

### Script para verificar Internet targets

```powershell
$targets = @("8.8.8.8", "1.1.1.1", "208.67.222.222")

foreach ($target in $targets) {
    $ping = Test-Connection -ComputerName $target -Count 1 -Quiet
    Write-Host "$target : $(if($ping) {'✓'} else {'✗'})"
}
```

---

## Contacto y Soporte

Si ninguna de estas soluciones funciona:

1. Recopila información:
   ```powershell
   # Copia esto en un archivo para reportar
   Get-Content logs/wadogmikrotik.log -Tail 100
   Get-Content configs/config.yaml | Select-String -NotMatch "password"
   ```

2. Verifica documentación:
   - [README.md](README.md) - Descripción general
   - [DEPLOYMENT.md](DEPLOYMENT.md) - Guía de instalación
   - [EXAMPLE_LOG.md](EXAMPLE_LOG.md) - Ejemplos de logs

3. Reporta el problema con:
   - Versión de Go (`go version`)
   - Versión de Windows (`winver`)
   - Logs completos (sin exponer credenciales)
   - Descripción exacta del problema
