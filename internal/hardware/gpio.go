package hardware

// PinMode defines the mode of a GPIO pin
type PinMode int

const (
	ModeInput PinMode = iota
	ModeOutput
)

// PinState defines the state of a GPIO pin
type PinState int

const (
	Low PinState = iota
	High
)

// GPIO defines the interface for GPIO operations
type GPIO interface {
	// Init initializes the GPIO subsystem
	Init() error

	// Close releases GPIO resources
	Close() error

	// SetMode sets the mode of a pin (input/output)
	SetMode(pin int, mode PinMode) error

	// Write sets the state of an output pin
	Write(pin int, state PinState) error

	// Read reads the state of an input pin
	Read(pin int) (PinState, error)

	// Toggle toggles the state of an output pin
	Toggle(pin int) error
}

// New creates a new GPIO instance appropriate for the current platform
func New() GPIO {
	return newGPIO()
}
