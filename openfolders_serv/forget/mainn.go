package main

import (
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
	VK_CONTROL      = 0x11
	VK_T            = 0x54
	VK_L            = 0x4C
	VK_RETURN       = 0x0D
	KEYEVENTF_KEYUP = 0x0002
)

type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

func sendKey(vk uint16) {
	var input INPUT
	input.Type = 1 // INPUT_KEYBOARD
	input.Ki = KEYBDINPUT{Vk: vk}
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))

	input.Ki.Flags = KEYEVENTF_KEYUP
	procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
}

func sendCtrlCombo(vk uint16) {
	sendKey(VK_CONTROL)
	sendKey(vk)
	sendKey(VK_CONTROL | KEYEVENTF_KEYUP)
}

func sendText(text string) {
	for _, r := range text {
		var input INPUT
		input.Type = 1
		input.Ki = KEYBDINPUT{Vk: uint16(r)}
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))

		input.Ki.Flags = KEYEVENTF_KEYUP
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
	}
}

func main() {
	hwnd, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("CabinetWClass"))), 0)
	if hwnd == 0 {
		return
	}
	procSetForeground.Call(hwnd)

	paths := []string{
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PRINCIPALES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\PARADAS_INTELIGENTES`,
		`D:\C\OneDriveP\OneDrive\2A-JOB02-EMOVEP\2026\CONTRATOS\PROYECTOS\SERVICE_ENLACES_CONSOLIDACION\INFO\RADARES_HOUSING`,
	}

	for _, path := range paths {
		// Nueva pestaña
		sendCtrlCombo(VK_T)
		time.Sleep(500 * time.Millisecond)

		// Barra de direcciones
		sendCtrlCombo(VK_L)
		time.Sleep(500 * time.Millisecond)

		// Escribir ruta completa
		sendText(path)
		time.Sleep(500 * time.Millisecond)

		// Enter
		sendKey(VK_RETURN)
		time.Sleep(1000 * time.Millisecond)
	}
}
