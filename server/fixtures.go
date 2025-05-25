package server

import (
	"encoding/json"
	"net/http"
)

func handleFixtures(w http.ResponseWriter, r *http.Request) {
	fixtures := league.GetFixtures()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fixtures)
}
