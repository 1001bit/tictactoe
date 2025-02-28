# Variables
DOCKER_COMPOSE = docker compose
TSCOMPILER = python3 typescript/tscompiler.py
TS_PATH = typescript

# Build and start
all: start

# Build the Docker containers
start:
	@echo "\nStarting Docker containers..."
	$(DOCKER_COMPOSE) up --build -d

# Compile typescript
tscompile:
	@echo "\nCompiling typescript..."
	$(TSCOMPILER) $(TS_PATH)

# Stop the Docker containers
down:
	@echo "\nStopping Docker containers..."
	$(DOCKER_COMPOSE) down

# Clean up Docker resources
clean:
	@echo "\nCleaning up Docker resources..."
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans