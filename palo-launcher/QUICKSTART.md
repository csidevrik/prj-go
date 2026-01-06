# Guía de Inicio Rápido - Palo Alto Launcher

## 📋 Instalación en 5 Pasos

### Paso 1: Instalar dependencias

```bash
# En Ubuntu/Debian:
sudo apt update
sudo apt install firefox brave-browser # o google-chrome-stable

# Verificar que tienes Rust instalado:
cargo --version

# Si no tienes Rust, instalarlo:
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source $HOME/.cargo/env
```

### Paso 2: Compilar e instalar

```bash
cd palo-launcher
./install.sh
```

Selecciona la opción que prefieras:
- **Opción 1**: Instalación global (requiere sudo)
- **Opción 2**: Instalación solo para tu usuario
- **Opción 3**: Solo compilar

### Paso 3: Configurar perfiles de Firefox

```bash
./setup-firefox-profiles.sh
```

Este script creará automáticamente los perfiles necesarios.

### Paso 4: Configurar la herramienta

```bash
palo-launcher config
```

Introduce:
- La IP de tu firewall Palo Alto
- Selecciona tu combinación de navegadores preferida

### Paso 5: Primera conexión

```bash
palo-launcher
```

**IMPORTANTE - Primera vez:**
1. Se abrirán dos navegadores
2. En cada uno, loguéate manualmente
3. **Marca "Recordar mis credenciales"** o similar
4. Las próximas veces ya estarás logueado automáticamente

## 🎯 Uso Diario

### Conexión rápida

```bash
palo-launcher
```

Selecciona la sección:
- Dashboard Principal
- Monitor
- Policies
- Objects
- Network
- Logs de Tráfico

### Ver configuración

```bash
palo-launcher show
```

### Reconfigurar

```bash
palo-launcher config
```

## 💡 Ejemplos de Flujo de Trabajo

### Ejemplo 1: Revisar y editar policies

```bash
# 1. Ejecutar palo-launcher
palo-launcher

# 2. Seleccionar "Policies"
[Selecciona opción 3]

# 3. Se abren dos navegadores:
#    - Firefox: Para editar policies
#    - Brave: Navega manualmente a Monitor para ver logs en tiempo real

# 4. Trabaja en ambos simultáneamente
```

### Ejemplo 2: Monitoreo de tráfico

```bash
palo-launcher
# Selecciona "Logs de Tráfico"

# Firefox muestra logs
# Brave usa para filtrar y buscar detalles específicos
```

### Ejemplo 3: Configuración de objetos

```bash
palo-launcher
# Selecciona "Objects"

# En un navegador: Address Objects
# En otro navegador: Service Objects
```

## 🔧 Solución Rápida de Problemas

### "No se encuentran los perfiles"

```bash
# Recrear perfiles
./setup-firefox-profiles.sh
```

### "Se cierra la sesión automáticamente"

**Causa**: No marcaste "Recordar credenciales" al loguear

**Solución**:
1. Abre el navegador manualmente: `firefox -P palo-firefox`
2. Ve a la IP de Palo Alto
3. Loguéate y MARCA "Recordar credenciales"
4. Cierra y vuelve a usar `palo-launcher`

### "Brave/Chrome no abre"

**Verifica la instalación**:
```bash
brave-browser --version
# o
google-chrome --version
```

**Si no está instalado**:
```bash
# Brave
sudo apt install brave-browser

# Chrome
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

### "Los perfiles se mezclan con mi navegador personal"

**No debería pasar**. Los perfiles son completamente independientes.

**Verificar**:
```bash
# Ver perfiles de Firefox
ls ~/.mozilla/firefox/
# Deberías ver: xxxxx.palo-firefox y xxxxx.palo-firefox-2

# Ver perfiles de Brave
ls ~/.config/BraveSoftware/Brave-Browser/
# Deberías ver: palo-brave
```

## 🚀 Tips Pro

### 1. Crear un alias

```bash
echo 'alias palo="palo-launcher"' >> ~/.bashrc
source ~/.bashrc

# Ahora solo escribe:
palo
```

### 2. Atajos de teclado del firewall

Una vez abierto, usa estos atajos en el GUI de Palo Alto:
- `Ctrl + S`: Guardar cambios
- `Ctrl + F`: Buscar
- `Ctrl + Shift + C`: Commit
- `?`: Ayuda de atajos

### 3. Mantener múltiples sesiones

Puedes tener TRES o MÁS navegadores abiertos:
```bash
# Edita el config
nano ~/.config/palo-launcher/config.json

# Cambia default_browsers a:
"default_browsers": ["firefox", "firefox2", "brave"]
```

### 4. Bookmarks organizados

En cada perfil, crea bookmarks para acceso instantáneo:
- Firefox: Sections de configuración
- Brave: Logs y monitoreo

### 5. Personaliza las secciones

Edita `src/main.rs` para añadir URLs específicas que uses frecuentemente.

## 📱 Integración con tu Workflow

### Con tmux

```bash
# Crear sesión tmux para Palo Alto
tmux new-session -s palo
palo-launcher
# Ctrl+B D para detach
```

### Con i3wm / sway

```bash
# Añade a tu config:
bindsym $mod+p exec palo-launcher
assign [title="Palo Alto"] workspace 9
```

### Con scripts personalizados

```bash
# Crear script de trabajo diario
cat > ~/bin/daily-palo.sh << 'EOF'
#!/bin/bash
echo "Iniciando workflow diario de Palo Alto..."
palo-launcher
# Otros comandos que uses...
EOF

chmod +x ~/bin/daily-palo.sh
```

## 📚 Próximos Pasos

1. ✅ Lee el `README.md` completo para entender todas las funcionalidades
2. ✅ Revisa `ADVANCED.md` para configuraciones avanzadas
3. ✅ Personaliza según tu workflow específico
4. ✅ Comparte mejoras con el equipo

## 🆘 Ayuda Adicional

Si tienes problemas:
1. Revisa los logs: `palo-launcher 2>&1 | tee debug.log`
2. Verifica los perfiles de navegador existen
3. Confirma que puedes acceder al firewall manualmente
4. Revisa el archivo de configuración: `~/.config/palo-launcher/config.json`

---

**¡Ya estás listo para administrar Palo Alto de forma más eficiente! 🎉**
