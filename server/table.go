package server

import (
	"encoding/json"
	"net/http"
)

func handleTable(w http.ResponseWriter, r *http.Request) {
	table := league.GetTable()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(table)
}
