# Configuración Avanzada - Palo Alto Launcher

## 🎯 Configuraciones para Power Users

### 1. Múltiples Firewalls

Si administras varios firewalls Palo Alto, puedes crear múltiples configuraciones:

#### Opción A: Múltiples archivos de configuración

```bash
# Crear diferentes configuraciones
mkdir -p ~/.config/palo-launcher/

# Configuración para Producción
cat > ~/.config/palo-launcher/config-prod.json << 'EOF'
{
  "firewall_ip": "192.168.1.1",
  "firefox_profile": "palo-prod-firefox",
  "brave_profile": "palo-prod-brave",
  "chrome_profile": "palo-prod-chrome",
  "default_browsers": ["firefox", "brave"]
}
EOF

# Configuración para Desarrollo
cat > ~/.config/palo-launcher/config-dev.json << 'EOF'
{
  "firewall_ip": "192.168.2.1",
  "firefox_profile": "palo-dev-firefox",
  "brave_profile": "palo-dev-brave",
  "chrome_profile": "palo-dev-chrome",
  "default_browsers": ["firefox", "chrome"]
}
EOF

# Crear alias para cada entorno
echo 'alias palo-prod="cp ~/.config/palo-launcher/config-prod.json ~/.config/palo-launcher/config.json && palo-launcher"' >> ~/.bashrc
echo 'alias palo-dev="cp ~/.config/palo-launcher/config-dev.json ~/.config/palo-launcher/config.json && palo-launcher"' >> ~/.bashrc

source ~/.bashrc
```

#### Opción B: Script wrapper personalizado

```bash
# Crear script ~/bin/palo-multi.sh
cat > ~/bin/palo-multi.sh << 'EOF'
#!/bin/bash

echo "Selecciona el firewall:"
echo "1. Producción (192.168.1.1)"
echo "2. Desarrollo (192.168.2.1)"
echo "3. QA (192.168.3.1)"
read -p "Opción: " option

case $option in
    1) FW_IP="192.168.1.1" ;;
    2) FW_IP="192.168.2.1" ;;
    3) FW_IP="192.168.3.1" ;;
    *) echo "Opción inválida"; exit 1 ;;
esac

# Actualizar temporalmente el config
CONFIG_FILE=~/.config/palo-launcher/config.json
TEMP_CONFIG=$(mktemp)

jq --arg ip "$FW_IP" '.firewall_ip = $ip' "$CONFIG_FILE" > "$TEMP_CONFIG"
mv "$TEMP_CONFIG" "$CONFIG_FILE"

palo-launcher
EOF

chmod +x ~/bin/palo-multi.sh
```

### 2. Extensiones de Navegador Recomendadas

Para maximizar tu productividad con Palo Alto, considera instalar estas extensiones en tus perfiles:

#### Firefox (Perfil palo-firefox)

1. **JSON Viewer** - Para visualizar respuestas API
   ```
   https://addons.mozilla.org/en-US/firefox/addon/json-viewer/
   ```

2. **Copy PlainText** - Copiar logs sin formato
   ```
   https://addons.mozilla.org/en-US/firefox/addon/copy-plaintext/
   ```

3. **Tab Session Manager** - Guardar sesiones de trabajo
   ```
   https://addons.mozilla.org/en-US/firefox/addon/tab-session-manager/
   ```

#### Brave/Chrome (Perfil palo-brave)

1. **JSON Formatter**
   ```
   Chrome Web Store: JSON Formatter
   ```

2. **Session Buddy** - Gestión de pestañas
   ```
   Chrome Web Store: Session Buddy
   ```

### 3. Bookmarks Organizados por Perfil

Organiza tus bookmarks en cada perfil del navegador:

**Firefox - Perfil para Policies:**
```
Palo Alto - Policies/
├── Security Rules
├── NAT Rules
├── Application Override
├── DoS Protection
└── Security Profiles
    ├── Antivirus
    ├── Anti-Spyware
    ├── Vulnerability Protection
    └── URL Filtering
```

**Brave - Perfil para Monitoring:**
```
Palo Alto - Monitor/
├── Traffic Logs
├── Threat Logs
├── System Logs
├── ACC (Application Command Center)
├── Dashboard
└── Reports
```

### 4. Snippets de Configuración Útiles

#### Archivo de configuración completo con todas las opciones

```json
{
  "firewall_ip": "192.168.1.1",
  "firefox_profile": "palo-firefox",
  "brave_profile": "palo-brave",
  "chrome_profile": "palo-chrome",
  "default_browsers": ["firefox", "brave"],
  "custom_sections": {
    "logs": "#monitor/logs/traffic",
    "rules": "#policies/security/rules",
    "objects": "#objects/address",
    "commit": "#commit",
    "dashboard": "#monitor/dashboard"
  }
}
```

### 5. Scripts Complementarios

#### Auto-refresh para mantener sesiones activas

Crea este script para evitar timeout de sesión:

```bash
# ~/bin/palo-keepalive.sh
#!/bin/bash

# Refresca las pestañas de Palo Alto cada 5 minutos
# Requiere: xdotool

while true; do
    sleep 300  # 5 minutos
    
    # Buscar ventanas de Firefox con "Palo Alto"
    WINDOWS=$(xdotool search --name "Palo Alto" 2>/dev/null)
    
    for WIN in $WINDOWS; do
        xdotool windowactivate $WIN
        xdotool key F5
        sleep 1
    done
done
```

Uso:
```bash
chmod +x ~/bin/palo-keepalive.sh
nohup ~/bin/palo-keepalive.sh &
```

#### Backup de configuración de perfiles

