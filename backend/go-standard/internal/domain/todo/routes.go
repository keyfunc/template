package todo

import (
	"log/slog"
	"mall/internal/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoDeps struct {
	Cfg    *config.Config
	Logger *slog.Logger
	DB     *pgxpool.Pool
	Mux    *http.ServeMux
}

func NewRoutes(deps *TodoDeps) {
	repository := NewRepository(RepositoryDeps{DB: deps.DB})
	service := NewService(ServiceDeps{Repo: repository})
	handler := NewHandler(HandlerDeps{Logger: deps.Logger, Service: service})
	// todo crud
	deps.Mux.HandleFunc("GET /api/v1/todo", handler.List)
	deps.Mux.HandleFunc("POST /api/v1/todo", handler.Create)
	deps.Mux.HandleFunc("PUT /api/v1/todo/{id}", handler.Update)
	deps.Mux.HandleFunc("DELETE /api/v1/todo/{id}", handler.Del)
}
