start:
	go run cmd/homestorage/main.go -config=./pkg/config/config.toml

build:
	go build -o ./bin/homestorage cmd/homestorage/main.go

start_web_prod:
	./bin/hfsweb -config=./pkg/config/config.toml

start_cli:
	go run cmd/hfscli/main.go

tests:
	go test -v ./...

templates:
	~/go/bin/templ generate -f web/templates/base.templ
	~/go/bin/templ generate -f web/templates/dashboard.templ