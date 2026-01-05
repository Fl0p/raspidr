# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

RaspiDR is a Go application designed to run on Raspberry Pi devices (specifically Raspberry Pi Zero 2W). The project uses cross-compilation to build ARM binaries from a development machine.

## Build Commands

```bash
# Build for local machine
make build

# Build for Raspberry Pi Zero 2W (64-bit ARM)
make build-pi

# Build for Raspberry Pi Zero 2W (32-bit ARM, if using 32-bit OS)
make build-pi32

# Run locally
make run
# or
go run .

# Deploy to Raspberry Pi (builds and copies via scp)
make deploy

# Clean build artifacts
make clean
```

## Deployment Configuration

The Makefile supports customizable deployment targets:
- `PI_HOST` - Raspberry Pi hostname (default: `raspberrypi.local`)
- `PI_USER` - SSH user (default: `pi`)
- `PI_PATH` - Deployment path (default: `/home/pi`)

Override with: `make deploy PI_HOST=mypi.local PI_USER=admin`

## Architecture

Single-file Go application (`main.go`) with no external dependencies. Binaries are output to `bin/` directory.
