package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/samnart/odh-traffic-system/traffic-service/cache"
	"github.com/samnart/odh-traffic-system/traffic-service/model"
)

// TrafficService handles traffic data operations
type TrafficService struct {
	// Can be expanded with dependencies like database connections & external APIs
}

// NewTrafficService creates a new traffic service instance
func NewTrafficService() *TrafficService {
	return &TrafficService{}
}

// GetTrafficSummary retrieves traffic summary data, using cache when available
func (s *TrafficService) GetTrafficSummary() (*model.TrafficData, error) {
	cacheKey := "traffic:summary"

	// Try to get from cache first
	if cached, err := cache.Get(cacheKey); err == nil && cached != "" {
		var data model.TrafficData
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			return &data, nil
		}
	}

	// Cache miss or unmarshaling error - fetch from source
	data, err := s.fetchTrafficDataFromSource()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch traffic data: %v", &err)
	}

	// Store in cache for future requests
	jsonData, err := json.Marshal(data)
	if err == nil {
		_ = cache.Set(cacheKey, string(jsonData), 5 * time.Minute)
	}

	return data, nil
}

// fetchTrafficDataFromSource simulates getting data from an external source
// In a real application, this would make API calls to Open Data Hub
func (s *TrafficService) fetchTrafficDataFromSource() (*model.TrafficData, error) {
	// Simulated fetch from Open Data Hub
	data := &model.TrafficData{
		TotalVehicles: 	1294,
		Timestamp: 		time.Now(),
		Location: 		"Main Highway",
		AverageSpeed: 	65.4,
	}

	// Simulate a small delay to mimic network latency
	time.Sleep(200 * time.Millisecond)

	return data, nil
}