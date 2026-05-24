package main

import (
    "syscall"
    "unsafe"
)

var (
    user32           = syscall.NewLazyDLL("user32.dll")
    procFindWindow   = user32.NewProc("FindWindowW")
    procSetForeground= user32.NewProc("SetForegroundWindow")
    procKeybdEvent   = user32.NewProc("keybd_event")
)

func main() {
    // Buscar ventana de Explorer
    hwnd, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("CabinetWClass"))), 0)
    if hwnd != 0 {
        procSetForeground.Call(hwnd)

        // Simular Ctrl+T
        const VK_CONTROL = 0x11
        const VK_T = 0x54
        const KEYEVENTF_KEYUP = 0x0002

        procKeybdEvent.Call(VK_CONTROL, 0, 0, 0)
        procKeybdEvent.Call(VK_T, 0, 0, 0)
        procKeybdEvent.Call(VK_T, 0, KEYEVENTF_KEYUP, 0)
        procKeybdEvent.Call(VK_CONTROL, 0, KEYEVENTF_KEYUP, 0)
    }
}
