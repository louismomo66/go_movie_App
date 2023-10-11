package main

import (
	"backend/internal/models/repository"
	"backend/internal/models/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 9000

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	//set application config
	var app application
	//read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgress connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "string secret")
	flag.Parse()
	//connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	//start a web server

	log.Println("Starting application on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
