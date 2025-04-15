package main

import (
	"log"
	"traffic-service/cache"
	"traffic-service/handler"

	"github.com/gin-gonic/gin"
	"github.com/samnart/odh-traffic-system/traffic-service/cache"
	"github.com/samnart/odh-traffic-system/traffic-service/handler"
)


func main() {
	log.Println("Starting Traffic Service...")

	cache.InitRedis()

	r := gin.Default()
	r.GET("/traffic/latest", handler.GetLatestTraffic)

	log.Println("Server running on port 8081")
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Server failed: %v", err)
	}
}