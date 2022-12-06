package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/asxraj/url-shortener/internal/database"
	"github.com/asxraj/url-shortener/internal/jsonlog"
	"github.com/asxraj/url-shortener/internal/models"
	"github.com/joho/godotenv"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
	smtp struct {
	}
	jwt struct {
		secret string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models models.Models
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

	cfg.dsn = os.Getenv("SHORTURL_DB_DSN")
	cfg.jwt.secret = os.Getenv("JWT_SECRET")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := database.OpenDB(cfg.dsn)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		logger: logger,
		models: models.New(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
