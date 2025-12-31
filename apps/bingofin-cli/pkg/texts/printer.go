package ui

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Style struct {
	UnderlineChar rune // '=' '-' etc.
	PadY          int  // líneas en blanco antes/después del título
	DoubleLine    bool // si true, imprime 2 líneas (una fina + otra fuerte)
	ThinChar      rune // '_' '-' etc. (solo si DoubleLine)
}

type Printer struct {
	w     io.Writer
	style Style
}

func New(w io.Writer) *Printer {
	return &Printer{
		w: w,
		style: Style{
			UnderlineChar: '=',
			ThinChar:      '_',
			PadY:          0,
			DoubleLine:    false,
		},
	}
}

func (p *Printer) WithStyle(s Style) *Printer {
	p.style = s
	return p
}

func (p *Printer) Blank(n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintln(p.w)
	}
}

func (p *Printer) Line(s string) {
	fmt.Fprintln(p.w, s)
}

func (p *Printer) Title(text string) {
	p.Blank(p.style.PadY)

	fmt.Fprintln(p.w, text)

	width := visibleWidth(text)

	if p.style.DoubleLine {
		fmt.Fprintln(p.w, strings.Repeat(string(p.style.ThinChar), width))
	}
	fmt.Fprintln(p.w, strings.Repeat(string(p.style.UnderlineChar), width))

	p.Blank(p.style.PadY)
}

func (p *Printer) Section(text string) {
	// Una sección es un título con línea simple por defecto
	old := p.style
	p.style.DoubleLine = false
	if p.style.UnderlineChar == 0 {
		p.style.UnderlineChar = '='
	}
	p.Title(text)
	p.style = old
}

// ---- helpers de ancho visible ----

// elimina códigos ANSI (si luego quieres colores)
var ansiRE = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func visibleWidth(s string) int {
	s = ansiRE.ReplaceAllString(s, "")

	// Simplificación práctica:
	// - contamos runas, pero evitamos contar variación/skin tone/joiners comunes
	// - emojis quedan como 1 “unidad” aprox (suficiente para subrayados bonitos)
	n := 0
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		switch r {
		case 0x200D, // zero-width joiner
			0xFE0F: // variation selector-16
			continue
		default:
			n++
		}
	}
	if n < 0 {
		return 0
	}
	return n
}
