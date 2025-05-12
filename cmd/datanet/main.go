package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Datanet - Drag & Drop CSV")
	w.Resize(fyne.NewSize(600, 400))

	infoLabel := widget.NewLabel("Arrastra un archivo .csv sobre esta ventana.")
	content := container.NewVBox(infoLabel)
	w.SetContent(content)

	// Configurar el manejador de archivos arrastrados
	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			path := uri.Path()
			if !isCSV(path) {
				infoLabel.SetText("Archivo no válido. Por favor, arrastra un archivo .csv.")
				continue
			}

			records, err := parseCSV(path)
			if err != nil {
				infoLabel.SetText(fmt.Sprintf("Error al leer el archivo: %v", err))
				continue
			}

			infoLabel.SetText(fmt.Sprintf("Archivo cargado: %s\nRegistros: %d", path, len(records)))
			// Aquí puedes procesar los registros según tus necesidades
		}
	})

	w.ShowAndRun()
}

// Verifica si el archivo tiene extensión .csv
func isCSV(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".csv"
}

// Lee y parsea el archivo CSV
func parseCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
