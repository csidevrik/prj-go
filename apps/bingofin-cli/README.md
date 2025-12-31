# 🎯 BingoFin CLI

Calculadora de plazos fijos escrita en Go. Simula estrategias de inversión con múltiples plazos fijos, renovaciones automáticas y aportes periódicos.

## 🚀 Características

* ✅ Cálculo preciso de intereses simples
* ✅ Simulación de renovaciones automáticas
* ✅ Estrategia de aportes periódicos
* ✅ Múltiples plazos simultáneos (ladder strategy)
* ✅ Exportación a CSV
* ✅ Tests unitarios completos

## 📦 Instalación

```bash
# Clonar el proyecto
git clone https://github.com/csidevrik/bingofin-cli.git
cd bingofin-cli

# Descargar dependencias
go mod tidy

# Ejecutar
go run main.go
```

## 🧪 Testing

```bash
# Ejecutar todos los tests
go test ./...

# Ejecutar tests con cobertura
go test -cover ./...

# Ejecutar tests con detalle
go test -v ./pkg/calculator
```

## 🏗️ Estructura del Proyecto

```
bingofin-cli/
├── cmd/                    # Comandos CLI (futuro: Cobra)
├── pkg/
│   ├── calculator/        # Motor de cálculo
│   │   ├── plazo.go      # Lógica de cálculo
│   │   └── plazo_test.go # Tests unitarios
│   ├── exporter/         # Exportadores
│   │   └── csv.go        # Export a CSV
│   └── models/           # Estructuras de datos
│       └── types.go      # Tipos principales
├── main.go               # Entry point
├── go.mod                # Dependencias
└── README.md
```

## 💡 Ejemplo de Uso

```go
// Configurar plazos fijos
plazos := []models.PlazoFijo{
    models.NewPlazoFijo("360 días", 360, 6.5, 4390.69, 0),
    models.NewPlazoFijo("30 días", 30, 4.6, 500.00, 11),
}

// Configurar simulación
config := models.ConfiguracionSimulacion{
    FechaInicio:        time.Now(),
    Plazos:             plazos,
    AportesDisponibles: 500.00,
    AportePorPeriodo:   100.00,
    PlazoConAportes:    "30 días",
    NumAportes:         5,
}

// Simular
resultado := calculator.SimularTodosLosPlazo(config)

// Exportar
exporter.ExportarResumenCSV(resultado, "resultado.csv")
```

## 📊 Fórmula de Cálculo

Interés Simple:

```
Interés = Capital × (Tasa/100) × (Días/360)
```

Capital Final:

```
Capital Final = Capital + Interés + Aporte
```

## 🎯 Próximas Funcionalidades

* [ ] CLI completo con Cobra
* [ ] Interfaz TUI con Bubble Tea
* [ ] Export a Excel (xlsx)
* [ ] Comparación de escenarios
* [ ] Optimización automática de estrategia
* [ ] Aplicación desktop con Wails

## 🤝 Contribuir

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature
3. Commit tus cambios
4. Push a la rama
5. Abre un Pull Request

## 📝 Licencia

MIT License - ver LICENSE para detalles

## 👤 Autor

**csidevrik** - IT Administrator @ EMOV

---

Hecho con ❤️ y Go
