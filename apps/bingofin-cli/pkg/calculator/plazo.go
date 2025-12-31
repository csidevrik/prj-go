package calculator

import (
	"time"

	"github.com/csidevrik/bingofin-cli/pkg/models"
)

// CalcularInteres calcula el interés simple para un plazo fijo
// Fórmula: Interés = Capital × (Tasa/100) × (Días/360)
func CalcularInteres(capital float64, tasa float64, dias int) float64 {
	return capital * (tasa / 100.0) * (float64(dias) / 360.0)
}

// GenerarPeriodos genera todos los períodos de un plazo fijo incluyendo renovaciones
func GenerarPeriodos(config models.PlazoFijo, fechaInicio time.Time, offsetDias int, aportesConfig models.ConfiguracionSimulacion) []models.Periodo {
	periodos := make([]models.Periodo, 0)

	// Ajustar fecha de inicio con offset (para plazos escalonados)
	fechaInicioActual := fechaInicio.AddDate(0, 0, offsetDias)
	capitalActual := config.Capital

	// Calcular número total de períodos (renovaciones + período inicial)
	numPeriodos := config.Renovaciones + 1

	for i := 0; i < numPeriodos; i++ {
		// Calcular fecha de vencimiento
		fechaVence := fechaInicioActual.AddDate(0, 0, config.Dias)

		// Calcular interés de este período
		interes := CalcularInteres(capitalActual, config.Tasa, config.Dias)
		capitalFinal := capitalActual + interes

		// Calcular aporte (si corresponde)
		aporte := 0.0
		if config.Nombre == aportesConfig.PlazoConAportes && i < aportesConfig.NumAportes {
			aporte = aportesConfig.AportePorPeriodo
		}

		// Nuevo capital para el próximo período
		nuevoCapital := capitalFinal + aporte

		// Crear período
		periodo := models.Periodo{
			Numero:        i + 1,
			FechaInicio:   fechaInicioActual,
			FechaVence:    fechaVence,
			CapitalInicio: capitalActual,
			Interes:       interes,
			CapitalFinal:  capitalFinal,
			Aporte:        aporte,
			NuevoCapital:  nuevoCapital,
			Dias:          config.Dias,
			Tasa:          config.Tasa,
		}

		periodos = append(periodos, periodo)

		// Preparar próximo período
		capitalActual = nuevoCapital
		fechaInicioActual = fechaVence.AddDate(0, 0, 1) // Día siguiente al vencimiento
	}

	return periodos
}

// SimularTodosLosPlazo simula todos los plazos fijos según la configuración
func SimularTodosLosPlazo(config models.ConfiguracionSimulacion) models.ResumenTotal {
	resultados := make([]models.ResultadoPlazo, 0)

	// Generar períodos para cada plazo con offset escalonado
	for idx, plazoConfig := range config.Plazos {
		periodos := GenerarPeriodos(plazoConfig, config.FechaInicio, idx, config)

		resultado := models.ResultadoPlazo{
			Configuracion: plazoConfig,
			Periodos:      periodos,
		}

		resultados = append(resultados, resultado)
	}

	// Calcular totales
	return CalcularTotales(resultados, config.AportesDisponibles)
}

// CalcularTotales calcula los totales consolidados de todos los plazos
func CalcularTotales(resultados []models.ResultadoPlazo, aportesDisponibles float64) models.ResumenTotal {
	var inversionInicial float64
	var interesTotal float64
	var capitalFinal float64

	// Inversión inicial = suma de capitales iniciales + aportes disponibles
	for _, resultado := range resultados {
		inversionInicial += resultado.Configuracion.Capital
	}
	inversionTotal := inversionInicial + aportesDisponibles

	// Sumar todos los intereses de todos los períodos
	for _, resultado := range resultados {
		for _, periodo := range resultado.Periodos {
			interesTotal += periodo.Interes
		}
	}

	// Capital final = suma de los capitales finales del último período de cada plazo
	for _, resultado := range resultados {
		if len(resultado.Periodos) > 0 {
			ultimoPeriodo := resultado.Periodos[len(resultado.Periodos)-1]
			capitalFinal += ultimoPeriodo.NuevoCapital
		}
	}

	// Calcular rendimiento porcentual
	rendimiento := (interesTotal / inversionTotal) * 100.0

	return models.ResumenTotal{
		InversionInicial: inversionTotal,
		InteresTotal:     interesTotal,
		CapitalFinal:     capitalFinal,
		Rendimiento:      rendimiento,
		Plazos:           resultados,
	}
}

// ObtenerPeriodoPorFecha busca qué período está activo en una fecha específica
func ObtenerPeriodoPorFecha(resultado models.ResultadoPlazo, fecha time.Time) *models.Periodo {
	for i := range resultado.Periodos {
		periodo := &resultado.Periodos[i]
		if (fecha.Equal(periodo.FechaInicio) || fecha.After(periodo.FechaInicio)) &&
			(fecha.Equal(periodo.FechaVence) || fecha.Before(periodo.FechaVence)) {
			return periodo
		}
	}
	return nil
}
