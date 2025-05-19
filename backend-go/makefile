# Project Metadata
APP_NAME := koala-world-backend
PKG := ./...
MAIN := cmd/server/main.go
BIN_DIR := bin
COVERAGE_FILE := coverage.out

# Go Tools
GO := go
GOLINT := golangci-lint

# Build Flags
LD_FLAGS := -s -w

# Default Target
.PHONY: all
all: test build

##----------------------##
## Development Targets  ##
##----------------------##

.PHONY: run
run:
	$(GO) run $(MAIN)

.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	$(GO) build -ldflags "$(LD_FLAGS)" -o $(BIN_DIR)/$(APP_NAME) $(MAIN)

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) $(COVERAGE_FILE) coverage.html

.PHONY: fmt
fmt:
	$(GO) fmt $(PKG)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: lint
lint:
	@$(GOLINT) run

##----------------------##
## Testing Targets      ##
##----------------------##

.PHONY: test
test:
	$(GO) test -v $(PKG)

.PHONY: test-cover
test-cover:
	$(GO) test -v -cover -coverprofile=$(COVERAGE_FILE) $(PKG)

.PHONY: test-html
test-html: test-cover
	@$(GO) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Open coverage.html in browser to view coverage report"

.PHONY: test-short
test-short:
	$(GO) test -short $(PKG)

.PHONY: test-race
test-race:
	$(GO) test -race $(PKG)

##----------------------##
## Docker Targets       ##
##----------------------##

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(APP_NAME):latest

##----------------------##
## Help                 ##
##----------------------##

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Common targets:"
	@echo "  run            Run the app"
	@echo "  build          Compile the app"
	@echo "  clean          Remove binaries and temp files"
	@echo "  fmt            Format source code"
	@echo "  tidy           Clean go.mod/go.sum"
	@echo "  lint           Run linter (requires golangci-lint)"
	@echo "  test           Run all tests"
	@echo "  test-cover     Run tests with coverage"
	@echo "  test-html      Generate coverage HTML report"
	@echo "  test-race      Run tests with race detection"
	@echo "  docker-build   Build Docker image"
	@echo "  docker-run     Run Docker container"
