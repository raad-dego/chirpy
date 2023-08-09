package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/raad-dego/chirpy/chirpyApi"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	apiCfg := &chirpyApi.ApiConfig{
		FileserverHits: 0,
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

	r := chi.NewRouter()
	fsHandler := apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	r.Handle("/app/*", fsHandler)
	r.Handle("/app", fsHandler)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", chirpyApi.HealthHandler)
	r.Mount("/api", apiRouter)
	adminRouter := chi.NewRouter()
	adminRouter.Get("/metrics", http.HandlerFunc(apiCfg.AdminMetrics().ServeHTTP))
	r.Mount("/admin", adminRouter)

	

	corsR := MiddlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsR,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
