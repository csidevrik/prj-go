#!/bin/bash

# Script de instalación para Palo Alto Launcher

set -e

# Colores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "════════════════════════════════════════════════════════════"
echo "       Instalador de Palo Alto Launcher v0.1.0"
echo "════════════════════════════════════════════════════════════"
echo ""

# Verificar que estamos en el directorio correcto
if [ ! -f "Cargo.toml" ]; then
    echo -e "${RED}✗ Error: Debes ejecutar este script desde el directorio palo-launcher${NC}"
    exit 1
fi

# Verificar si Rust está instalado
if ! command -v cargo &> /dev/null; then
    echo -e "${RED}✗ Rust/Cargo no está instalado${NC}"
    echo ""
    echo "Por favor instala Rust primero:"
    echo "  curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh"
    exit 1
fi

echo -e "${GREEN}✓${NC} Rust/Cargo encontrado"
echo ""

# Compilar el proyecto
echo -e "${BLUE}Compilando palo-launcher...${NC}"
cargo build --release

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓${NC} Compilación exitosa"
else
    echo -e "${RED}✗ Error en la compilación${NC}"
    exit 1
fi

echo ""
echo "════════════════════════════════════════════════════════════"
echo "           Opciones de instalación"
echo "════════════════════════════════════════════════════════════"
echo ""
echo "1. Instalar en /usr/local/bin (requiere sudo, disponible globalmente)"
echo "2. Instalar en ~/.local/bin (solo tu usuario, no requiere sudo)"
echo "3. Solo compilar (no instalar, ejecutar desde target/release/)"
echo ""

read -p "Selecciona una opción [1-3]: " option

case $option in
    1)
        echo ""
        echo -e "${BLUE}Instalando en /usr/local/bin...${NC}"
        sudo cp target/release/palo-launcher /usr/local/bin/
        sudo chmod +x /usr/local/bin/palo-launcher
        echo -e "${GREEN}✓${NC} Instalado en /usr/local/bin/palo-launcher"
        INSTALLED_PATH="/usr/local/bin/palo-launcher"
        ;;
    2)
        echo ""
        echo -e "${BLUE}Instalando en ~/.local/bin...${NC}"
        mkdir -p ~/.local/bin
        cp target/release/palo-launcher ~/.local/bin/
        chmod +x ~/.local/bin/palo-launcher
        echo -e "${GREEN}✓${NC} Instalado en ~/.local/bin/palo-launcher"
        
        # Verificar si ~/.local/bin está en el PATH
        if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
            echo ""
            echo -e "${YELLOW}⚠ Advertencia: ~/.local/bin no está en tu PATH${NC}"
            echo ""
            echo "Añade esta línea a tu ~/.bashrc o ~/.zshrc:"
            echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
            echo ""
            echo "Luego ejecuta: source ~/.bashrc  (o source ~/.zshrc)"
        fi
        INSTALLED_PATH="~/.local/bin/palo-launcher"
        ;;
    3)
        echo ""
        echo -e "${GREEN}✓${NC} Compilación completada"
        echo "El binario está en: $(pwd)/target/release/palo-launcher"
        echo ""
        echo "Para ejecutarlo:"
        echo "  ./target/release/palo-launcher"
        INSTALLED_PATH="./target/release/palo-launcher"
        ;;
    *)
        echo -e "${RED}Opción inválida${NC}"
        exit 1
        ;;
esac

echo ""
echo "════════════════════════════════════════════════════════════"
echo -e "${GREEN}✓ Instalación completada${NC}"
echo "════════════════════════════════════════════════════════════"
echo ""

# Preguntar si desea configurar los perfiles de Firefox
read -p "¿Deseas configurar los perfiles de Firefox ahora? (s/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[SsYy]$ ]]; then
    echo ""
    ./setup-firefox-profiles.sh
fi

echo ""
echo "════════════════════════════════════════════════════════════"
echo "              Próximos pasos"
echo "════════════════════════════════════════════════════════════"
echo ""
echo "1. Configura la IP de tu Palo Alto:"
echo "   ${YELLOW}palo-launcher config${NC}"
echo ""
echo "2. Conecta al firewall:"
echo "   ${YELLOW}palo-launcher${NC}"
echo ""
echo "3. La primera vez, loguéate en cada navegador y marca"
echo "   'Recordar credenciales' para mantener la sesión"
echo ""
echo "4. (Opcional) Crea un alias para acceso rápido:"
echo "   ${YELLOW}echo 'alias palo=\"palo-launcher\"' >> ~/.bashrc${NC}"
echo ""
echo -e "${GREEN}¡Disfruta de una administración más eficiente de Palo Alto!${NC}"
echo ""
