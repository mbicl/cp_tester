package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("cp %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	server := http.NewServeMux()

	// POST / - handle incoming code submissions
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			log.Printf("Method not allowed: %s %s", r.Method, r.URL.Path)
			return
		}
		log.Printf("Received code submission from %s", r.RemoteAddr)
		var body []byte
		if r.Body != nil {
			body, _ = io.ReadAll(r.Body)
		}
		// pretty print the JSON code content
		var prettyJSON map[string]interface{}
		if err := json.Unmarshal(body, &prettyJSON); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		prettyBody, _ := json.MarshalIndent(prettyJSON, "", "  ")
		log.Printf("Code content:\n%s", string(prettyBody))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code received"))
	})

	log.Println("Starting server on :10043")
	http.ListenAndServe(":10043", server)
}
