#!/bin/bash

# GitHub Release Script for CoinPilot
# This script helps create a GitHub release with pre-built binaries

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
REPO="Crows-Storm/coinpilot"
BUILD_DIR="build"

# Check if version is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Please provide a version tag${NC}"
    echo "Usage: $0 <version> [release-notes]"
    echo "Example: $0 v1.0.0 'Initial release with basic trading functionality'"
    exit 1
fi

VERSION=$1
RELEASE_NOTES=${2:-"Release $VERSION"}

echo -e "${GREEN}Creating GitHub Release for CoinPilot $VERSION${NC}"

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo -e "${RED}Error: GitHub CLI (gh) is not installed${NC}"
    echo "Please install it from: https://cli.github.com/"
    echo "Or use the manual method described below."
    exit 1
fi

# Check if user is authenticated
if ! gh auth status &> /dev/null; then
    echo -e "${YELLOW}You need to authenticate with GitHub CLI${NC}"
    echo "Run: gh auth login"
    exit 1
fi

# Check if build directory exists
if [ ! -d "$BUILD_DIR" ]; then
    echo -e "${YELLOW}Build directory not found. Building binaries...${NC}"
    make build-all
    if [ "$(uname)" = "Darwin" ]; then
        make build-macos-universal
    fi
fi

echo -e "${GREEN}Creating release $VERSION...${NC}"

# Create the release
gh release create "$VERSION" \
    --repo "$REPO" \
    --title "CoinPilot $VERSION" \
    --notes "$RELEASE_NOTES" \
    "$BUILD_DIR/coinpilot-windows-amd64.exe#CoinPilot for Windows (64-bit)" \
    "$BUILD_DIR/coinpilot-windows-386.exe#CoinPilot for Windows (32-bit)" \
    "$BUILD_DIR/coinpilot-macos-intel#CoinPilot for macOS (Intel)" \
    "$BUILD_DIR/coinpilot-macos-arm64#CoinPilot for macOS (Apple Silicon)" \
    "$BUILD_DIR/coinpilot-linux-amd64#CoinPilot for Linux (64-bit)" \
    "$BUILD_DIR/coinpilot-linux-arm64#CoinPilot for Linux (ARM64)"

# Add universal macOS binary if it exists
if [ -f "$BUILD_DIR/coinpilot-macos-universal" ]; then
    gh release upload "$VERSION" \
        --repo "$REPO" \
        "$BUILD_DIR/coinpilot-macos-universal#CoinPilot for macOS (Universal)"
fi

echo -e "${GREEN}✅ Release $VERSION created successfully!${NC}"
echo -e "${GREEN}🔗 View at: https://github.com/$REPO/releases/tag/$VERSION${NC}"