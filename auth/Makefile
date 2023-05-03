AUTH_SERVICE_BINARY=auth

# build and run backend
up_build: b_auth
	docker compose down
	docker compose up -d --build

# build services
b_auth:
	env GOOS=linux CGO_ENABLED=0 go build -o bin/${AUTH_SERVICE_BINARY} ./

# manage infrastructure
up:
	docker compose up -d --build

down:
	docker compose down

down_volumes:
	docker compose down -v

update_database:
	docker exec account_service_db psql -U postgres -tc "CREATE TABLE "users" ("id" SERIAL UNIQUE, "email" varchar NOT NULL UNIQUE, "lang" varchar NOT NULL, "phone" TEXT NOT NULL UNIQUE, "active" boolean NOT NULL DEFAULT true, "rating" int NOT NULL DEFAULT 0, "password" TEXT, "updated_at" TIMESTAMP NOT NULL DEFAULT (now()), "created_at" TIMESTAMP NOT NULL DEFAULT (now()))"
