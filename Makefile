# Variables
BINARY_NAME = smithron
SRC_DIR = .
OUT_DIR = bin
GO_FILES = $(wildcard $(SRC_DIR)/*.go)

# Default target
all: build

# Build the Go binary
build: $(GO_FILES)
	@mkdir -p $(OUT_DIR)
	go build -o $(OUT_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	rm -rf $(OUT_DIR)

.PHONY: all build test clean