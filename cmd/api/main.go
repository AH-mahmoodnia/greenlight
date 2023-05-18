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

// Application version number.
const version = "1.0.0"

// Holds configuration settings:
// port is the port numebr that server will listen on.
// env is the current operating environment (development, staging, production, ...)
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

// Holds dependencies for HTTP handlers, helpers, middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare an instance of config struct.
	var cfg config

	//Read the value of port and env from command-line and put it into
	//our config struct.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "Postgresql DSN")
	flag.Parse()

	// simple logger which write into stdout and prefix with local date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection pool established.")

	// Declare an instance of application struct and fill the config and logger fields of it.
	app := &application{
		config: cfg,
		logger: logger,
	}
	//Decalre a server and fill the fileds with above variables.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server and log it on stdout.
	logger.Printf("starting %s server on %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stablish a connection passing the above ctx as a parameter
	// If the connection couldn't be estblished successfully within the deadline above
	// this will return error
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
