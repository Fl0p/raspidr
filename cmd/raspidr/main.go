package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/fl0p/raspidr/internal/hardware"
)

const (
	// Example: LED on GPIO pin 17
	ledPin = 17
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== RaspiDR ===")
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println()

	// Initialize GPIO
	gpio := hardware.New()
	if err := gpio.Init(); err != nil {
		log.Fatalf("Failed to initialize GPIO: %v", err)
	}
	defer gpio.Close()

	// Configure LED pin as output
	if err := gpio.SetMode(ledPin, hardware.ModeOutput); err != nil {
		log.Fatalf("Failed to set pin mode: %v", err)
	}

	// Blink LED 3 times
	fmt.Println("Blinking LED on pin", ledPin)
	for i := 0; i < 3; i++ {
		if err := gpio.Write(ledPin, hardware.High); err != nil {
			log.Printf("Error writing to pin: %v", err)
		}
		time.Sleep(500 * time.Millisecond)

		if err := gpio.Write(ledPin, hardware.Low); err != nil {
			log.Printf("Error writing to pin: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Done!")
}
