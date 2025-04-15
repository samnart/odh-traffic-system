package handler

import (
	"net/http"
	"traffic-service/service"

	"github.com/gin-gonic/gin"
)

func GetLatestTraffic(c *gin.Context) {
	data, err := service.FetchTrafficData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch traffic data"})
		return
	}

	c.JSON(http.StatusOK, data)
}