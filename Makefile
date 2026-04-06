.PHONY: build clean run build-linux build-darwin build-windows build-all install uninstall
default: run

GO=go
APP_NAME=cpt
SERVER_NAME=cpt_server
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BUILD_DIR=build
CLI_APP=./src/cmd/cli
SERVER_APP=./src/cmd/server
LDFLAGS=-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

build:
	@echo "Building $(APP_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="$(LDFLAGS)" $(CLI_APP)
	@echo "Building $(SERVER_NAME) version $(VERSION)..."
	@$(GO) build -o $(BUILD_DIR)/$(SERVER_NAME) -ldflags="$(LDFLAGS)" $(SERVER_APP)
	@echo "Build completed.  $(BUILD_DIR)/$(APP_NAME) and $(BUILD_DIR)/$(SERVER_NAME) created."

build-app:
	@echo "Building $(APP_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="$(LDFLAGS)" $(CLI_APP)
	@echo "Build completed.  $(BUILD_DIR)/$(APP_NAME) created."

build-server:
	@echo "Building $(SERVER_NAME) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(SERVER_NAME) -ldflags="$(LDFLAGS)" $(SERVER_APP)
	@echo "Build completed.  $(BUILD_DIR)/$(SERVER_NAME) created."

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed."

run: build
	@echo "Running $(SERVER_NAME) in the background..."
	@./$(BUILD_DIR)/$(SERVER_NAME) &
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

app: build-app
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)
	
server: build-server
	@echo "Running $(SERVER_NAME)..."
	@./$(BUILD_DIR)/$(SERVER_NAME)

build-linux:
	@echo "Building for Linux (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_linux_amd64 -ldflags="$(LDFLAGS)" $(CLI_APP)
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(SERVER_NAME)_linux_amd64 -ldflags="$(LDFLAGS)" $(SERVER_APP)

build-darwin:
	@echo "Building for macOS (arm64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_darwin_arm64 -ldflags="$(LDFLAGS)" $(CLI_APP)
	@GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(SERVER_NAME)_darwin_arm64 -ldflags="$(LDFLAGS)" $(SERVER_APP)

build-windows:
	@echo "Building for Windows (amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)_windows_amd64.exe -ldflags="$(LDFLAGS)" $(CLI_APP)
	@GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(SERVER_NAME)_windows_amd64.exe -ldflags="$(LDFLAGS)" $(SERVER_APP)

build-all: build-linux build-darwin build-windows
	@echo "All builds completed."

install: build
	@echo "Installing $(APP_NAME) to /usr/local/bin..."
	@install -m 755 $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "Installation completed."

uninstall:
	@echo "Uninstalling $(APP_NAME) from /usr/local/bin..."
	@rm -f /usr/local/bin/$(APP_NAME)
	@echo "Uninstallation completed."