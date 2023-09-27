package main

import (
	"fmt"
	"hermes-api/db"
	"hermes-api/helpers"
	"hermes-api/router"
	"hermes-api/services"
	"log"
	"net/http"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	fmt.Println("API listening on port", app.Config.Port)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {

	db_port := helpers.GetEnv("DB_PORT", "5432")
	db_host := helpers.GetEnv("DB_HOST", "localhost")
	db_user := helpers.GetEnv("DB_USER", "postgres")
	db_pass := helpers.GetEnv("DB_PASS", "postgres")
	db_name := helpers.GetEnv("DB_NAME", "postgres")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_pass, db_name)

	dbConn, err := db.NewAdapter(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	defer dbConn.CloseDbConnection()

	var cfg Config
	cfg.Port = helpers.GetEnv("PORT", "8080")

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
