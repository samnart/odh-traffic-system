package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/v5/middleware"
	"github.com/rs/cors"

	"github.com/samnart/odh-traffic-system/traffic-service/cache"
	"github.com/samnart/odh-traffic-system/traffic-service/handler"
	"github.com/samnart/odh-traffic-system/traffic-service/middleware"
)


func main() {
	log.Println("Starting Traffic Service...")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(logging.LoggingMiddleware)
	r.Use(recovery.RecoveryMiddleware)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:		[]string{"*"},
		AllowedMethods:		[]string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:		[]string{"Accept", "Authorization", "Content-Type"},
		AllowedCredentials:	true,
	})
	r.Use(corsHandler.Handler)

	// Init Redis
	err := cache.InitRedis()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	// API routes
	r.Route("/traffic", func (r chi.Router) {
		r.Get("/summary", handler.GetTrafficSummary)		
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	srv := &http.Server{
		Addr: 			":" port,
		Handler: 		r,
		ReadTimeout:	15 * time.Second,
		WriteTimeout:	15 * time.Second,
	}

	log.Printf("Traffic Service listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}