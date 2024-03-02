FROM golang:1.22.0 AS builder
LABEL authors="Georgy Solonitsky"
LABEL description="Simple http server, wich returns client headers, works on 80 port"
WORKDIR /build
COPY ./main/main.go main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /headersserver main.go
FROM alpine:3.19.1 AS publish
EXPOSE 80
WORKDIR /app
COPY --from=builder /headersserver headersserver
RUN apk add --no-cache bash
ENTRYPOINT ["/app/headersserver"]