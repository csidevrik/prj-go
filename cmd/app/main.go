package main

import (
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const scale = 0.2

func verticalLabel(text string, x, y float32) *canvas.Image {
	img := image.NewRGBA(image.Rect(0, 0, 60, 200))

	// fondo transparente
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	// simular texto con líneas
	for i := 0; i < 6; i++ {
		rect := image.Rect(20, 10+i*25, 40, 15+i*25)
		draw.Draw(img, rect, &image.Uniform{color.Black}, image.Point{}, draw.Src)
	}

	cImg := canvas.NewImageFromImage(img)
	cImg.Move(fyne.NewPos(x, y))
	cImg.Resize(fyne.NewSize(40, 200))

	return cImg
}

func main() {
	a := app.NewWithID("woodero")
	w := a.NewWindow("Woodero - Canvas Base")

	s := float32(scale)

	// dimensiones reales
	sheetW := float32(2450)
	sheetH := float32(1230)

	// contenedor libre (tipo CAD)
	root := container.NewWithoutLayout()

	// 🪵 PLANCHA
	sheet := canvas.NewRectangle(color.NRGBA{R: 240, G: 240, B: 240, A: 255})
	sheet.Move(fyne.NewPos(50, 50)) // margen
	sheet.Resize(fyne.NewSize(sheetW*s, sheetH*s))

	root.Add(sheet)

	// 📏 COTA SUPERIOR (2450 mm)
	topText := canvas.NewText(strconv.Itoa(int(sheetW))+" mm", color.Black)
	topText.Move(fyne.NewPos(50+(sheetW*s)/2-40, 20))

	root.Add(topText)

	// 📏 COTA LATERAL (1230 mm) - versión limpia
	sideLabel := strconv.Itoa(int(sheetH)) + " mm"

	// posición base (centrado verticalmente)
	startX := 50 + (sheetW * s) + 40
	startY := 50 + (sheetH*s)/2 - float32(len(sideLabel))*7

	for i, ch := range sideLabel {
		t := canvas.NewText(string(ch), color.Black)
		t.Move(fyne.NewPos(startX, startY+float32(i*14)))
		root.Add(t)
	}
	// w.SetContent(sideLabel)

	w.SetContent(root)
	w.Resize(fyne.NewSize(800, 400))
	w.ShowAndRun()
}
