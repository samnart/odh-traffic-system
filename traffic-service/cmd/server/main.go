package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/samnart/odh-traffic-system/traffic-service/internal/handler"
	"github.com/samnart/odh-traffic-system/traffic-service/internal/middleware/logging"
	"github.com/samnart/odh-traffic-system/traffic-service/internal/middleware/recovery"
	"github.com/samnart/odh-traffic-system/traffic-service/internal/service"
	"github.com/samnart/odh-traffic-system/traffic-service/pkg/cache"
)

func main() {
	log.Println("Starting Traffic Service...")

	// Initialize Redis cache
	if err := cache.InitRedis(cache.DefaultConfig()); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer cache.Close()

	// Setup services and handlers
	trafficService := service.NewTrafficService()
	trafficHandler := handler.NewTrafficHandler(trafficService)

	// Set up Gin router with middlewares
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.Logger())
	router.Use(recovery.Recovery())

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Register routes
	traffic := router.Group("/traffic")
	{
		traffic.GET("/summary", trafficHandler.GetTrafficSummary)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr: 			":" + port,
		Handler: 		router,
		ReadTimeout:	15 * time.Second,
		WriteTimeout:	15 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Traffic Service listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down to server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}