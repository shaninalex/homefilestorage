start:
	go run cmd/homestorage/main.go -config=./pkg/config/config.toml

build:
	env GOOS=linux GOARCH=arm CGO_ENABLED=1 go build -tags=sqlite_omit_load_extension -o ./bin/homestorage cmd/homestorage/main.go

start_web_prod:
	./bin/homestorage -config=./pkg/config/config.toml

tests:
	go test -v ./...

templates:
	~/go/bin/templ generate -f web/templates/base.templ
	~/go/bin/templ generate -f web/templates/dashboard.templ