package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func HandleCompetitiveCompanion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON payload
	var problemContent ProblemContent
	err := json.NewDecoder(r.Body).Decode(&problemContent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Load configuration
	config, err := LoadConfig()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract platform, contestId, and problemId from the URL
	platform, contestId, problemId, err := ExtractPlatformAndProblemName(problemContent.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Received problem: %s (platform: %s, contestId: %s, problemId: %s)", problemContent.Name, platform, contestId, problemId)

	config.CPPath += "/" + platform
	if contestId != "" {
		config.CPPath += "/" + contestId
	}

	// Create necessary folders for the problem
	if err := createFolders(config.CPPath); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save the tests of the problem
	if err := saveTests(problemContent.Tests, config.CPPath); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create the solution file if it doesn't exist
	if err := createSolutionFile(config, &problemContent, problemId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
