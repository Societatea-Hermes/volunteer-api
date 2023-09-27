package main

import (
	"fmt"
	"hermes-api/db"
	"hermes-api/helpers"
	"hermes-api/router"
	"hermes-api/services"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

var port = helpers.GetEnv("PORT", "8080")

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

	dsn := os.Getenv("DSN") // data source name
	log.Printf("DSN %s\n", dsn)
	dbConn, err := db.NewAdapter(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	defer dbConn.CloseDbConnection()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
