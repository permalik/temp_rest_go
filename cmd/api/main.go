package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/permalik/temp_rest_go/internal/data"
)

type application struct {
	config data.AppConfig
	logger *log.Logger
}

func main() {
	var cfg data.AppConfig

	// TODO: envvar from flag, then .env, then default
	// TODO: if containerized, then env
	flag.StringVar(&cfg.Env, "env", "development", "Environment(development|staging|production)")
	flag.IntVar(&cfg.Port, "port", 9000, "api server port")
	flag.Parse()
	cfg.Version = "0.1.0"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	var h http.Handler
	r := http.NewServeMux()
	r.HandleFunc("GET /v0/healthcheck", app.healthcheck)
	r.HandleFunc("POST /v0/item", app.create_item)
	r.HandleFunc("GET /v0/items", app.read_items)
	r.HandleFunc("GET /v0/item/{id}", app.read_item)
	r.HandleFunc("GET /", app.home)
	h = r

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           h,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	logger.Printf("\nestablish server connection\nenv: %s\naddr: %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
