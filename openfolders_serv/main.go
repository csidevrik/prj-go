package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	procFindWindow    = user32.NewProc("FindWindowW")
	procSetForeground = user32.NewProc("SetForegroundWindow")
	procSendInput     = user32.NewProc("SendInput")
)

const (
	INPUT_KEYBOARD  = 1
	KEYEVENTF_KEYUP = 0x0002
	VK_CONTROL      = 0x11
	VK_T            = 0x54
	VK_L            = 0x4C
	VK_RETURN       = 0x0D
)

type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
	_    [8]byte // Padding para alineación 64-bit
}

type KEYBDINPUT struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

func keyPress(vk uint16, down bool) {
	input := INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   KEYBDINPUT{Vk: vk},
	}
	if !down {
		input.Ki.Flags = KEYEVENTF_KEYUP
	}
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
}

// Combo correcto: mantiene Ctrl presionado mientras presiona la otra tecla
func sendCtrlCombo(vk uint16) {
	keyPress(VK_CONTROL, true) // Ctrl DOWN
	time.Sleep(100 * time.Millisecond)
	keyPress(vk, true) // Tecla DOWN
	time.Sleep(50 * time.Millisecond)
	keyPress(vk, false) // Tecla UP
	time.Sleep(50 * time.Millisecond)
	keyPress(VK_CONTROL, false) // Ctrl UP
}

// Escribe texto usando Unicode directo (funciona con cualquier layout de teclado)
func sendTextUnicode(text string) {
	const KEYEVENTF_UNICODE = 0x0004

	for _, r := range text {
		input := INPUT{
			Type: INPUT_KEYBOARD,
			Ki: KEYBDINPUT{
				Scan:  uint16(r),
				Flags: KEYEVENTF_UNICODE,
			},
		}
		// Key down
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		// Key up
		input.Ki.Flags = KEYEVENTF_UNICODE | KEYEVENTF_KEYUP
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		time.Sleep(5 * time.Millisecond)
	}
}

func main() {
	// ═══════════════════════════════════════════════════════════
	// 🎯 CONFIGURACIÓN: Rutas a abrir en pestañas
	// ═══════════════════════════════════════════════════════════
	paths := []string{
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PRINCIPALES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PARADAS_INTELIGENTES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\RADARES_HOUSING`,
	}

	// ═══════════════════════════════════════════════════════════
	// 🔍 Buscar ventana del Explorador
	// ═══════════════════════════════════════════════════════════
	hwnd, _, _ := procFindWindow.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("CabinetWClass"))),
		0,
	)

	if hwnd == 0 {
		fmt.Println("❌ No se encontró ventana del Explorador")
		fmt.Println("💡 Abre manualmente una ventana del Explorador y ejecuta de nuevo")
		return
	}

	// ═══════════════════════════════════════════════════════════
	// 🎯 Traer ventana al frente
	// ═══════════════════════════════════════════════════════════
	procSetForeground.Call(hwnd)
	time.Sleep(1000 * time.Millisecond)

	// ═══════════════════════════════════════════════════════════
	// 📁 Abrir pestañas
	// ═══════════════════════════════════════════════════════════
	fmt.Printf("🚀 Abriendo %d pestañas...\n", len(paths))

	for i, path := range paths {
		fmt.Printf("[%d/%d] %s\n", i+1, len(paths), path)

		// Ctrl+T (Nueva pestaña)
		sendCtrlCombo(VK_T)
		time.Sleep(1000 * time.Millisecond)

		// Ctrl+L (Barra de direcciones)
		sendCtrlCombo(VK_L)
		time.Sleep(800 * time.Millisecond)

		// Escribir ruta
		sendTextUnicode(path)
		time.Sleep(500 * time.Millisecond)

		// Enter
		keyPress(VK_RETURN, true)
		keyPress(VK_RETURN, false)
		time.Sleep(2000 * time.Millisecond)
	}

	fmt.Println("✅ Completado")
}
