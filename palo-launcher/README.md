# Palo Alto Launcher 🚀

Herramienta CLI en Rust para gestionar conexiones al Firewall Palo Alto con múltiples navegadores y sesiones persistentes.

## 🎯 Características

- ✅ Abre múltiples navegadores simultáneamente (Firefox, Brave, Chrome)
- ✅ Mantiene sesiones persistentes con perfiles separados
- ✅ Menú interactivo para acceder a diferentes secciones del firewall
- ✅ Configuración simple y reutilizable
- ✅ Evita el tedioso proceso de login constante

## 📦 Instalación

### Compilar desde el código fuente

```bash
# Clonar o copiar el proyecto
cd palo-launcher

# Compilar
cargo build --release

# Instalar (opcional)
sudo cp target/release/palo-launcher /usr/local/bin/
```

## 🔧 Configuración Inicial

### 1. Configurar la herramienta

```bash
palo-launcher config
```

Esto te preguntará:
- La IP de administración de tu Palo Alto
- Qué navegadores quieres usar por defecto

### 2. Crear perfiles de navegador

Los perfiles son CRUCIALES para mantener sesiones persistentes. Aquí están las instrucciones:

#### Firefox

```bash
# Crear perfil para Firefox
firefox -ProfileManager

# En el gestor de perfiles:
# 1. Click en "Crear perfil"
# 2. Nombre: palo-firefox
# 3. Si usas dos ventanas de Firefox, crea también: palo-firefox-2
# 4. NO marques "Usar el perfil seleccionado sin preguntar"
```

**Alternativa rápida (desde terminal):**
```bash
firefox -CreateProfile "palo-firefox"
firefox -CreateProfile "palo-firefox-2"
```

#### Brave

```bash
# Brave usará automáticamente un directorio de perfil en:
# ~/.config/BraveSoftware/Brave-Browser/palo-brave/
# Se creará automáticamente al primer uso
```

#### Chrome

```bash
# Chrome usará automáticamente un directorio de perfil en:
# ~/.config/google-chrome/palo-chrome/
# Se creará automáticamente al primer uso
```

### 3. Primera conexión y login

```bash
# Ejecutar la herramienta
palo-launcher

# O explícitamente:
palo-launcher connect
```

**IMPORTANTE:** La primera vez que uses cada navegador:
1. Se abrirán los navegadores con los perfiles nuevos
2. Loguéate manualmente en cada uno
3. **Marca la opción "Recordar credenciales"** o similar
4. Las siguientes veces ya estarás logueado automáticamente

## 🎮 Uso

### Modo interactivo (recomendado)

```bash
palo-launcher
```

Esto abrirá un menú donde puedes elegir:
- Dashboard Principal
- Monitor
- Policies
- Objects
- Network
- Logs de Tráfico
- Configuración Personalizada

### Comandos disponibles

```bash
# Conectar (modo interactivo)
palo-launcher connect

# Ver configuración actual
palo-launcher show

# Reconfigurar
palo-launcher config
```

## 💡 Ventajas de usar perfiles separados

1. **Sesiones independientes**: Cada navegador mantiene su propia sesión con Palo Alto
2. **Sin conflictos**: Puedes tener la misma página abierta en diferentes navegadores
3. **Persistencia**: Las sesiones se mantienen incluso después de cerrar los navegadores
4. **Organización**: Tus navegadores personales no se mezclan con tu trabajo de Palo Alto

## 🔍 Cómo funciona

La herramienta:

1. Abre Firefox con el perfil `palo-firefox` usando `firefox -P palo-firefox --new-instance`
2. Abre Brave/Chrome con un directorio de datos específico usando `--user-data-dir`
3. Cada perfil almacena sus propias cookies, sesiones y configuraciones
4. Al detectar inactividad en Palo Alto, solo necesitas refrescar la página (no reloguear)

## 🛠️ Solución de problemas

### "Firefox no se encuentra"
```bash
# Instalar Firefox
sudo apt install firefox  # Debian/Ubuntu
sudo dnf install firefox  # Fedora
```

### "Brave no se encuentra"
```bash
# Instalar Brave
sudo curl -fsSLo /usr/share/keyrings/brave-browser-archive-keyring.gpg \
  https://brave-browser-apt-release.s3.brave.com/brave-browser-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/brave-browser-archive-keyring.gpg] \
  https://brave-browser-apt-release.s3.brave.com/ stable main" | \
  sudo tee /etc/apt/sources.list.d/brave-browser-release.list

sudo apt update
sudo apt install brave-browser
```

### "Las sesiones no persisten"

Verifica que:
1. Los perfiles de Firefox se crearon correctamente: `ls ~/.mozilla/firefox/`
2. No estás usando modo incógnito/privado
3. Las cookies no están bloqueadas en el navegador
4. El certificado SSL de Palo Alto está aceptado en cada navegador

### "Quiero cambiar los navegadores por defecto"

```bash
palo-launcher config
# Selecciona una nueva combinación
```

## 📋 Ejemplos de uso típicos

### Caso 1: Revisar Policies y Monitor simultáneamente

```bash
palo-launcher
# Selecciona "Policies"
```

Luego en uno de los navegadores, navega manualmente a Monitor. Ambas sesiones se mantienen.

### Caso 2: Trabajo diario

```bash
# Crear un alias en tu ~/.bashrc o ~/.zshrc
alias palo='palo-launcher'

# Ahora solo ejecuta:
palo
```

### Caso 3: Logs mientras configuras políticas

1. Abre `palo-launcher`
2. Selecciona "Dashboard Principal"
3. En Firefox navega a Policies
4. En Brave navega a Monitor > Logs
5. Trabaja fluidamente entre ambos

## 🎨 Personalización

### Cambiar los nombres de perfiles

Edita el archivo de configuración:
```bash
# La configuración se guarda en:
# Linux: ~/.config/palo-launcher/config.json

nano ~/.config/palo-launcher/config.json
```

Ejemplo de configuración:
```json
{
  "firewall_ip": "192.168.1.1",
  "firefox_profile": "palo-firefox",
  "brave_profile": "palo-brave",
  "chrome_profile": "palo-chrome",
  "default_browsers": ["firefox", "brave"]
}
```

## 🚀 Consejos Pro

1. **Crea perfiles temáticos**: Usa extensiones diferentes en cada perfil
   - Firefox: Para Policies (con extensión de JSON viewer)
   - Brave: Para Monitor (con bloqueador de ads desactivado)

2. **Atajos de teclado**: Aprende los shortcuts de Palo Alto para navegar más rápido

3. **Múltiples firewalls**: Si administras varios Palo Alto, crea un alias por cada uno:
   ```bash
   alias palo-prod='palo-launcher'
   alias palo-dev='PALO_IP=192.168.2.1 palo-launcher'
   ```

4. **Bookmarks en perfiles**: Añade bookmarks en cada perfil de navegador para acceso aún más rápido

## 📝 Notas

- La herramienta NO almacena credenciales, solo gestiona perfiles de navegador
- Cada perfil es independiente y se comporta como un navegador completamente nuevo
- Los perfiles creados ocupan espacio mínimo (~50MB por perfil)
- Puedes usar los mismos perfiles para otros sistemas web si lo deseas

## 🤝 Contribuciones

¿Ideas para mejorar? ¡Las sugerencias son bienvenidas!

## 📄 Licencia

MIT License - Úsalo libremente para tu administración de Palo Alto

---

**Desarrollado con ❤️ para administradores de Palo Alto que valoran su tiempo**
