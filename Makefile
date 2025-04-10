# Variables
BINARY_NAME = smithron
SRC_DIR = .
OUT_DIR = bin
GO_FILES = $(shell git ls-files '*.go')

# Default target
all: build

# Build the Go binary
build: $(OUT_DIR)/$(BINARY_NAME)

$(OUT_DIR)/$(BINARY_NAME): $(GO_FILES)
	@mkdir -p $(OUT_DIR)
	go build -o $(OUT_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run tests
test:
	go test ./...

# Run formatter
fmt:
	go fmt ./...

# Clean up build artifacts
clean:
	rm -rf $(OUT_DIR)

.PHONY: all build test clean