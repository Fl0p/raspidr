# RaspiDR

Go application for Raspberry Pi Zero 2W.

## Getting Started

### Build

```bash
# Local build
make build

# Build for Raspberry Pi (ARM64)
make build-pi
```

### Deploy

```bash
make deploy PI_HOST=raspberrypi.local PI_USER=pi
```

## Project Structure

```
cmd/raspidr/    - main application entry point
internal/       - private application packages
docs/           - documentation
```
