package server

import (
	"encoding/json"
	"net/http"
)

func handleResults(w http.ResponseWriter, r *http.Request) {
	results := league.GetResults()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
