package logging

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs request information and timing
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method

		log.Printf("[%d] %s %s - %v", statusCode, &method, &path, latency)
	}
}