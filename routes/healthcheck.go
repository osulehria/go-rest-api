package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/osulehria/go-rest-api/db"
)

type HealthCheckResource interface {
	Routes() chi.Router
	Get(http.ResponseWriter, *http.Request)
}

type healthCheckConfig struct {
	db db.RedisDB
}

func NewHealthCheckResource(db db.RedisDB) HealthCheckResource {
	return &healthCheckConfig{db: db}
}

func (hr *healthCheckConfig) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", hr.Get)

	return r
}

func (hr *healthCheckConfig) Get(w http.ResponseWriter, r *http.Request) {
	var status string

	if err := hr.db.Up(context.Background()); err != nil {
		status = "down"
	} else {
		status = "up"
	}

	w.Write([]byte(status))
}
