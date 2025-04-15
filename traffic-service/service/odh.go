

package service

import (
	"encoding/json"
	"time"

	"github.com/samnart/odh-traffic-system/traffic-service/cache"
)

type TrafficData struct {
	TotalVehicles int       `json:"totalVehicles"`
	Timestamp     time.Time `json:"timestamp"`
}

func GetTrafficSummary() (*TrafficData, error) {
	cacheKey := "traffic:summary"
	if cached, err := cache.Get(cacheKey); err == nil && cached != "" {
		var data TrafficData
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			return &data, nil
		}
	}

	// Simulated fetch from Open Data Hub (replace with real call)
	data := &TrafficData{
		TotalVehicles: 1294,
		Timestamp:     time.Now(),
	}

	jsonData, _ := json.Marshal(data)
	_ = cache.Set(cacheKey, string(jsonData), 5*time.Minute)

	return data, nil
}
