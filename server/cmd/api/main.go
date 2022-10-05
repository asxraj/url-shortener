package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server listen to port")
	flag.StringVar(&cfg.env, "env", "development", "development|staging|production")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Println("Starting up server on port", srv.Addr)
	log.Fatal(srv.ListenAndServe())
	log.Fatal(err)

}
