#!/bin/bash

# Detect the operating system
OS=$(uname -s)
ARCH=$(uname -m)

# Determine download URL and archive filename based on OS and ARCH
if [[ "$OS" == "Linux" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        GO_ARCHIVE="go1.20.8.linux-amd64.tar.gz"
        GO_URL="https://go.dev/dl/$GO_ARCHIVE"
    elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
        GO_ARCHIVE="go1.20.8.linux-arm64.tar.gz"
        GO_URL="https://go.dev/dl/$GO_ARCHIVE"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "Darwin" ]]; then
    if [[ "$ARCH" == "x86_64" ]]; then
        GO_ARCHIVE="go1.20.8.darwin-amd64.tar.gz"
        GO_URL="https://go.dev/dl/$GO_ARCHIVE"
    elif [[ "$ARCH" == "arm64" ]]; then
        GO_ARCHIVE="go1.20.8.darwin-arm64.tar.gz"
        GO_URL="https://go.dev/dl/$GO_ARCHIVE"
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
else
    echo "Unsupported OS: $OS"
    exit 1
fi


# Parse command-line arguments
FORCE_CLEAN=false
if [[ "$1" == "--clean" ]]; then
    FORCE_CLEAN=true
    echo "Force clean enabled. Existing files will be removed."
fi

# Check if Go is already installed on the system
if command -v go &> /dev/null; then
    echo "Go is already installed on the system: $(go version)"
    SYSTEM_GO=true
    SYSTEM_GO=false
else
    SYSTEM_GO=false
fi

# If --clean is used, remove existing files, even if Go is installed
if [[ "$FORCE_CLEAN" == true ]]; then
    echo "Cleaning up existing Go files..."
    rm -rf "$GO_ARCHIVE" go ./tools/trond/trond
fi

# If Go is not installed, download and extract it
if [[ "$SYSTEM_GO" == false ]]; then
    # Check if the Go archive is already downloaded
    if [[ -f "$GO_ARCHIVE" ]]; then
        echo "$GO_ARCHIVE already exists. Skipping download."
    else
        echo "Downloading Go from $GO_URL..."
        curl -LO "$GO_URL"
    fi

    # Extract Golang to current directory
    if [[ -d "go" ]]; then
        echo "Go is already extracted. Skipping extraction."
    else
        echo "Extracting Go..."
        tar -xzf "$GO_ARCHIVE"
    fi

    # Set up Go binary path for the current directory
    GO_BIN="./go/bin"
    export PATH="$GO_BIN:$PATH"
else
    GO_BIN="$(dirname $(command -v go))"
fi

# Verify Go binary
if ! command -v go &> /dev/null; then
    echo "Go setup failed."
    exit 1
fi

echo "Go setup successful. Version: $(go version)"

# Check for the main.go file in ./tools/trond/
MAIN_GO_PATH="./tools/trond/main.go"
if [[ ! -f "$MAIN_GO_PATH" ]]; then
    echo "Error: $MAIN_GO_PATH not found!"
    exit 1
fi

# Change to the tools/trond directory for building
echo "Changing to ./tools/trond/ directory..."
cd ./tools/trond || exit

# Build the Go program
echo "Building main.go..."
"$GO_BIN/go" build -o trond main.go

# Check if the build was successful
if [[ -f "./trond" ]]; then
    echo "Build successful! The output binary is ./tools/trond/trond"
else
    echo "Build failed."
    exit 1
fi

# Move the binary to the original directory
echo "Moving the binary to the original directory..."
mv trond ../../

# Return to the original directory
echo "Returning to the original directory..."
cd - >/dev/null

if [[ -f "./trond" ]]; then
    echo "Binary moved successfully! Run ./trond to execute."
else
    echo "Failed to move the binary."
    exit 1
fi

echo "Done!"