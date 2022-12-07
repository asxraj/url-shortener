package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/asxraj/url-shortener/internal/database"
	"github.com/asxraj/url-shortener/internal/jsonlog"
	"github.com/asxraj/url-shortener/internal/mailer"
	"github.com/asxraj/url-shortener/internal/models"
	"github.com/joho/godotenv"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
	jwt  struct {
		secret string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models models.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server listen to port")
	flag.StringVar(&cfg.env, "env", "development", "development|staging|production")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 2525, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "9629c412b8b7b3", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "3dfb5523361d5c", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "SHORTURL <no-reply@shorturl.asxraj.com>", "SMTP sender")

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
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
