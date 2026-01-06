#!/bin/bash

# Script para configurar automáticamente los perfiles de Firefox para Palo Alto Launcher
# Este script facilita la creación de perfiles sin usar la GUI

set -e

echo "════════════════════════════════════════════════════════════"
echo "  Configurador de Perfiles Firefox - Palo Alto Launcher"
echo "════════════════════════════════════════════════════════════"
echo ""

# Colores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Verificar si Firefox está instalado
if ! command -v firefox &> /dev/null && ! command -v firefox-esr &> /dev/null; then
    echo -e "${RED}✗ Firefox no está instalado${NC}"
    echo ""
    echo "Por favor instala Firefox primero:"
    echo "  Ubuntu/Debian: sudo apt install firefox"
    echo "  Fedora: sudo dnf install firefox"
    echo "  Arch: sudo pacman -S firefox"
    exit 1
fi

FIREFOX_CMD="firefox"
if command -v firefox-esr &> /dev/null; then
    FIREFOX_CMD="firefox-esr"
fi

echo -e "${GREEN}✓${NC} Firefox encontrado: $FIREFOX_CMD"
echo ""

# Crear primer perfil
echo "Creando perfil: palo-firefox..."
$FIREFOX_CMD -CreateProfile "palo-firefox" 2>/dev/null || true
echo -e "${GREEN}✓${NC} Perfil 'palo-firefox' creado"

# Crear segundo perfil (para usar dos ventanas de Firefox)
echo "Creando perfil: palo-firefox-2..."
$FIREFOX_CMD -CreateProfile "palo-firefox-2" 2>/dev/null || true
echo -e "${GREEN}✓${NC} Perfil 'palo-firefox-2' creado"

echo ""
echo "════════════════════════════════════════════════════════════"
echo -e "${GREEN}✓ Perfiles creados exitosamente${NC}"
echo "════════════════════════════════════════════════════════════"
echo ""
echo "Perfiles disponibles:"
echo "  • palo-firefox"
echo "  • palo-firefox-2"
echo ""
echo -e "${YELLOW}SIGUIENTE PASO:${NC}"
echo "1. Ejecuta: palo-launcher config"
echo "2. Configura la IP de tu Palo Alto"
echo "3. Ejecuta: palo-launcher"
echo "4. La primera vez, loguéate y marca 'Recordar credenciales'"
echo ""
echo "Los perfiles se encuentran en: ~/.mozilla/firefox/"
echo ""

# Opcional: abrir Firefox con el primer perfil para configuración inicial
read -p "¿Deseas abrir Firefox con el perfil 'palo-firefox' ahora para configurarlo? (s/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[SsYy]$ ]]; then
    echo "Abriendo Firefox con perfil palo-firefox..."
    echo "Configura lo que necesites y cierra Firefox cuando termines."
    $FIREFOX_CMD -P palo-firefox &
fi

echo ""
echo -e "${GREEN}¡Listo! Ya puedes usar palo-launcher${NC}"
