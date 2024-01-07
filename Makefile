db/migrate: ## Migrate database structure
	@scripts/migrate.sh up

db/up: ## Apply all the migration to the latest version to the local database
	@make db/migrate

db/drop: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh drop -f

docker/up: ## Run docker compose
	docker-compose -f ./docker/docker-compose.yaml up -d

docker/down: ## Stop docker compose
	docker-compose -f ./docker/docker-compose.yaml down