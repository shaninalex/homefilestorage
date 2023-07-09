start:
	docker compose -f ./build/docker-compose.yml up -d --build

down:
	docker compose -f ./build/docker-compose.yml down 

restart_app:
	docker compose -f ./build/docker-compose.yml up -d --no-deps --build filemanager

restart: down start
