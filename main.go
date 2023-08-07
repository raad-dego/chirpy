package main

import (
	"log"
	"net/http"
	"github.com/raad-dego/chirpy/api"
	"github.com/go-chi/chi"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	apiCfg := &ApiConfig{
		fileserverHits: 0,
	}

	// ServeMux
	// mux := http.NewServeMux()
	// mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	// mux.HandleFunc("/healthz", healthHandler)
	// mux.HandleFunc("/metrics", apiCfg.metricsHandler)
	// corsMux := MiddlewareCors(mux)
	// srv := &http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: corsMux,
	// }
	// Chi router
	r := chi.NewRouter()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	r.Handle("/app/*", fsHandler)
	r.Handle("/app", fsHandler)
	r.Get("/healthz", HealthHandler)
	r.Get("/metrics", apiCfg.metricsHandler)

	corsR := MiddlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsR,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
