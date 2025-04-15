package model

type TrafficEntry struct {
	ID				string	`json:"id"`
	Location		string	`json:"location"`
	DateObserved	string	`json:"dateObserved"`
	VehicleCount	int		`json:"vehicleCount"`
	AverageSpeed	float64	`json:"averageSpeed"`
}