```bash
# Backup de perfiles de Firefox
#!/bin/bash

BACKUP_DIR=~/backups/firefox-profiles/$(date +%Y%m%d)
mkdir -p "$BACKUP_DIR"

# Backup perfil palo-firefox
PROFILE_PATH=$(find ~/.mozilla/firefox -name "*palo-firefox*" -type d)
if [ -n "$PROFILE_PATH" ]; then
    tar -czf "$BACKUP_DIR/palo-firefox.tar.gz" "$PROFILE_PATH"
    echo "Backup creado: $BACKUP_DIR/palo-firefox.tar.gz"
fi

# Backup perfil palo-firefox-2
PROFILE_PATH_2=$(find ~/.mozilla/firefox -name "*palo-firefox-2*" -type d)
if [ -n "$PROFILE_PATH_2" ]; then
    tar -czf "$BACKUP_DIR/palo-firefox-2.tar.gz" "$PROFILE_PATH_2"
    echo "Backup creado: $BACKUP_DIR/palo-firefox-2.tar.gz"
fi
```

### 6. Integración con i3/sway (Window Managers)

Si usas un tiling window manager, puedes configurar layouts específicos:

#### i3 config
```
# ~/.config/i3/config

# Workspace dedicado para Palo Alto
assign [title="Palo Alto Networks"] workspace 9

# Atajos
bindsym $mod+p exec palo-launcher

# Layout automático para Palo Alto
for_window [title="Palo Alto Networks"] layout splitv
```

#### sway config
```
# ~/.config/sway/config

assign [title="Palo Alto Networks"] workspace 9
bindsym $mod+p exec palo-launcher
for_window [title="Palo Alto Networks"] layout splitv
```

### 7. Personalización de URLs

Puedes modificar el código fuente para añadir URLs personalizadas:

```rust
// En src/main.rs, en la función interactive_connect():

let sections = vec![
    "Dashboard Principal",
    "Monitor",
    "Policies",
    "Objects",
    "Network",
    "Logs de Tráfico",
    // Añade tus propias secciones aquí:
    "Commit Status",
    "Reports",
    "User-ID",
    "GlobalProtect",
    "Configuración Personalizada",
];

// Y en el match statement:
6 => format!("https://{}/#commit", self.config.firewall_ip),
7 => format!("https://{}/#monitor/reports", self.config.firewall_ip),
8 => format!("https://{}/#objects/userid", self.config.firewall_ip),
9 => format!("https://{}/#globalprotect/portal", self.config.firewall_ip),
```

Luego recompila:
```bash
cargo build --release
```

### 8. Temas y Apariencia

#### Firefox - Tema oscuro para trabajo nocturno

1. Abre Firefox con perfil palo-firefox
2. Instala "Dark Reader" o tema oscuro
3. Configura excepciones para el GUI de Palo Alto si es necesario

#### Brave - Optimización de rendimiento

```
brave://settings/
- Hardware acceleration: ON
- Memory saver: OFF (para mantener pestañas activas)
- Cookies: Allow for [tu-palo-alto-ip]
```

### 9. Automatización con Systemd

Inicia palo-launcher automáticamente al login:

```bash
# ~/.config/systemd/user/palo-launcher.service
[Unit]
Description=Palo Alto Launcher Auto-start
After=graphical-session.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/palo-launcher
RemainAfterExit=yes

[Install]
WantedBy=default.target
```

Habilitar:
```bash
systemctl --user enable palo-launcher.service
systemctl --user start palo-launcher.service
```

### 10. Integración con rofi/dmenu

Lanza palo-launcher desde tu launcher favorito:

```bash
# ~/.local/bin/palo-menu
#!/bin/bash

CHOICE=$(echo -e "Conectar\nConfigurar\nVer Config" | rofi -dmenu -p "Palo Alto")

case "$CHOICE" in
    "Conectar") palo-launcher connect ;;
    "Configurar") palo-launcher config ;;
    "Ver Config") palo-launcher show | rofi -dmenu ;;
esac
```

### 11. Logging y Debugging

Para troubleshooting, puedes habilitar logs:

```bash
# Ejecutar con verbose output
RUST_LOG=debug palo-launcher

# Guardar logs en archivo
palo-launcher 2>&1 | tee ~/palo-launcher.log
```

### 12. Variables de Entorno

Puedes usar variables de entorno para configuración rápida:

```bash
# ~/.bashrc o ~/.zshrc

export PALO_IP="192.168.1.1"
export PALO_BROWSER_1="firefox"
export PALO_BROWSER_2="brave"

# Función helper
palo() {
    case $1 in
        prod) PALO_IP="192.168.1.1" ;;
        dev) PALO_IP="192.168.2.1" ;;
    esac
    palo-launcher
}
```

## 🚀 Tips Finales

1. **Mantén tus perfiles limpios**: Limpia caché periódicamente
2. **Actualiza las extensiones**: Mantén las extensiones actualizadas
3. **Backup regular**: Respalda tus perfiles antes de actualizaciones importantes
4. **Usa atajos de teclado**: Aprende los shortcuts del GUI de Palo Alto
5. **Monitorea recursos**: Los múltiples navegadores consumen RAM, ciérralos cuando no uses

## 📚 Recursos Adicionales

- [Documentación oficial de Palo Alto](https://docs.paloaltonetworks.com/)
- [Best Practices de Palo Alto](https://live.paloaltonetworks.com/t5/community-blogs/best-practice-internet-gateway-design-guide/ba-p/453783)
- [Firefox Profile Manager](https://support.mozilla.org/en-US/kb/profile-manager-create-remove-switch-firefox-profiles)

---

**¿Tienes más ideas o configuraciones? ¡Compártelas!**
