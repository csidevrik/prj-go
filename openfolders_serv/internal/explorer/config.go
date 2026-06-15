package explorer

import "time"

const (
	// Window timings - Cuando se abre/enfoca Explorer
	WindowLaunchDelay   = 2000 * time.Millisecond
	WindowFocusDelay    = 1000 * time.Millisecond

	// Key combo timings - Para Ctrl+T y Ctrl+L
	NewTabDelay         = 1000 * time.Millisecond
	AddressBarDelay     = 800 * time.Millisecond

	// Text input timings - Cuando escribes el path
	TextInputDelay      = 1500 * time.Millisecond
	CharacterDelay      = 5 * time.Millisecond

	// Enter key timings - Cuando presionas Enter
	EnterKeyPressDelay  = 300 * time.Millisecond
	NavigationDelay     = 3000 * time.Millisecond
)
