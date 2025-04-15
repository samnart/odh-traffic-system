package model

import (
	"time"
)

// TrafficData represents traffic information
type TrafficData struct {
	TotalVehicles	int			`json:"totalVehicles"`
	Timestamp		time.Time	`json:"timestamp"`
	Location		string		`json:"location,omitempty"`
	AverageSpeed	float64		`json:"averageSpeed,omitempty"`
}