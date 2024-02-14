// lib/sayFormat.go

package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Formato representa la estructura de datos para cada formato en el archivo JSON
type Formato struct {
	NAME      string `json:"NAME"`
	MACFORMAT string `json:"MACFORMAT"`
}

// ObtenerFormatos lee el contenido del archivo formats.json y devuelve la lista de formatos
func ObtenerFormatos() ([]Formato, error) {
	// Obtener la ruta completa del archivo formats.json
	rutaArchivo := filepath.Join(".", "/conf/formats.json")

	// Leer el contenido del archivo
	contenido, err := os.ReadFile(rutaArchivo)
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %v", err)
	}

	// Deserializar el contenido JSON
	var formatos []Formato
	if err := json.Unmarshal(contenido, &formatos); err != nil {
		return nil, fmt.Errorf("error al deserializar JSON: %v", err)
	}

	return formatos, nil
}
