package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"volunteer-api/db"
	"volunteer-api/models"
	"volunteer-api/router"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models models.Models
}

var port = os.Getenv("APP_PORT")

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	var cfg Config
	cfg.Port = port
	db_port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	config := db.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     db_port,
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dbAdapter, err := db.NewAdapter(config)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
	defer db.CloseDbConnection(dbAdapter)

	app := &Application{
		Config: cfg,
		Models: models.New(dbAdapter),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
