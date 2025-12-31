package models

import "time"

// PlazoFijo representa la configuración de un plazo fijo
type PlazoFijo struct {
	Nombre       string  // "30 días", "60 días", etc.
	Dias         int     // Duración del plazo en días
	Tasa         float64 // Tasa de interés anual (porcentaje)
	Capital      float64 // Capital inicial
	Renovaciones int     // Número de veces que se renueva
	Color        string  // Color para visualización (opcional)
}

// Periodo representa un período individual de inversión
type Periodo struct {
	Numero        int       // Número del período (1, 2, 3, ...)
	FechaInicio   time.Time // Fecha de inicio del período
	FechaVence    time.Time // Fecha de vencimiento
	CapitalInicio float64   // Capital al inicio del período
	Interes       float64   // Interés generado en este período
	CapitalFinal  float64   // Capital + Interés
	Aporte        float64   // Aporte adicional al final del período
	NuevoCapital  float64   // Capital Final + Aporte (capital para próximo período)
	Dias          int       // Días del plazo
	Tasa          float64   // Tasa aplicada
}

// ResultadoPlazo contiene todos los períodos de un plazo fijo
type ResultadoPlazo struct {
	Configuracion PlazoFijo // Configuración original
	Periodos      []Periodo // Todos los períodos calculados
}

// ResumenTotal contiene los totales consolidados
type ResumenTotal struct {
	InversionInicial float64          // Capital inicial total + aportes disponibles
	InteresTotal     float64          // Suma de todos los intereses
	CapitalFinal     float64          // Capital final total
	Rendimiento      float64          // Porcentaje de rendimiento
	Plazos           []ResultadoPlazo // Resultados de cada plazo
}

// ConfiguracionSimulacion contiene los parámetros de la simulación
type ConfiguracionSimulacion struct {
	FechaInicio        time.Time   // Fecha de inicio de la simulación
	Plazos             []PlazoFijo // Configuración de todos los plazos
	AportesDisponibles float64     // Dinero adicional disponible para aportes
	AportePorPeriodo   float64     // Cuánto aportar en cada período elegible
	PlazoConAportes    string      // Nombre del plazo que recibe aportes ("30 días")
	NumAportes         int         // Número de aportes a realizar
}

// NewPlazoFijo crea un nuevo plazo fijo con valores por defecto
func NewPlazoFijo(nombre string, dias int, tasa float64, capital float64, renovaciones int) PlazoFijo {
	return PlazoFijo{
		Nombre:       nombre,
		Dias:         dias,
		Tasa:         tasa,
		Capital:      capital,
		Renovaciones: renovaciones,
		Color:        "", // Se puede asignar después
	}
}

// NewConfiguracionSimulacion crea una configuración por defecto
func NewConfiguracionSimulacion(fechaInicio time.Time) *ConfiguracionSimulacion {
	return &ConfiguracionSimulacion{
		FechaInicio:        fechaInicio,
		Plazos:             []PlazoFijo{},
		AportesDisponibles: 0,
		AportePorPeriodo:   0,
		PlazoConAportes:    "",
		NumAportes:         0,
	}
}
