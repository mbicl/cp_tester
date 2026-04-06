package main

import (
	"log"
	"net/http"

	"github.com/mbicl/cp_tester/src/internal/server"
)

func main() {
	srv := http.NewServeMux()

	srv.HandleFunc("/health", server.HealthCheckHandler)
	srv.HandleFunc("/", server.HandleCompetitiveCompanion)

	log.Printf("Starting server on :10043")
	log.Fatal(http.ListenAndServe(":10043", server.LoggingMiddleware(srv)))
}
