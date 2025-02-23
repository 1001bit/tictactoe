# Variables
DOCKER_COMPOSE = docker compose

# Build and start
all: start

# Build the Docker containers
start:
	@echo "\nStarting Docker containers..."
	$(DOCKER_COMPOSE) up --build -d

# Stop the Docker containers
down:
	@echo "\nStopping Docker containers..."
	$(DOCKER_COMPOSE) down

# Clean up Docker resources
clean:
	@echo "\nCleaning up Docker resources..."
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans