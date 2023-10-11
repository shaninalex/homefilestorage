start_api:
	go run cmd/hfsapp/main.go -config=./pkg/config/config.toml

start_web:
	go run cmd/hfsweb/main.go -config=./pkg/config/config.toml

build_web:
	go build -o ./bin/hfsweb cmd/hfsweb/main.go

start_web_prod:
	./bin/hfsweb -config=./pkg/config/config.toml

start_cli:
	go run cmd/hfscli/main.go

tests:
	go test -v ./...
