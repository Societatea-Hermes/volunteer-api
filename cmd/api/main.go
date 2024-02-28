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

var port string

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func configureApp() (int, string, string, string, string) {
	p, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = "8080"
	}
	port = p

	db_port_s, ok := os.LookupEnv("DB_PORT")
	if !ok {
		db_port_s = "4530"
	}
	db_port, _ := strconv.Atoi(db_port_s)

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		host = "localhost"
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		user = "postgres"
	}

	pass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		pass = "postgres"
	}

	db_name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		db_name = "volunteers"
	}

	return db_port, host, user, pass, db_name
}

func main() {
	var cfg Config
	cfg.Port = port

	db_port, host, pass, user, db_name := configureApp()

	config := db.DatabaseConfig{
		Host:     host,
		Port:     db_port,
		Username: user,
		Password: pass,
		DBName:   db_name,
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
