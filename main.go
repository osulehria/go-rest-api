package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/osulehria/go-rest-api/db"
	"github.com/osulehria/go-rest-api/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	healthcheck := routes.NewHealthCheckResource(db.NewRedisDB(redisClient))

	r.Mount("/healthz", healthcheck.Routes())

	http.ListenAndServe(":3333", r)
}
