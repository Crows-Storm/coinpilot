# Release Guide

This document explains how to create GitHub Releases for CoinPilot.

## Method 1: Automated Script (Recommended)

### Prerequisites

1. Install GitHub CLI:
   ```bash
   # macOS
   brew install gh
   
   # Windows (using Chocolatey)
   choco install gh
   
   # Or download from: https://cli.github.com/
   ```

2. Login to GitHub CLI:
   ```bash
   gh auth login
   ```

### Create Release

1. Ensure all changes are committed and pushed to main branch

2. Run the release script:
   ```bash
   # Basic usage
   ./scripts/release.sh v1.0.0
   
   # With release notes
   ./scripts/release.sh v1.0.0 "Initial release with basic trading functionality"
   ```

The script will automatically:
- Build binaries for all platforms
- Create GitHub Release
- Upload all binary files
- Set appropriate file descriptions

## Method 2: Manual Release Creation

### 1. Build Binaries

```bash
# Build for all platforms
make build-all

# If on macOS, create universal binary
make build-macos-universal
```

### 2. Create Release on GitHub

1. Visit https://github.com/Crows-Storm/coinpilot/releases
2. Click "Create a new release"
3. Fill in the following information:
   - **Tag version**: `v1.0.0` (follow semantic versioning)
   - **Release title**: `CoinPilot v1.0.0`
   - **Description**: Describe new features and improvements in this version

### 3. Upload Binary Files

Drag and drop the following files to the "Attach binaries" area on the Release page:

#### Windows Versions
- `build/coinpilot-windows-amd64.exe` → Rename to `coinpilot-windows-x64.exe`
- `build/coinpilot-windows-386.exe` → Rename to `coinpilot-windows-x86.exe`

#### macOS Versions
- `build/coinpilot-macos-universal` → Rename to `coinpilot-macos-universal`
- `build/coinpilot-macos-intel` → Rename to `coinpilot-macos-intel`
- `build/coinpilot-macos-arm64` → Rename to `coinpilot-macos-arm64`

#### Linux Versions
- `build/coinpilot-linux-amd64` → Rename to `coinpilot-linux-x64`
- `build/coinpilot-linux-arm64` → Rename to `coinpilot-linux-arm64`

### 4. Publish

Click "Publish release" to complete the release.

## Version Naming Convention

Use Semantic Versioning:

- `v1.0.0` - Major version (breaking changes)
- `v1.1.0` - Minor version (new features)
- `v1.0.1` - Patch version (bug fixes)

## Release Notes Template

```markdown
## 🚀 New Features
- Added trade recording functionality
- Support for multi-exchange management

## 🐛 Bug Fixes
- Fixed CSV file reading issues
- Improved error handling

## 📦 Downloads

### Windows
- [Windows 64-bit](link) - Recommended for most Windows users
- [Windows 32-bit](link) - For older systems

### macOS
- [macOS Universal](link) - Supports both Intel and Apple Silicon
- [macOS Intel](link) - Intel Mac only
- [macOS Apple Silicon](link) - M1/M2 Mac only

### Linux
- [Linux 64-bit](link) - For most Linux distributions
- [Linux ARM64](link) - For ARM64 architecture

## 📋 Installation Instructions

After downloading the appropriate file for your platform:

1. **Windows**: Double-click to run or use in command line
2. **macOS**: May need to allow in "System Preferences > Security & Privacy"
3. **Linux**: Add execute permission `chmod +x coinpilot-linux-x64`

## 🔄 Updates

If you have a previous version installed, simply download the new version and replace the old file.
```

## Automated CI/CD (Future Enhancement)

Consider setting up GitHub Actions to automate the release process:

1. Automatically trigger on new tag push
2. Auto-build for all platforms
3. Auto-create Release and upload files
4. Auto-generate changelog

This will be implemented as the project matures.

---

[中文版本](RELEASE_ZH.md)