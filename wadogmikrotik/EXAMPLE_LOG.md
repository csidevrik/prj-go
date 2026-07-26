# Ejemplo de Log de Ejecución Normal

Este documento muestra ejemplos de logs que verás en `logs/wadogmikrotik.log` en diferentes situaciones.

## Inicio Normal del Servicio

```
time="2024-07-26T14:30:00Z" level=info msg="Iniciando servicio de monitoreo de conectividad"
time="2024-07-26T14:30:05Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:10Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:15Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:20Z" level=debug msg="Conectividad a Internet verificada"
```

**Interpretación**: El servicio está ejecutándose normalmente. Verifica conectividad cada 5 segundos y la Internet está disponible.

---

## Pérdida de Conectividad Detectada (Escenario Exitoso)

```
time="2024-07-26T14:30:20Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:25Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:30Z" level=warn msg="Pérdida de conectividad a Internet detectada"
time="2024-07-26T14:30:30Z" level=info msg="Iniciando diagnóstico de antenas MikroTik"
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado B" side=B ip=192.168.1.2 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado A" side=A ip=192.168.1.1 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Ambas antenas responden, ejecutando recuperación remota"
time="2024-07-26T14:30:30Z" level=info msg="Enviando comando de reinicio" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=info msg="Comando /system reboot enviado" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=info msg="Enviando comando de reinicio" side="Lado B" ip=192.168.1.2
time="2024-07-26T14:30:31Z" level=info msg="Comando /system reboot enviado" side="Lado B" ip=192.168.1.2
time="2024-07-26T14:30:31Z" level=info msg="Recuperación realizada y cooldown activado"
time="2024-07-26T14:30:36Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:41Z" level=debug msg="Conectividad a Internet verificada"
```

**Interpretación**: 
1. Internet perdido a las 14:30:30
2. Se iniciaron diagnósticos
3. Ambas antenas respondieron al ping
4. Se enviaron comandos de reboot a Lado A y Lado B
5. El servicio vuelve a monitoreo normal
6. Internet se recupera después del reboot

---

## Fallo de Conectividad en Antena (Sin Recuperación)

```
time="2024-07-26T14:30:30Z" level=warn msg="Pérdida de conectividad a Internet detectada"
time="2024-07-26T14:30:30Z" level=info msg="Iniciando diagnóstico de antenas MikroTik"
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado B" side=B ip=192.168.1.2 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado A" side=A ip=192.168.1.1 reachable=false
time="2024-07-26T14:30:31Z" level=warn msg="Error durante diagnóstico/recuperación: Lado A no responde al ping"
```

**Interpretación**:
1. Internet perdido
2. Diagnóstico iniciado
3. Lado B respondió correctamente
4. Lado A NO respondió al ping
5. Recuperación ABORTADA (prevención de ejecutar reboot en equipos inaccesibles)
6. Posible problema: Lado A caído, conectividad SSH interrumpida, problema en la red local

---

## Período de Cooldown Activo

```
time="2024-07-26T14:30:31Z" level=info msg="Recuperación realizada y cooldown activado"
time="2024-07-26T14:30:36Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:34:00Z" level=warn msg="Pérdida de conectividad a Internet detectada"
time="2024-07-26T14:34:00Z" level=info msg="Esperando cooldown antes del nuevo intento de recuperación"
time="2024-07-26T14:34:05Z" level=debug msg="Conectividad a Internet verificada"
```

**Interpretación**:
1. Recuperación ejecutada a las 14:30:31
2. Segunda pérdida de Internet a las 14:34:00 (dentro de la ventana cooldown)
3. El servicio NO ejecuta recuperación (espera el período cooldown)
4. Esto previene ciclos de reinicio infinito

---

## Error SSH - Autenticación Fallida

```
time="2024-07-26T14:30:30Z" level=warn msg="Pérdida de conectividad a Internet detectada"
time="2024-07-26T14:30:30Z" level=info msg="Iniciando diagnóstico de antenas MikroTik"
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado B" side=B ip=192.168.1.2 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado A" side=A ip=192.168.1.1 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Ambas antenas responden, ejecutando recuperación remota"
time="2024-07-26T14:30:30Z" level=info msg="Enviando comando de reinicio" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=error msg="Error durante diagnóstico/recuperación: falló reinicio Lado A: error conectando por SSH a 192.168.1.1: ssh: unable to authenticate, attempted methods [password], no supported methods remain"
```

**Interpretación**:
- Error de autenticación SSH
- Posibles causas:
  - Usuario/password incorrectos en config.yaml
  - Usuario no existe en la antena
  - Contraseña cambió
  - SSH deshabilitado en la antena

---

## Error SSH - Conexión Rechazada

