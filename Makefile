.PHONY: build build-pi run clean

BINARY_NAME=raspidr
PI_HOST?=raspberrypi.local
PI_USER?=pi
PI_PATH?=/home/pi

# Build for local machine
build:
	go build -o bin/$(BINARY_NAME) .

# Build for Raspberry Pi Zero 2W (ARM64)
build-pi:
	GOOS=linux GOARCH=arm64 go build -o bin/$(BINARY_NAME)-linux-arm64 .

# Build for Raspberry Pi Zero 2W (32-bit, if using 32-bit OS)
build-pi32:
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/$(BINARY_NAME)-linux-arm .

# Run locally
run:
	go run .

# Deploy to Raspberry Pi
deploy: build-pi
	scp bin/$(BINARY_NAME)-linux-arm64 $(PI_USER)@$(PI_HOST):$(PI_PATH)/$(BINARY_NAME)
	ssh $(PI_USER)@$(PI_HOST) "chmod +x $(PI_PATH)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	rm -rf bin/
