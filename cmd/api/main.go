package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Application version number.
const version = "1.0.0"

// Holds configuration settings:
// port is the port numebr that server will listen on.
// env is the current operating environment (development, staging, production, ...)
type config struct {
	port int
	env  string
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
	flag.Parse()

	// simple logger which write into stdout and prefix with local date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of application struct and fill the config and logger fields of it.
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Declare a new router and add a /v1/healthcheck route which dispathes requests
	//to the healthcheckHandler method.
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//Decalre a server and fill the fileds with above variables.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server and log it on stdout.
	logger.Printf("starting %s server on %d", cfg.env, cfg.port)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

}
