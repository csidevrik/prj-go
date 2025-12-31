package main

import (
	"fmt"
	"os"
	"time"

	"github.com/csidevrik/bingofin-cli/pkg/calculator"
	"github.com/csidevrik/bingofin-cli/pkg/exporter"
	"github.com/csidevrik/bingofin-cli/pkg/models"
	ui "github.com/csidevrik/bingofin-cli/pkg/texts"
)

func main() {
	u := ui.New(os.Stdout).WithStyle(ui.Style{
		UnderlineChar: '=',
		PadY:          1,
	})
	u.Title("🎯 BingoFin CLI - Calculadora de Plazos Fijos y mas")

	// Configurar fecha de inicio
	fechaInicio := time.Date(2025, 12, 22, 0, 0, 0, 0, time.UTC)

	// Configurar todos los plazos fijos
	plazos := []models.PlazoFijo{
		models.NewPlazoFijo("360 días", 360, 6.5, 4390.69, 0),
		models.NewPlazoFijo("270 días", 270, 5.6, 500.00, 0),
		models.NewPlazoFijo("180 días", 180, 5.4, 500.00, 1),
		models.NewPlazoFijo("90 días", 90, 5.1, 500.00, 3),
		models.NewPlazoFijo("60 días", 60, 4.9, 500.00, 5),
		models.NewPlazoFijo("30 días", 30, 4.6, 500.00, 11),
	}

	// Configurar simulación
	config := models.ConfiguracionSimulacion{
		FechaInicio:        fechaInicio,
		Plazos:             plazos,
		AportesDisponibles: 500.00,    // $500 adicionales
		AportePorPeriodo:   100.00,    // $100 por período
		PlazoConAportes:    "60 días", // Solo el plazo de 60 días recibe aportes
		NumAportes:         5,         // Primeros 5 períodos
	}

	// Ejecutar simulación
	u.Line("📊 Ejecutando simulación...")
	resultado := calculator.SimularTodosLosPlazo(config)

	// Mostrar resultados

	u.Section("💰 RESULTADOS DE LA SIMULACIÓN")

	fmt.Printf("Inversión Total:    $%12.3f\n", resultado.InversionInicial)
	fmt.Printf("Interés Generado:   $%12.2f\n", resultado.InteresTotal)
	fmt.Printf("Capital Final:      $%12.2f\n", resultado.CapitalFinal)
	fmt.Printf("Rendimiento:        %12.2f%%\n", resultado.Rendimiento)
	u.Line("*-*-**-*-**-*-**-*-**-*-*")

	// Mostrar detalles por plazo
	u.Section("📋 DETALLE POR PLAZO")

	for _, plazo := range resultado.Plazos {
		// Calcular totales del plazo
		var interesTotal float64
		for _, p := range plazo.Periodos {
			interesTotal += p.Interes
		}
		ultimoPeriodo := plazo.Periodos[len(plazo.Periodos)-1]

		fmt.Printf("\n%s:\n", plazo.Configuracion.Nombre)
		fmt.Printf("  Capital Inicial:  $%8.2f\n", plazo.Configuracion.Capital)
		fmt.Printf("  Tasa:             %8.2f%%\n", plazo.Configuracion.Tasa)
		fmt.Printf("  Períodos:         %8d\n", len(plazo.Periodos))
		fmt.Printf("  Interés Total:    $%8.2f\n", interesTotal)
		fmt.Printf("  Capital Final:    $%8.2f\n", ultimoPeriodo.NuevoCapital)
	}

	// Exportar a CSV
	fmt.Println()
	u.Section("📁 Exportando resultados...")
	nombreArchivo := exporter.GenerarNombreArchivo("bingofin-simulacion")
	err := exporter.ExportarResumenCSV(resultado, nombreArchivo)
	if err != nil {
		fmt.Printf("❌ Error exportando: %v\n", err)
		return
	}
	fmt.Printf("✅ Exportado exitosamente a: %s\n", nombreArchivo)

	// Mostrar detalle del plazo más interesante (30 días con aportes)

	u.Section("🔍 DETALLE DEL PLAZO DE 60 DÍAS (con aportes)")

	plazo30 := resultado.Plazos[4] // Último en el array

	fmt.Printf("%-8s %-12s %-12s %-12s %-10s %-12s\n",
		"Período", "Capital", "Interés", "Cap.Final", "Aporte", "Nuevo Cap.")
	fmt.Println("-------------------------------------------------------------------")

	for _, p := range plazo30.Periodos {
		fmt.Printf("%-8d $%10.2f $%10.2f $%10.2f $%8.2f $%10.2f\n",
			p.Numero,
			p.CapitalInicio,
			p.Interes,
			p.CapitalFinal,
			p.Aporte,
			p.NuevoCapital,
		)
	}

	fmt.Println()
	fmt.Println("✨ Simulación completada exitosamente")
}
