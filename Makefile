start_api:
	go run cmd/hfsapp/main.go -config=./pkg/config/config.toml	

start_cli:
	go run cmd/hfscli/main.go

tests:
	go test -v ./...
