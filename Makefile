start:
	docker compose -f docker-compose.yml up -d --build

down:
	docker compose -f docker-compose.yml down -v

restart: down start