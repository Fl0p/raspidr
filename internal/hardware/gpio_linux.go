//go:build linux && (arm64 || arm)

package hardware

import (
	"fmt"
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

type gpioLinux struct {
	pins map[int]rpio.Pin
}

func newGPIO() GPIO {
	log.Println("[GPIO] Initializing real GPIO for Raspberry Pi")
	return &gpioLinux{
		pins: make(map[int]rpio.Pin),
	}
}

func (g *gpioLinux) Init() error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open GPIO: %w", err)
	}
	log.Println("[GPIO] GPIO initialized successfully")
	return nil
}

func (g *gpioLinux) Close() error {
	log.Println("[GPIO] Closing GPIO")
	return rpio.Close()
}

func (g *gpioLinux) SetMode(pin int, mode PinMode) error {
	p := rpio.Pin(pin)
	g.pins[pin] = p

	switch mode {
	case ModeInput:
		p.Input()
		log.Printf("[GPIO] Pin %d: set as INPUT", pin)
	case ModeOutput:
		p.Output()
		log.Printf("[GPIO] Pin %d: set as OUTPUT", pin)
	}
	return nil
}

func (g *gpioLinux) Write(pin int, state PinState) error {
	p, ok := g.pins[pin]
	if !ok {
		return fmt.Errorf("pin %d not configured", pin)
	}

	switch state {
	case High:
		p.High()
	case Low:
		p.Low()
	}
	return nil
}

func (g *gpioLinux) Read(pin int) (PinState, error) {
	p, ok := g.pins[pin]
	if !ok {
		return Low, fmt.Errorf("pin %d not configured", pin)
	}

	if p.Read() == rpio.High {
		return High, nil
	}
	return Low, nil
}

func (g *gpioLinux) Toggle(pin int) error {
	p, ok := g.pins[pin]
	if !ok {
		return fmt.Errorf("pin %d not configured", pin)
	}
	p.Toggle()
	return nil
}
