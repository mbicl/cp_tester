package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded successfully: %+v", config)

	server := http.NewServeMux()

	server.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server.HandleFunc("/", handleCompetitiveCompanion)

	log.Printf("Starting server on :10043")
	log.Fatal(http.ListenAndServe(":10043", loggingMiddleware(server)))
}

// handleCompetitiveCompanion handles the incoming requests from Competitive Companion and creates necessary files for writing the solution and testing the problem locally.
func handleCompetitiveCompanion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var problemContent ProblemContent
	err := json.NewDecoder(r.Body).Decode(&problemContent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = LoadConfig()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Middlewares

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		log.Printf("%d %s %s", rec.status, r.Method, r.URL.Path)
	})
}
