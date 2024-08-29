
.PHONY: up down migrate build_event
# Run SQL migration file
build_event: 
	docker-compose up --build event-service
migrate: 
	docker exec -i postgres_event_db psql -U postgres -d booking_db -f ./db/migrations/001_create_events_table.sql

# Bring up the PostgreSQL container
up:
	docker-compose up -d

# Bring down the PostgreSQL container
down:
	docker-compose down

# # Run SQL migration file
# migrate:
# 	docker exec -i langapp-db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) -f /docker-entrypoint-initdb.d/init.sql