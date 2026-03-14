# Documentyzer

Generador de documentación para proyectos escrito en Go. Extrae metadatos del proyecto y produce documentos usando plantillas Markdown.

## Índice

- Instalación
- Requisitos
- Construcción y ejecución
- Uso básico
- Estructura del proyecto
- Plantillas
- Contribuir
- Licencia

## Requisitos

- Go 1.18 o superior

## Instalación

Clona el repositorio y sitúate en la carpeta del proyecto:

```bash
git clone <repo-url>
cd prj-go/documentyzer
```

## Construcción y ejecución

Compilar el binario:

```bash
go build -o documentyzer-bin ./cmd
```

En Windows el ejecutable será `documentyzer-bin.exe` y puede ejecutarse así:

```powershell
.\documentyzer-bin.exe
```

Ejecutar directamente con `go run` durante desarrollo:

```bash
go run ./cmd
```

## Uso básico

La herramienta extrae metadatos del proyecto y genera documentación en formato Markdown usando las plantillas de la carpeta `templates`.

Ejemplos de uso (ajusta flags según `cmd/main.go`):

Generar documentación para el directorio actual y escribirla en `docs/`:

```bash
./documentyzer-bin -input ./ -output docs/
```

Mostrar ayuda y flags disponibles:

```bash
./documentyzer-bin -h
```

> Nota: Los flags y opciones exactas están definidos en [cmd/main.go](cmd/main.go). Revisa ese archivo para parámetros adicionales.

## Estructura del proyecto

- Código de entrada: [cmd/main.go](cmd/main.go)
- Lógica interna:
	- [internal/git/branches.go](internal/git/branches.go)
	- [internal/meta/meta.go](internal/meta/meta.go)
	- [internal/project/init.go](internal/project/init.go)
	- [internal/scanner/scanner.go](internal/scanner/scanner.go)
	- [internal/tui/tui.go](internal/tui/tui.go)
- Plantillas: [templates/project.md.tmpl](templates/project.md.tmpl), [templates/root.md.tmpl](templates/root.md.tmpl)

## Plantillas

Las plantillas usan la sintaxis del paquete `text/template` de Go. Modifica los archivos en `templates/` para personalizar el formato de salida.

## Desarrollo y pruebas

Formatear el código:

```bash
gofmt -w .
```

Ejecutar tests (si existieran):

```bash
go test ./...
```

## Contribuir

1. Abre un issue para discutir cambios importantes.
2. Crea una rama por feature: `git checkout -b feat/nueva-funcionalidad`.
3. Envía un pull request contra `main`.

## Licencia

Revisa el fichero de licencia en la raíz del repositorio: [../LICENSE](../LICENSE)

---

Si quieres, puedo actualizar los ejemplos de ejecución con los flags reales leyendo y extrayendo las opciones desde [cmd/main.go](cmd/main.go). ¿Quieres que lo haga?

