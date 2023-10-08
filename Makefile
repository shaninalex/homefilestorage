start_api:
	go run cmd/hfsapp/main.go

start_cli:
	go run cmd/hfscli/main.go

tests:
	go test -v ./...
