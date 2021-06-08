# Go REST API Playground

A playground to check out how to write REST APIs in Golang, with testing. Uses [Redis](https://github.com/go-redis/redis) as the data store and [Chi](https://github.com/go-chi/chi) as the HTTP router.

## Prerequisites

- Go 1.16+
- Docker

## Start

```bash
docker compose up -d # starts redis
go run main.go
```

## Test

```bash
go test -cover -v ./...
```
