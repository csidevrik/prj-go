package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/csidevrik/bingofin-cli/pkg/models"
)

// ExportarResumenCSV exporta el resumen total a un archivo CSV
func ExportarResumenCSV(resumen models.ResumenTotal, nombreArchivo string) error {
	file, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir encabezado del resumen
	writer.Write([]string{"RESUMEN TOTAL"})
	writer.Write([]string{"Inversión Inicial", fmt.Sprintf("$%.2f", resumen.InversionInicial)})
	writer.Write([]string{"Interés Total", fmt.Sprintf("$%.2f", resumen.InteresTotal)})
	writer.Write([]string{"Capital Final", fmt.Sprintf("$%.2f", resumen.CapitalFinal)})
	writer.Write([]string{"Rendimiento", fmt.Sprintf("%.2f%%", resumen.Rendimiento)})
	writer.Write([]string{}) // Línea vacía

	// Escribir detalle de cada plazo
	for _, resultado := range resumen.Plazos {
		writer.Write([]string{fmt.Sprintf("PLAZO: %s", resultado.Configuracion.Nombre)})
		writer.Write([]string{
			"Período",
			"Fecha Inicio",
			"Fecha Vence",
			"Capital Inicio",
			"Interés",
			"Capital Final",
			"Aporte",
			"Nuevo Capital",
		})

		for _, periodo := range resultado.Periodos {
			writer.Write([]string{
				fmt.Sprintf("%d", periodo.Numero),
				periodo.FechaInicio.Format("2006-01-02"),
				periodo.FechaVence.Format("2006-01-02"),
				fmt.Sprintf("%.2f", periodo.CapitalInicio),
				fmt.Sprintf("%.2f", periodo.Interes),
				fmt.Sprintf("%.2f", periodo.CapitalFinal),
				fmt.Sprintf("%.2f", periodo.Aporte),
				fmt.Sprintf("%.2f", periodo.NuevoCapital),
			})
		}

		// Totales del plazo
		var totalInteres float64
		for _, p := range resultado.Periodos {
			totalInteres += p.Interes
		}
		ultimoPeriodo := resultado.Periodos[len(resultado.Periodos)-1]

		writer.Write([]string{
			"TOTAL",
			"",
			"",
			fmt.Sprintf("%.2f", resultado.Configuracion.Capital),
			fmt.Sprintf("%.2f", totalInteres),
			"",
			"",
			fmt.Sprintf("%.2f", ultimoPeriodo.NuevoCapital),
		})
		writer.Write([]string{}) // Línea vacía entre plazos
	}

	return nil
}

// ExportarPlazoCSV exporta un solo plazo a CSV con formato detallado
func ExportarPlazoCSV(resultado models.ResultadoPlazo, nombreArchivo string) error {
	file, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Información del plazo
	writer.Write([]string{"Plazo Fijo", resultado.Configuracion.Nombre})
	writer.Write([]string{"Días", fmt.Sprintf("%d", resultado.Configuracion.Dias)})
	writer.Write([]string{"Tasa", fmt.Sprintf("%.2f%%", resultado.Configuracion.Tasa)})
	writer.Write([]string{"Capital Inicial", fmt.Sprintf("$%.2f", resultado.Configuracion.Capital)})
	writer.Write([]string{"Renovaciones", fmt.Sprintf("%d", resultado.Configuracion.Renovaciones)})
	writer.Write([]string{})

	// Encabezado de períodos
	writer.Write([]string{
		"Período",
		"Fecha Inicio",
		"Fecha Vence",
		"Días",
		"Capital Inicio",
		"Tasa",
		"Interés",
		"Capital Final",
		"Aporte",
		"Nuevo Capital",
	})

	// Datos de períodos
	for _, periodo := range resultado.Periodos {
		writer.Write([]string{
			fmt.Sprintf("%d", periodo.Numero),
			periodo.FechaInicio.Format("2006-01-02"),
			periodo.FechaVence.Format("2006-01-02"),
			fmt.Sprintf("%d", periodo.Dias),
			fmt.Sprintf("%.2f", periodo.CapitalInicio),
			fmt.Sprintf("%.2f%%", periodo.Tasa),
			fmt.Sprintf("%.2f", periodo.Interes),
			fmt.Sprintf("%.2f", periodo.CapitalFinal),
			fmt.Sprintf("%.2f", periodo.Aporte),
			fmt.Sprintf("%.2f", periodo.NuevoCapital),
		})
	}

	return nil
}

// GenerarNombreArchivo genera un nombre de archivo con timestamp
func GenerarNombreArchivo(prefijo string) string {
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("%s-%s.csv", prefijo, timestamp)
}
