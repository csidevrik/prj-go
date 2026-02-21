package project

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"documentyzer/internal/meta"
)

// InitProject crea la estructura inicial para un proyecto/script
func InitProject(folderPath string) error {
	// Verificar que la carpeta existe
	info, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("la carpeta no existe: %s", folderPath)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("la ruta no es una carpeta: %s", folderPath)
	}

	// Extraer el nombre de la carpeta
	folderName := filepath.Base(folderPath)

	// Crear meta.json
	metaData := &meta.Meta{
		Name:        folderName,
		Description: fmt.Sprintf("Descripción de %s", folderName),
		Distros: []meta.Distro{
			{
				Name:    "Ubuntu",
				Version: "22.04+",
				Status:  "untested",
			},
		},
		Version:     "1.0.0",
		LastUpdated: time.Now().Format("2006-01-02"),
		Tags:        []string{"script", "automation"},
	}

	if err := meta.Save(folderPath, metaData); err != nil {
		return fmt.Errorf("error creando meta.json: %w", err)
	}
	fmt.Printf("✅ Creado: meta.json\n")

	// Crear README.md desde template
	readmeContent := fmt.Sprintf(`# %s

## Descripción

Descripción de %s

## Información

- **Versión**: 1.0.0
- **Última actualización**: %s

## Distribuciones Soportadas

- Ubuntu (22.04+) - untested

## Instalación

Ejecutar el script de instalación:

    chmod +x install.sh
    ./install.sh

## Uso

[Describe cómo usar este script/proyecto]

## Notas

[Información adicional importante]
`, folderName, folderName, time.Now().Format("2006-01-02"))

	readmePath := filepath.Join(folderPath, "README.md")
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("error creando README.md: %w", err)
	}
	fmt.Printf("✅ Creado: README.md\n")

	// Crear install.sh
	installContent := `#!/bin/bash

# Script de instalación para ` + folderName + `
# Actualiza este script según tus necesidades

set -e

echo "Instalando ` + folderName + `..."

# TODO: Agregar lógica de instalación aquí

echo "✅ Instalación completada"
`

	installPath := filepath.Join(folderPath, "install.sh")
	if err := os.WriteFile(installPath, []byte(installContent), 0755); err != nil {
		return fmt.Errorf("error creando install.sh: %w", err)
	}
	fmt.Printf("✅ Creado: install.sh (ejecutable)\n")

	return nil
}

// InitProjectWithTemplate crea la estructura usando templates (para uso futuro)
func InitProjectWithTemplate(folderPath string, templatePath string) error {
	// Esta función podría usar los templates del directorio templates/
	// Por ahora, la mantenemos para futura expansión
	return InitProject(folderPath)
}
