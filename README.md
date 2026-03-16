# CoinPilot

CoinPilot is an offline cryptocurrency portfolio management CLI tool for recording trades, calculating positions, analyzing returns, and displaying asset dashboards.

## Features

- **Offline Operation**: No internet connectivity required
- **Trade Management**: Record BUY/SELL trades with automatic position calculation
- **Multi-Exchange Support**: Track trades across different exchanges
- **CSV Storage**: Simple, readable data storage format
- **CLI Interface**: Intuitive command-line interface using Cobra

## Installation

### Prerequisites

- Go 1.25.6 or later

### Build from Source

```bash
git clone https://github.com/Crows-Storm/coinpilot.git
cd coinpilot
go mod tidy
go build -o coinpilot cmd/coinpilot/main.go
```

### Cross-Platform Build

#### Build for Windows (from any platform)
```bash
# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o coinpilot.exe cmd/coinpilot/main.go

# Windows 32-bit
GOOS=windows GOARCH=386 go build -o coinpilot-32.exe cmd/coinpilot/main.go
```

#### Build for macOS (from any platform)
```bash
# macOS Intel (x64)
GOOS=darwin GOARCH=amd64 go build -o coinpilot-macos-intel cmd/coinpilot/main.go

# macOS Apple Silicon (ARM64)
GOOS=darwin GOARCH=arm64 go build -o coinpilot-macos-arm64 cmd/coinpilot/main.go

# Universal macOS binary (requires macOS)
lipo -create -output coinpilot-macos-universal coinpilot-macos-intel coinpilot-macos-arm64
```

#### Build for Linux (from any platform)
```bash
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o coinpilot-linux cmd/coinpilot/main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o coinpilot-linux-arm64 cmd/coinpilot/main.go
```

#### Build All Platforms at Once
```bash
# Use the provided Makefile
make build-all
```

## Usage

### Basic Commands

```bash
# Add a trade
./coinpilot trade add --type BUY --symbol BTC --quantity 0.5 --price 45000

# List all trades
./coinpilot trade list

# List positions
./coinpilot position list

# Delete a trade
./coinpilot trade delete <trade-id>

# Get help
./coinpilot --help
```

### Trade Types

- `BUY`: Purchase cryptocurrency
- `SELL`: Sell cryptocurrency
- `TRANSFER_IN`: Transfer into exchange/wallet
- `TRANSFER_OUT`: Transfer out of exchange/wallet
- `FEE`: Trading or network fees

### Examples

```bash
# Buy Bitcoin
./coinpilot trade add -t BUY -s BTC -q 0.1 -p 50000 -e binance

# Sell Ethereum
./coinpilot trade add -t SELL -s ETH -q 2.0 -p 3000 -e binance -f 5.0

# View all positions
./coinpilot position list
```

## Data Storage

All data is stored locally in CSV files:
- `trades.csv`: Trade records
- `positions.csv`: Calculated positions (cached)
- `prices.csv`: Manual price entries
- `config.yaml`: Configuration settings

## Development

This project follows Go best practices with a clean architecture:

```
cmd/           # Application entry points
internal/      # Private application code
  cli/         # Command-line interface
  models/      # Data models
  services/    # Business logic
pkg/           # Public library code
  utils/       # Utility functions
```

## License

[License information to be added]

---

[中文版本](README_zh.md)