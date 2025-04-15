package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samnart/odh-traffic-system/traffic-service/internal/service"
)

// TrafficHandler handles HTTP requests related to traffic data
type TrafficHandler struct {
	trafficService *service.TrafficService
}

// NewTrafficHandler creates a new traffic handler instance
func NewTrafficHandler(trafficService *service.TrafficService) *TrafficHandler {
	return &TrafficHandler{
		trafficService: trafficService,
	}
}

// GetTrafficSummary handles the request for traffic summary data
func (h *TrafficHandler) GetTrafficSummary(c *gin.Context) {
	data, err := h.trafficService.GetTrafficSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}