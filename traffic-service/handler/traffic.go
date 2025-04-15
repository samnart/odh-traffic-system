package handler

import (
	"encoding/json"
	"net/http"

	"github.com/samnart/odh-traffic-system/traffic-service/service"
)

func GetTrafficSummary(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetTrafficSummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}