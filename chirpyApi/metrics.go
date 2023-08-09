package chirpyApi

import (
	"fmt"
	"net/http"
)

type ApiConfig struct {
	FileserverHits int
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits++
		next.ServeHTTP(w, r)
	})
}

// func (cfg *ApiConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.FileserverHits)))
// }

func (cfg *ApiConfig) AdminMetrics() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
		</html>
		`

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, html, cfg.FileserverHits)
	})
}
