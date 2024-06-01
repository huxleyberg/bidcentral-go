package main

import (
	"net/http"
	"os"

	"auction/internal/app"
	"auction/internal/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	a := app.New(providePostgres())

	//this is for running api requests on the same container
	err := http.ListenAndServe(":80", a.Handler())
	if err != nil {
		log.ZeroLogger.Fatal().Err(err).Msg("HTTP server failed to start")
	}
}

func providePostgres() *gorm.DB {
	postgresDbUrl := os.Getenv("POSTGRES_DB_URL")
	db, err := gorm.Open(postgres.Open(postgresDbUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		}})

	if err != nil {
		log.ZeroLogger.Fatal().Err(err).Msg("postgres database connection error")
	}
	return db
}
