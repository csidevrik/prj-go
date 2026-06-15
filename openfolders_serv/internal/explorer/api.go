package explorer

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
	_    [8]byte
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
	time.Sleep(10 * time.Millisecond)
}

func sendCtrlCombo(vk uint16) {
	keyPress(VK_CONTROL, true)
	time.Sleep(100 * time.Millisecond)
	keyPress(vk, true)
	time.Sleep(50 * time.Millisecond)
	keyPress(vk, false)
	time.Sleep(50 * time.Millisecond)
	keyPress(VK_CONTROL, false)
}

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
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		input.Ki.Flags = KEYEVENTF_UNICODE | KEYEVENTF_KEYUP
		procSendInput.Call(1, uintptr(unsafe.Pointer(&input)), unsafe.Sizeof(input))
		time.Sleep(5 * time.Millisecond)
	}
}

func OpenExplorerWithPaths(paths []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("no paths provided")
	}

	hwnd, _, _ := procFindWindow.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("CabinetWClass"))),
		0,
	)

	if hwnd == 0 {
		return fmt.Errorf("explorer window not found - please open a File Explorer window first")
	}

	procSetForeground.Call(hwnd)
	time.Sleep(1000 * time.Millisecond)

	for i, path := range paths {
		fmt.Printf("[%d/%d] Opening %s\n", i+1, len(paths), path)

		sendCtrlCombo(VK_T)
		time.Sleep(1000 * time.Millisecond)

		sendCtrlCombo(VK_L)
		time.Sleep(800 * time.Millisecond)

		sendTextUnicode(path)
		time.Sleep(1500 * time.Millisecond)

		keyPress(VK_RETURN, true)
		time.Sleep(300 * time.Millisecond)
		keyPress(VK_RETURN, false)
		time.Sleep(3000 * time.Millisecond)
	}

	return nil
}
