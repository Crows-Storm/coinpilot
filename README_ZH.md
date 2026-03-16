<div align="center">

<img src="assets/logo.svg" alt="CoinPilot Logo" width="120" height="120">

# 🚀 CoinPilot

**离线加密货币投资组合管理CLI工具**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.25.6+-00ADD8?logo=go)](https://golang.org/)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey)](https://github.com/Crows-Storm/coinpilot/releases)
[![Release](https://img.shields.io/github/v/release/Crows-Storm/coinpilot)](https://github.com/Crows-Storm/coinpilot/releases)

*完全隐私保护的离线加密货币投资组合管理工具*

[📥 下载](https://github.com/Crows-Storm/coinpilot/releases) • [📖 使用文档](#使用方法) • [🐛 报告问题](https://github.com/Crows-Storm/coinpilot/issues) • [💡 功能建议](https://github.com/Crows-Storm/coinpilot/issues)

</div>

---

## ✨ 功能特性

🔒 **100% 离线** - 无需互联网连接，数据完全本地存储  
📊 **交易管理** - 记录买入/卖出交易并自动计算持仓  
🏢 **多交易所支持** - 跨不同交易所跟踪交易记录  
💰 **投资组合分析** - 实时持仓计算和盈亏分析  
📁 **CSV存储** - 简单可读的数据存储格式，完全可控  
🖥️ **跨平台支持** - 原生支持Windows、macOS和Linux  
⚡ **CLI界面** - 基于Cobra构建的快速直观命令行界面  

## 🚀 快速开始

### 安装

#### 方式一：下载预编译二进制文件（推荐）

访问我们的[发布页面](https://github.com/Crows-Storm/coinpilot/releases)并下载适合您系统的二进制文件：

- **Windows**: `coinpilot-windows-x64.exe`
- **macOS**: `coinpilot-macos-universal`（Intel + Apple Silicon通用版）
- **Linux**: `coinpilot-linux-x64`

#### 方式二：从源码构建

```bash
git clone https://github.com/Crows-Storm/coinpilot.git
cd coinpilot
go mod tidy
go build -o coinpilot cmd/coinpilot/main.go
```

### 第一步

```bash
# 添加您的第一笔交易
./coinpilot trade add --type BUY --symbol BTC --quantity 0.1 --price 45000

# 查看投资组合
./coinpilot position list

# 列出所有交易
./coinpilot trade list

# 获取帮助
./coinpilot --help
```

## 📖 使用方法

### 交易管理

#### 添加交易
```bash
# 购买比特币
./coinpilot trade add -t BUY -s BTC -q 0.1 -p 50000 -e binance

# 出售以太坊并包含手续费
./coinpilot trade add -t SELL -s ETH -q 2.0 -p 3000 -e binance -f 5.0
```

#### 查看交易
```bash
# 列出所有交易
./coinpilot trade list

# 删除特定交易
./coinpilot trade delete <交易ID>
```

### 投资组合管理

```bash
# 查看当前持仓
./coinpilot position list

# 获取任何命令的详细帮助
./coinpilot trade --help
./coinpilot position --help
```

### 支持的交易类型

| 类型 | 描述 |
|------|------|
| `BUY` | 购买加密货币 |
| `SELL` | 出售加密货币 |
| `TRANSFER_IN` | 转入交易所/钱包 |
| `TRANSFER_OUT` | 转出交易所/钱包 |
| `FEE` | 交易或网络手续费 |

### 命令参考

```bash
coinpilot trade add [flags]    # 添加新交易
coinpilot trade list           # 列出所有交易
coinpilot trade delete <id>    # 删除交易
coinpilot position list        # 显示当前持仓
coinpilot --version           # 显示版本信息
coinpilot --help              # 显示帮助
```

## 🏗️ 开发

### 构建不同平台版本

```bash
# 构建当前平台版本
make build

# 构建所有平台版本
make build-all

# 构建macOS通用二进制文件（仅限macOS）
make build-macos-universal
```

#### 手动跨平台构建

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

### 项目结构

```
coinpilot/
├── cmd/coinpilot/          # 应用程序入口点
├── internal/               # 私有应用程序代码
│   ├── cli/               # 命令行界面
│   ├── models/            # 数据模型
│   ├── services/          # 业务逻辑
│   └── filehandler/       # 文件操作
├── pkg/utils/             # 公共工具函数
├── build/                 # 构建产物
└── scripts/               # 构建和发布脚本
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行测试并生成覆盖率报告
make test-coverage

# 格式化代码
make fmt
```

## 📁 数据存储

所有数据都本地存储在当前目录的CSV文件中：

- **`trades.csv`** - 包含完整交易历史的交易记录
- **`positions.csv`** - 计算的持仓（自动生成的缓存）
- **`prices.csv`** - 手动价格条目（未来功能）
- **`config.yaml`** - 配置设置（未来功能）

### 数据隐私

- ✅ **无云存储** - 一切都保存在您的机器上
- ✅ **无需API密钥** - 不需要交易所连接
- ✅ **人类可读格式** - 可以在Excel中打开的CSV文件
- ✅ **轻松备份** - 只需复制CSV文件
- ✅ **便携性** - 可以将数据移动到任何地方

## 🤝 贡献

我们欢迎贡献！请查看我们的[贡献指南](CONTRIBUTING.md)了解详情。

### 开发环境设置

1. **Fork仓库**
2. **克隆您的fork**: `git clone https://github.com/yourusername/coinpilot.git`
3. **创建分支**: `git checkout -b feature/amazing-feature`
4. **进行更改并测试**: `make test`
5. **提交**: `git commit -m 'Add amazing feature'`
6. **推送**: `git push origin feature/amazing-feature`
7. **打开Pull Request**

## 📋 系统要求

- **Go 1.25.6+** 用于从源码构建
- **无运行时依赖** - 包含静态二进制文件

## 🗺️ 路线图

- [x] **阶段1**: 基础交易记录和持仓计算
- [ ] **阶段2**: 多交易所支持、高级分析、CSV导入/导出
- [ ] **阶段3**: 性能优化、高级报告、Web仪表板

查看我们的[项目路线图](https://github.com/Crows-Storm/coinpilot/projects)了解详细进展。

## 📄 许可证

本项目采用MIT许可证 - 查看[LICENSE](LICENSE)文件了解详情。

## 🙏 致谢

- 使用[Cobra](https://github.com/spf13/cobra)作为CLI框架构建
- 受隐私导向投资组合管理需求启发
- 感谢所有贡献者，特别是[Crows-Storm](https://github.com/Crows-Storm)

## 📞 支持

- 📖 **文档**: 查看此README和命令帮助（`--help`）
- 🐛 **错误报告**: [GitHub Issues](https://github.com/Crows-Storm/coinpilot/issues)
- 💡 **功能请求**: [GitHub Issues](https://github.com/Crows-Storm/coinpilot/issues)
- 💬 **讨论**: [GitHub Discussions](https://github.com/Crows-Storm/coinpilot/discussions)

---

<div align="center">

**⭐ 如果您觉得这个项目有用，请给我们一个星标！**

[English Version](README.md) | [发布说明](https://github.com/Crows-Storm/coinpilot/releases) | [贡献指南](CONTRIBUTING.md)

</div>