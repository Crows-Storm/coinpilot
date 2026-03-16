# CoinPilot

CoinPilot 是一个离线加密货币投资组合管理CLI工具，用于记录交易、计算持仓、分析收益和显示资产仪表板。

## 功能特性

- **离线操作**: 无需互联网连接
- **交易管理**: 记录买入/卖出交易并自动计算持仓
- **多交易所支持**: 跨不同交易所跟踪交易
- **CSV存储**: 简单、可读的数据存储格式
- **CLI界面**: 使用Cobra构建的直观命令行界面

## 安装

### 系统要求

- Go 1.25.6 或更高版本

### 从源码构建

```bash
git clone https://github.com/Crows-Storm/coinpilot.git
cd coinpilot
go mod tidy
go build -o coinpilot cmd/coinpilot/main.go
```

### 跨平台构建

#### 构建Windows版本（从任何平台）
```bash
# Windows 64位
GOOS=windows GOARCH=amd64 go build -o coinpilot.exe cmd/coinpilot/main.go

# Windows 32位
GOOS=windows GOARCH=386 go build -o coinpilot-32.exe cmd/coinpilot/main.go
```

#### 构建macOS版本（从任何平台）
```bash
# macOS Intel (x64)
GOOS=darwin GOARCH=amd64 go build -o coinpilot-macos-intel cmd/coinpilot/main.go

# macOS Apple Silicon (ARM64)
GOOS=darwin GOARCH=arm64 go build -o coinpilot-macos-arm64 cmd/coinpilot/main.go

# 通用macOS二进制文件（需要macOS环境）
lipo -create -output coinpilot-macos-universal coinpilot-macos-intel coinpilot-macos-arm64
```

#### 构建Linux版本（从任何平台）
```bash
# Linux 64位
GOOS=linux GOARCH=amd64 go build -o coinpilot-linux cmd/coinpilot/main.go

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o coinpilot-linux-arm64 cmd/coinpilot/main.go
```

#### 一次构建所有平台
```bash
# 使用提供的Makefile
make build-all
```

## 使用方法

### 基本命令

```bash
# 添加交易
./coinpilot trade add --type BUY --symbol BTC --quantity 0.5 --price 45000

# 列出所有交易
./coinpilot trade list

# 列出持仓
./coinpilot position list

# 删除交易
./coinpilot trade delete <交易ID>

# 获取帮助
./coinpilot --help
```

### 交易类型

- `BUY`: 购买加密货币
- `SELL`: 出售加密货币
- `TRANSFER_IN`: 转入交易所/钱包
- `TRANSFER_OUT`: 转出交易所/钱包
- `FEE`: 交易或网络手续费

### 使用示例

```bash
# 购买比特币
./coinpilot trade add -t BUY -s BTC -q 0.1 -p 50000 -e binance

# 出售以太坊
./coinpilot trade add -t SELL -s ETH -q 2.0 -p 3000 -e binance -f 5.0

# 查看所有持仓
./coinpilot position list
```

## 数据存储

所有数据都本地存储在CSV文件中：
- `trades.csv`: 交易记录
- `positions.csv`: 计算的持仓（缓存）
- `prices.csv`: 手动价格条目
- `config.yaml`: 配置设置

## 开发

本项目遵循Go最佳实践，采用清晰的架构：

```
cmd/           # 应用程序入口点
internal/      # 私有应用程序代码
  cli/         # 命令行界面
  models/      # 数据模型
  services/    # 业务逻辑
pkg/           # 公共库代码
  utils/       # 工具函数
```

## 许可证

[许可证信息待添加]

---

[English Version](README.md)