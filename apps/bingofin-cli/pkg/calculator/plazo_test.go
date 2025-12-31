package calculator

import (
	"math"
	"testing"
	"time"

	"github.com/csidevrik/bingofin-cli/pkg/models"
)

// floatEquals compara dos floats con tolerancia de error
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestCalcularInteres(t *testing.T) {
	tests := []struct {
		name     string
		capital  float64
		tasa     float64
		dias     int
		expected float64
	}{
		{
			name:     "Plazo 30 días - $500 @ 4.6%",
			capital:  500.00,
			tasa:     4.6,
			dias:     30,
			expected: 1.92, // 500 * 0.046 * (30/360) = 1.916...
		},
		{
			name:     "Plazo 360 días - $4390.69 @ 6.5%",
			capital:  4390.69,
			tasa:     6.5,
			dias:     360,
			expected: 285.39, // 4390.69 * 0.065 * (360/360) = 285.3948...
		},
		{
			name:     "Plazo 60 días - $500 @ 4.9%",
			capital:  500.00,
			tasa:     4.9,
			dias:     60,
			expected: 4.08, // 500 * 0.049 * (60/360) = 4.083...
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalcularInteres(tt.capital, tt.tasa, tt.dias)

			// Comparar con tolerancia de $0.01
			if !floatEquals(result, tt.expected, 0.01) {
				t.Errorf("CalcularInteres() = %.2f, esperado %.2f", result, tt.expected)
			}
		})
	}
}

func TestGenerarPeriodos(t *testing.T) {
	fechaInicio := time.Date(2025, 12, 22, 0, 0, 0, 0, time.UTC)

	// Configurar plazo de 30 días con 2 renovaciones (3 períodos totales)
	plazo := models.NewPlazoFijo("30 días", 30, 4.6, 500.00, 2)

	// Configuración de aportes
	config := models.ConfiguracionSimulacion{
		FechaInicio:        fechaInicio,
		AportesDisponibles: 500,
		AportePorPeriodo:   100,
		PlazoConAportes:    "30 días",
		NumAportes:         2, // Solo primeros 2 períodos reciben aportes
	}

	periodos := GenerarPeriodos(plazo, fechaInicio, 0, config)

	// Verificar que se generaron 3 períodos
	if len(periodos) != 3 {
		t.Errorf("Esperaban 3 períodos, obtuvieron %d", len(periodos))
	}

	// Verificar primer período
	if periodos[0].Numero != 1 {
		t.Errorf("Período 1: número incorrecto")
	}

	if !floatEquals(periodos[0].CapitalInicio, 500.00, 0.01) {
		t.Errorf("Período 1: capital inicio = %.2f, esperado 500.00", periodos[0].CapitalInicio)
	}

	if !floatEquals(periodos[0].Interes, 1.92, 0.01) {
		t.Errorf("Período 1: interés = %.2f, esperado ~1.92", periodos[0].Interes)
	}

	if !floatEquals(periodos[0].Aporte, 100.00, 0.01) {
		t.Errorf("Período 1: aporte = %.2f, esperado 100.00", periodos[0].Aporte)
	}

	// El nuevo capital debería ser: 500 + 1.92 + 100 = 601.92
	expectedNuevoCapital := 500.00 + periodos[0].Interes + 100.00
	if !floatEquals(periodos[0].NuevoCapital, expectedNuevoCapital, 0.01) {
		t.Errorf("Período 1: nuevo capital = %.2f, esperado %.2f",
			periodos[0].NuevoCapital, expectedNuevoCapital)
	}

	// Verificar que el segundo período inicia con el capital acumulado
	if !floatEquals(periodos[1].CapitalInicio, periodos[0].NuevoCapital, 0.01) {
		t.Errorf("Período 2: debería iniciar con capital del período 1")
	}

	// Verificar que el tercer período NO recibe aporte (NumAportes = 2)
	if periodos[2].Aporte != 0.0 {
		t.Errorf("Período 3: no debería recibir aporte")
	}
}

func TestSimularTodosLosPlazo(t *testing.T) {
	fechaInicio := time.Date(2025, 12, 22, 0, 0, 0, 0, time.UTC)

	// Crear configuración simple: un plazo de 360 días
	config := models.ConfiguracionSimulacion{
		FechaInicio:        fechaInicio,
		AportesDisponibles: 0,
		AportePorPeriodo:   0,
		NumAportes:         0,
		Plazos: []models.PlazoFijo{
			models.NewPlazoFijo("360 días", 360, 6.5, 4390.69, 0),
		},
	}

	resultado := SimularTodosLosPlazo(config)

	// Verificar inversión inicial
	if !floatEquals(resultado.InversionInicial, 4390.69, 0.01) {
		t.Errorf("Inversión inicial = %.2f, esperado 4390.69", resultado.InversionInicial)
	}

	// Verificar interés total
	if !floatEquals(resultado.InteresTotal, 285.39, 0.5) {
		t.Errorf("Interés total = %.2f, esperado ~285.39", resultado.InteresTotal)
	}

	// Verificar capital final
	expectedCapitalFinal := 4390.69 + resultado.InteresTotal
	if !floatEquals(resultado.CapitalFinal, expectedCapitalFinal, 0.01) {
		t.Errorf("Capital final = %.2f, esperado %.2f",
			resultado.CapitalFinal, expectedCapitalFinal)
	}

	// Verificar que se calculó rendimiento
	if resultado.Rendimiento <= 0 {
		t.Errorf("Rendimiento debería ser positivo")
	}
}

func TestCalcularTotales(t *testing.T) {
	// Crear resultado simulado con 2 plazos
	resultados := []models.ResultadoPlazo{
		{
			Configuracion: models.NewPlazoFijo("Plazo 1", 360, 6.5, 1000.00, 0),
			Periodos: []models.Periodo{
				{
					Numero:        1,
					CapitalInicio: 1000.00,
					Interes:       65.00,
					CapitalFinal:  1065.00,
					Aporte:        0,
					NuevoCapital:  1065.00,
				},
			},
		},
		{
			Configuracion: models.NewPlazoFijo("Plazo 2", 360, 5.0, 500.00, 0),
			Periodos: []models.Periodo{
				{
					Numero:        1,
					CapitalInicio: 500.00,
					Interes:       25.00,
					CapitalFinal:  525.00,
					Aporte:        0,
					NuevoCapital:  525.00,
				},
			},
		},
	}

	aportesDisponibles := 100.00
	totales := CalcularTotales(resultados, aportesDisponibles)

	// Inversión inicial debería ser: 1000 + 500 + 100 = 1600
	if !floatEquals(totales.InversionInicial, 1600.00, 0.01) {
		t.Errorf("Inversión inicial = %.2f, esperado 1600.00", totales.InversionInicial)
	}

	// Interés total: 65 + 25 = 90
	if !floatEquals(totales.InteresTotal, 90.00, 0.01) {
		t.Errorf("Interés total = %.2f, esperado 90.00", totales.InteresTotal)
	}

	// Capital final: 1065 + 525 = 1590
	if !floatEquals(totales.CapitalFinal, 1590.00, 0.01) {
		t.Errorf("Capital final = %.2f, esperado 1590.00", totales.CapitalFinal)
	}

	// Rendimiento: (90 / 1600) * 100 = 5.625%
	expectedRendimiento := 5.625
	if !floatEquals(totales.Rendimiento, expectedRendimiento, 0.01) {
		t.Errorf("Rendimiento = %.2f%%, esperado %.2f%%",
			totales.Rendimiento, expectedRendimiento)
	}
}
