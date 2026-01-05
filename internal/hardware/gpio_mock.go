//go:build !linux || (!arm64 && !arm)

package hardware

import (
	"fmt"
	"log"
)

// gpioMock is a mock GPIO implementation for development on non-Pi platforms
type gpioMock struct {
	pinStates map[int]PinState
	pinModes  map[int]PinMode
}

func newGPIO() GPIO {
	log.Println("[GPIO-MOCK] Using mock GPIO (not running on Raspberry Pi)")
	return &gpioMock{
		pinStates: make(map[int]PinState),
		pinModes:  make(map[int]PinMode),
	}
}

func (g *gpioMock) Init() error {
	log.Println("[GPIO-MOCK] Init called - simulating GPIO initialization")
	return nil
}

func (g *gpioMock) Close() error {
	log.Println("[GPIO-MOCK] Close called - simulating GPIO cleanup")
	return nil
}

func (g *gpioMock) SetMode(pin int, mode PinMode) error {
	modeStr := "INPUT"
	if mode == ModeOutput {
		modeStr = "OUTPUT"
	}
	log.Printf("[GPIO-MOCK] Pin %d: SetMode(%s)", pin, modeStr)
	g.pinModes[pin] = mode
	return nil
}

func (g *gpioMock) Write(pin int, state PinState) error {
	if g.pinModes[pin] != ModeOutput {
		return fmt.Errorf("pin %d is not configured as output", pin)
	}
	stateStr := "LOW"
	if state == High {
		stateStr = "HIGH"
	}
	log.Printf("[GPIO-MOCK] Pin %d: Write(%s)", pin, stateStr)
	g.pinStates[pin] = state
	return nil
}

func (g *gpioMock) Read(pin int) (PinState, error) {
	if g.pinModes[pin] != ModeInput {
		return Low, fmt.Errorf("pin %d is not configured as input", pin)
	}
	state := g.pinStates[pin]
	stateStr := "LOW"
	if state == High {
		stateStr = "HIGH"
	}
	log.Printf("[GPIO-MOCK] Pin %d: Read() -> %s", pin, stateStr)
	return state, nil
}

func (g *gpioMock) Toggle(pin int) error {
	if g.pinModes[pin] != ModeOutput {
		return fmt.Errorf("pin %d is not configured as output", pin)
	}
	current := g.pinStates[pin]
	newState := Low
	if current == Low {
		newState = High
	}
	log.Printf("[GPIO-MOCK] Pin %d: Toggle() %v -> %v", pin, current, newState)
	return g.Write(pin, newState)
}
