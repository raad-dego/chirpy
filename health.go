package main

import (
	"net/http"
)

// Handler for /healthz
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
