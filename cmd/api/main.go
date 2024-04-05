package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	env  string
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// TODO: envvar from flag, then .env, then default
	// TODO: if containerized, then env
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.IntVar(&cfg.port, "port", 9000, "api server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("TEMP_REST_GO_DSN"), "PostgreSQL DSN")
	// flag.StringVar(&cfg.db.dsn, "db-dsn", "host=localhost port=5432 user=au4 dbname=db sslmode=disable", "PostgreSQL DSN")
	// TODO: update the above the utilize one or all of the below strategies as well if prudent
	// flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://db:password@localhost/db", "PostgreSQL DSN")
	// connStr := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	// Set each value dynamically w/ Sprintf
	// connStr = fmt.Sprintf(connStr, host, port, user, dbname, sslmode)
	flag.Parse()
	// version := "0.1.0"

	db, err := db_open(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Printf("established db connection pool")

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
		Addr:              fmt.Sprintf(":%d", cfg.port),
		Handler:           h,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	logger.Printf("\nestablish server connection\nenv: %s\naddr: %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func db_open(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
