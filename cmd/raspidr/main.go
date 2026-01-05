package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
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
}
