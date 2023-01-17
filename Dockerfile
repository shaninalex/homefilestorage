FROM golang:1.19-alpine AS build

WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 go build -o /src/application

FROM scratch
COPY --from=build /src/config.toml /src/config.toml
COPY --from=build /src/application /src/application
ENTRYPOINT ["/src/application", "-config=/src/config.toml"]
