package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wevnasc/hermes/email"
	"github.com/wevnasc/hermes/server"
)

var (
	ServerAddr = os.Getenv("HTTP_SERVER_ADDR")
)

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags|log.Lshortfile)

	h := email.NewHandlers(logger)
	mux := http.NewServeMux()

	h.SetupRoutes(mux)
	srv := server.New(mux, ServerAddr)

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("server failed to start: %v", err)
	}

	return nil
}
