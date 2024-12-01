# Variables
APP_NAME = backend
BUILD_DIR = .
WEB_APP_DIR = ./web/
WEB_APP_BUILD_DIR = ./web/build
MAIN_SRC = cmd/api/api.go

# Default target
all: build

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR) $(WEB_APP_BUILD_DIR)
	@npm --prefix $(WEB_APP_DIR) install
	@npm --prefix $(WEB_APP_DIR) run build
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_SRC)
	@echo "Build completed."

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Clean the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed."

docker-image:
	docker build ./ -f ./build/Dockerfile \
		--tag nginx-admin-panel-backend \
		--build-arg GO_VER=1.23 \
		--build-arg VERSION=0.0.1 \
		--build-arg ALPINE_VER=3.20

docker-run:
	docker-compose build
	docker-compose up

.PHONY: all build run clean docker-image docker-run
