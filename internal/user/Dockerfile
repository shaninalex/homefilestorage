FROM golang:1.20
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o userservice .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /build/userservice ./
CMD ["./userservice"]