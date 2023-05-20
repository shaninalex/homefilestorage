start:
	docker compose -f docker-compose.tmp.yml up -d --build

down:
	docker compose -f docker-compose.tmp.yml down -v

restart: down start