package main

import "net/http"

func handleReadinessProbe(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status": "ok"})
}
