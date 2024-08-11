# Makefile at the root of the project

.PHONY: build up down logs air

# Build all Docker images
build:
	docker-compose build

# Start all services using Docker Compose
up:
	docker-compose up -d

# Stop all running services
down:
	docker-compose down

# Show logs from all services
logs:
	docker-compose logs -f

# Start Air daemon for hot-reloading (requires Air to be installed)
air-log-ingestion:
	cd log_ingestion && air

air-log-processing:
	cd log_processing && air

air-data-storage:
	cd data_storage && air

air-analytics:
	cd analytics && air

air-monitoring:
	cd monitoring && air
