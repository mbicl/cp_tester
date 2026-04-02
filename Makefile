.PHONY: build clean run build-linux build-darwin build-windows build-all install

GO=go
APP_NAME=cp
VERSION=$(shell git describe --tags --always)
BUILD_DIR=build
MAIN=cmd/cli/main.go

build:
	@echo "Building $(APP_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN)
	@echo "Build completed: $(BUILD_DIR)/$(APP_NAME)"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed."

run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

build-linux:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_linux_amd64 -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN)

build-darwin:
	@echo "Building for macOS (arm64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_darwin_arm64 -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN)

build-windows:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_windows_amd64.exe -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN)

build-all: build-linux build-darwin build-windows
	@echo "All builds completed."

install: build
	@echo "Installing $(APP_NAME) to /usr/local/bin..."
	@install -m 755 $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "Installation completed."