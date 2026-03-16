# 发布指南 / Release Guide

本文档说明如何为CoinPilot创建GitHub Release。

## 方法一：使用自动化脚本（推荐）

### 前提条件

1. 安装GitHub CLI：
   ```bash
   # macOS
   brew install gh
   
   # Windows (使用Chocolatey)
   choco install gh
   
   # 或者从官网下载：https://cli.github.com/
   ```

2. 登录GitHub CLI：
   ```bash
   gh auth login
   ```

### 创建Release

1. 确保所有更改已提交并推送到main分支

2. 运行发布脚本：
   ```bash
   # 基本用法
   ./scripts/release.sh v1.0.0
   
   # 带发布说明
   ./scripts/release.sh v1.0.0 "初始版本，包含基本交易功能"
   ```

脚本会自动：
- 构建所有平台的二进制文件
- 创建GitHub Release
- 上传所有二进制文件
- 设置适当的文件描述

## 方法二：手动创建Release

### 1. 构建二进制文件

```bash
# 构建所有平台
make build-all

# 如果在macOS上，创建通用二进制文件
make build-macos-universal
```

### 2. 在GitHub上创建Release

1. 访问 https://github.com/Crows-Storm/coinpilot/releases
2. 点击 "Create a new release"
3. 填写以下信息：
   - **Tag version**: `v1.0.0` (遵循语义化版本)
   - **Release title**: `CoinPilot v1.0.0`
   - **Description**: 描述此版本的新功能和改进

### 3. 上传二进制文件

将以下文件拖拽到Release页面的"Attach binaries"区域：

#### Windows版本
- `build/coinpilot-windows-amd64.exe` → 重命名为 `coinpilot-windows-x64.exe`
- `build/coinpilot-windows-386.exe` → 重命名为 `coinpilot-windows-x86.exe`

#### macOS版本
- `build/coinpilot-macos-universal` → 重命名为 `coinpilot-macos-universal`
- `build/coinpilot-macos-intel` → 重命名为 `coinpilot-macos-intel`
- `build/coinpilot-macos-arm64` → 重命名为 `coinpilot-macos-arm64`

#### Linux版本
- `build/coinpilot-linux-amd64` → 重命名为 `coinpilot-linux-x64`
- `build/coinpilot-linux-arm64` → 重命名为 `coinpilot-linux-arm64`

### 4. 发布

点击 "Publish release" 完成发布。

## 版本命名规范

使用语义化版本控制 (Semantic Versioning)：

- `v1.0.0` - 主要版本（重大更改）
- `v1.1.0` - 次要版本（新功能）
- `v1.0.1` - 补丁版本（bug修复）

## 发布说明模板

```markdown
## 🚀 新功能
- 添加了交易记录功能
- 支持多交易所管理

## 🐛 Bug修复
- 修复了CSV文件读取问题
- 改进了错误处理

## 📦 下载

### Windows
- [Windows 64位](链接) - 推荐大多数Windows用户
- [Windows 32位](链接) - 适用于较老的系统

### macOS
- [macOS通用版](链接) - 同时支持Intel和Apple Silicon
- [macOS Intel版](链接) - 仅适用于Intel Mac
- [macOS Apple Silicon版](链接) - 仅适用于M1/M2 Mac

### Linux
- [Linux 64位](链接) - 适用于大多数Linux发行版
- [Linux ARM64](链接) - 适用于ARM64架构

## 📋 安装说明

下载对应平台的文件后：

1. **Windows**: 双击运行或在命令行中使用
2. **macOS**: 可能需要在"系统偏好设置 > 安全性与隐私"中允许运行
3. **Linux**: 添加执行权限 `chmod +x coinpilot-linux-x64`

## 🔄 更新

如果你已经安装了之前的版本，只需下载新版本替换旧文件即可。
```

## 自动化CI/CD（未来改进）

考虑设置GitHub Actions来自动化发布流程：

1. 当推送新tag时自动触发
2. 自动构建所有平台
3. 自动创建Release并上传文件
4. 自动生成changelog

这将在项目成熟后实施。

---

[English Version](RELEASE.md)