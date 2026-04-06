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

	platform, contestId, problemId, err := ExtractPlatformAndProblemName(problemContent.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Received problem: %s (platform: %s, contestId: %s, problemId: %s)", problemContent.Name, platform, contestId, problemId)

	w.WriteHeader(http.StatusOK)
}
