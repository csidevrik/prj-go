package main

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	procFindWindow      = user32.NewProc("FindWindowW")
	procSetForeground   = user32.NewProc("SetForegroundWindow")
	procGetForeground   = user32.NewProc("GetForegroundWindow")
	procGetActiveWindow = user32.NewProc("GetActiveWindow")
	procKeybdEvent      = user32.NewProc("keybd_event")
)

const (
	VK_CONTROL      = 0x11
	VK_T            = 0x54
	VK_L            = 0x4C
	VK_RETURN       = 0x0D
	KEYEVENTF_KEYUP = 0x0002
)

func sendKey(vk byte) {
	procKeybdEvent.Call(uintptr(vk), 0, 0, 0)
	procKeybdEvent.Call(uintptr(vk), 0, KEYEVENTF_KEYUP, 0)
}

func sendCtrlCombo(vk byte) {
	procKeybdEvent.Call(VK_CONTROL, 0, 0, 0)
	procKeybdEvent.Call(uintptr(vk), 0, 0, 0)
	procKeybdEvent.Call(uintptr(vk), 0, KEYEVENTF_KEYUP, 0)
	procKeybdEvent.Call(VK_CONTROL, 0, KEYEVENTF_KEYUP, 0)
}

func main() {
	// Buscar ventana de Explorer
	hwnd, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("CabinetWClass"))), 0)
	if hwnd == 0 {
		return
	}
	procSetForeground.Call(hwnd)

	// Paths que quieres abrir
	paths := []string{
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PRINCIPALES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PARADAS_INTELIGENTES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\RADARES_HOUSING`,
	}

	for _, path := range paths {
		// Crear nueva pestaña
		sendCtrlCombo(VK_T)
		time.Sleep(500 * time.Millisecond)

		// Focus en barra de direcciones
		sendCtrlCombo(VK_L)
		time.Sleep(500 * time.Millisecond)

		// Escribir la ruta (simulación de teclado)
		for _, r := range path {
			procKeybdEvent.Call(uintptr(r), 0, 0, 0)
			procKeybdEvent.Call(uintptr(r), 0, KEYEVENTF_KEYUP, 0)
		}

		// Enter
		sendKey(VK_RETURN)
		time.Sleep(1000 * time.Millisecond)
	}
}