```
time="2024-07-26T14:30:30Z" level=warn msg="Pérdida de conectividad a Internet detectada"
time="2024-07-26T14:30:30Z" level=info msg="Iniciando diagnóstico de antenas MikroTik"
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado B" side=B ip=192.168.1.2 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado A" side=A ip=192.168.1.1 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Ambas antenas responden, ejecutando recuperación remota"
time="2024-07-26T14:30:30Z" level=info msg="Enviando comando de reinicio" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=error msg="Error durante diagnóstico/recuperación: falló reinicio Lado A: error conectando por SSH a 192.168.1.1: dial tcp 192.168.1.1:22: connectex: No connection could be made because the target machine actively refused it."
```

**Interpretación**:
- SSH puerto 22 no está disponible
- Posibles causas:
  - SSH deshabilitado en la antena
  - IP incorrecta
  - Firewall bloqueando puerto 22
  - Antena apagada o reiniciando

---

## Ejemplo Completo de una Recuperación

```
# 14:30 - Servicio inicia
time="2024-07-26T14:30:00Z" level=info msg="Servicio iniciado"
time="2024-07-26T14:30:00Z" level=info msg="Iniciando servicio de monitoreo de conectividad"

# 14:30:00-14:30:25 - Monitoreo normal
time="2024-07-26T14:30:05Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:10Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:15Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:20Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:25Z" level=debug msg="Conectividad a Internet verificada"

# 14:30:30 - PÉRDIDA DETECTADA
time="2024-07-26T14:30:30Z" level=warn msg="Pérdida de conectividad a Internet detectada"

# 14:30:30 - DIAGNÓSTICO
time="2024-07-26T14:30:30Z" level=info msg="Iniciando diagnóstico de antenas MikroTik"
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado B" side=B ip=192.168.1.2 reachable=true
time="2024-07-26T14:30:30Z" level=info msg="Resultado del ping de Lado A" side=A ip=192.168.1.1 reachable=true

# 14:30:30 - RECUPERACIÓN
time="2024-07-26T14:30:30Z" level=info msg="Ambas antenas responden, ejecutando recuperación remota"
time="2024-07-26T14:30:30Z" level=info msg="Enviando comando de reinicio" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=info msg="Comando /system reboot enviado" side="Lado A" ip=192.168.1.1
time="2024-07-26T14:30:31Z" level=info msg="Enviando comando de reinicio" side="Lado B" ip=192.168.1.2
time="2024-07-26T14:30:31Z" level=info msg="Comando /system reboot enviado" side="Lado B" ip=192.168.1.2
time="2024-07-26T14:30:31Z" level=info msg="Recuperación realizada y cooldown activado"

# 14:30:35-14:30:50 - Esperar a que antenas se reinicien
time="2024-07-26T14:30:35Z" level=debug msg="Error haciendo ping a 8.8.8.8: sendto udp4 0.0.0.0:0->8.8.8.8:0: i/o timeout"
time="2024-07-26T14:30:40Z" level=debug msg="Error haciendo ping a 8.8.8.8: sendto udp4 0.0.0.0:0->8.8.8.8:0: i/o timeout"
time="2024-07-26T14:30:45Z" level=debug msg="Error haciendo ping a 8.8.8.8: sendto udp4 0.0.0.0:0->8.8.8.8:0: i/o timeout"

# 14:30:50 - INTERNET RECUPERADO
time="2024-07-26T14:30:50Z" level=debug msg="Conectividad a Internet verificada"
time="2024-07-26T14:30:55Z" level=debug msg="Conectividad a Internet verificada"

# 14:35 - Cooldown vence, servicio listo para nueva recuperación
```

---

## Niveles de Log

| Nivel | Color | Significado | Acción |
|-------|-------|------------|--------|
| **DEBUG** | Cyan | Información detallada, verbosa | Información, no requiere acción |
| **INFO** | Blanco | Eventos importantes, operación normal | Información, monitorear |
| **WARN** | Amarillo | Advertencia, algo inusual | Revisar, posible acción |
| **ERROR** | Rojo | Error, fallo en operación | **Acción requerida** |

---

## Análisis de Patrones

### Patrón A: Todo funciona correctamente
```
DEBUG: Conectividad verificada (cada 5 seg)
Acción: Ninguna, sistema OK
```

### Patrón B: Pérdida ocasional de Internet
```
WARN: Pérdida detectada
INFO: Ambas antenas responden
INFO: Recuperación realizada
DEBUG: Conectividad verificada (después)
Acción: Normal, monitoreo continuo
```

### Patrón C: Antena caída
```
WARN: Pérdida detectada
INFO: Lado B responde
INFO: Lado A NO responde
ERROR: Recuperación abortada
Acción: Revisar Lado A inmediatamente
```

### Patrón D: Error de autenticación
```
WARN: Pérdida detectada
INFO: Ambas antenas responden
ERROR: unable to authenticate
Acción: Revisar credenciales en config.yaml
```

### Patrón E: Múltiples intentos de recuperación
```
(Primer ciclo de recuperación)
(Segundo ciclo dentro de cooldown)
INFO: Esperando cooldown...
Acción: Normal, cooldown previniendo ciclos infinitos
```
