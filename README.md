<div align="center">

<img src="assets/logo.svg" alt="CoinPilot Logo" width="120" height="120">

# 🚀 CoinPilot

**Offline Cryptocurrency Portfolio Management CLI Tool**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.25.6+-00ADD8?logo=go)](https://golang.org/)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey)](https://github.com/Crows-Storm/coinpilot/releases)
[![Release](https://img.shields.io/github/v/release/Crows-Storm/coinpilot)](https://github.com/Crows-Storm/coinpilot/releases)

*Take control of your cryptocurrency portfolio with complete privacy and offline operation*

[📥 Download](https://github.com/Crows-Storm/coinpilot/releases) • [📖 Documentation](#usage) • [🐛 Report Bug](https://github.com/Crows-Storm/coinpilot/issues) • [💡 Request Feature](https://github.com/Crows-Storm/coinpilot/issues)

</div>

---

## ✨ Features

🔒 **100% Offline** - No internet connectivity required, your data stays local  
📊 **Trade Management** - Record BUY/SELL trades with automatic position calculation  
🏢 **Multi-Exchange Support** - Track trades across different exchanges  
💰 **Portfolio Analytics** - Real-time position calculations and P&L analysis  
📁 **CSV Storage** - Simple, readable data storage format you control  
🖥️ **Cross-Platform** - Native binaries for Windows, macOS, and Linux  
⚡ **CLI Interface** - Fast, intuitive command-line interface built with Cobra  

## 🚀 Quick Start

### Installation

#### Option 1: Download Pre-built Binaries (Recommended)

Visit our [Releases page](https://github.com/Crows-Storm/coinpilot/releases) and download the appropriate binary for your system:

- **Windows**: `coinpilot-windows-x64.exe`
- **macOS**: `coinpilot-macos-universal` (Intel + Apple Silicon)
- **Linux**: `coinpilot-linux-x64`

#### Option 2: Build from Source

```bash
git clone https://github.com/Crows-Storm/coinpilot.git
cd coinpilot
go mod tidy
go build -o coinpilot cmd/coinpilot/main.go
```

### First Steps

```bash
# Add your first trade
./coinpilot trade add --type BUY --symbol BTC --quantity 0.1 --price 45000

# View your portfolio
./coinpilot position list

# List all trades
./coinpilot trade list

# Get help
./coinpilot --help
```

### Build from Source

```bash
git clone https://github.com/Crows-Storm/coinpilot.git
cd coinpilot
go mod tidy
go build -o coinpilot cmd/coinpilot/main.go
```

## 📖 Usage

### Trade Management

#### Adding Trades
```bash
# Buy Bitcoin
./coinpilot trade add -t BUY -s BTC -q 0.1 -p 50000 -e binance

# Sell Ethereum with fee
./coinpilot trade add -t SELL -s ETH -q 2.0 -p 3000 -e binance -f 5.0
```

#### Viewing Trades
```bash
# List all trades
./coinpilot trade list

# Delete a specific trade
./coinpilot trade delete <trade-id>
```

### Portfolio Management

```bash
# View current positions
./coinpilot position list

# Get detailed help for any command
./coinpilot trade --help
./coinpilot position --help
```

### Supported Trade Types

| Type | Description |
|------|-------------|
| `BUY` | Purchase cryptocurrency |
| `SELL` | Sell cryptocurrency |
| `TRANSFER_IN` | Transfer into exchange/wallet |
| `TRANSFER_OUT` | Transfer out of exchange/wallet |
| `FEE` | Trading or network fees |

### Command Reference

```bash
coinpilot trade add [flags]    # Add a new trade
coinpilot trade list           # List all trades  
coinpilot trade delete <id>    # Delete a trade
coinpilot position list        # Show current positions
coinpilot --version           # Show version info
coinpilot --help              # Show help
```

## 🏗️ Development

### Building for Different Platforms

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Build macOS universal binary (macOS only)
make build-macos-universal
```

#### Manual Cross-Platform Builds

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o coinpilot.exe cmd/coinpilot/main.go

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o coinpilot-macos-intel cmd/coinpilot/main.go

# macOS Apple Silicon  
GOOS=darwin GOARCH=arm64 go build -o coinpilot-macos-arm64 cmd/coinpilot/main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o coinpilot-linux cmd/coinpilot/main.go
```

### Project Structure

```
coinpilot/
├── cmd/coinpilot/          # Application entry point
├── internal/               # Private application code
│   ├── cli/               # Command-line interface
│   ├── models/            # Data models
│   ├── services/          # Business logic
│   └── filehandler/       # File operations
├── pkg/utils/             # Public utility functions
├── build/                 # Build artifacts
└── scripts/               # Build and release scripts
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
make test-coverage

# Format code
make fmt
```

## 📁 Data Storage

All data is stored locally in CSV files in your current directory:

- **`trades.csv`** - Trade records with full transaction history
- **`positions.csv`** - Calculated positions (auto-generated cache)
- **`prices.csv`** - Manual price entries (future feature)
- **`config.yaml`** - Configuration settings (future feature)

### Data Privacy

- ✅ **No cloud storage** - Everything stays on your machine
- ✅ **No API keys required** - No exchange connections needed  
- ✅ **Human-readable format** - CSV files you can open in Excel
- ✅ **Easy backup** - Just copy the CSV files
- ✅ **Portable** - Move your data anywhere

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Setup

1. **Fork the repository**
2. **Clone your fork**: `git clone https://github.com/yourusername/coinpilot.git`
3. **Create a branch**: `git checkout -b feature/amazing-feature`
4. **Make changes and test**: `make test`
5. **Commit**: `git commit -m 'Add amazing feature'`
6. **Push**: `git push origin feature/amazing-feature`
7. **Open a Pull Request**

## 📋 Requirements

- **Go 1.25.6+** for building from source
- **No runtime dependencies** - static binaries included

## 🗺️ Roadmap

- [x] **Phase 1**: Basic trade recording and position calculation
- [ ] **Phase 2**: Multi-exchange support, advanced analytics, CSV import/export
- [ ] **Phase 3**: Performance optimization, advanced reporting, web dashboard

See our [Project Roadmap](https://github.com/Crows-Storm/coinpilot/projects) for detailed progress.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Inspired by the need for privacy-focused portfolio management
- Thanks to all contributors, especially [Crows-Storm](https://github.com/Crows-Storm)

## 📞 Support

- 📖 **Documentation**: Check this README and command help (`--help`)
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/Crows-Storm/coinpilot/issues)
- 💡 **Feature Requests**: [GitHub Issues](https://github.com/Crows-Storm/coinpilot/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/Crows-Storm/coinpilot/discussions)

---

<div align="center">

**⭐ Star this project if you find it useful!**

[中文版本](README_ZH.md) | [Release Notes](https://github.com/Crows-Storm/coinpilot/releases) | [Contributing](CONTRIBUTING.md)

</div>