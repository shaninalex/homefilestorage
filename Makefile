start:
	docker compose -f ./build/docker-compose.yml up -d --build

down:
	docker compose -f ./build/docker-compose.yml down -v

restart: down start
