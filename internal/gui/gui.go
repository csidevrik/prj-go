// internal/gui/gui.go
package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartGUI() {
	a := app.New()
	w := a.NewWindow("Datanet")

	menu := container.NewVBox(
		widget.NewLabel("Menú"),
		widget.NewButton("Cargar CSV", func() {
			// Lógica para cargar CSV
		}),
	)

	w.SetContent(container.NewHSplit(menu, widget.NewLabel("Área principal")))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
