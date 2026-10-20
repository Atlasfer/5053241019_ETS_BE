package main

import (
	"encoding/json"
	"net/http"
)

func tulisJSON(w http.ResponseWriter, kode int, data any) {
	w.Header().Set("Content-Type", "application/json") // set SEBELUM WriteHeader
	w.WriteHeader(kode)
	json.NewEncoder(w).Encode(data)
}

func tulisError(w http.ResponseWriter, kode int, pesan string) {
	tulisJSON(w, kode, map[string]string{"error": pesan})
}

type apiError struct {
	Kode  int
	Pesan string
}

func tulisApiError(w http.ResponseWriter, e *apiError) {
	tulisError(w, e.Kode, e.Pesan)
}