package app

import (
	"auction/internal/db"
	"auction/internal/log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type App struct {
}

func (a App) Handler() http.HandlerFunc {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return router.ServeHTTP
}

func New(postgresDbProvider *gorm.DB) App {
	postgresDbUrl := os.Getenv("POSTGRES_DB_URL")
	if err := db.Migrate(postgresDbUrl); err != nil {
		log.ZeroLogger.Fatal().Err(err).Msg("Failed to run database migrations")
	}
	return App{}
}
