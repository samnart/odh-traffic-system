package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/samnart/odh-traffic-system/traffic-service/cache"
	"github.com/samnart/odh-traffic-system/traffic-service/model"
)

const (
	odhURL 		= "https://mobility.api.opendatahub.bz.it/v2/flat/TrafficFlowObserved"
	cacheKey 	= "traffic_latest"
	cacheTTL	= 30 * time.Second
)

func FetchTrafficData() ([]model.TrafficEntry, error) {
	// Check cache
	if cached, err := cache.Get(cacheKey); err == nil {
		log.Println("Serving traffic data from cache")
		var data []model.TrafficEntry
		if err := json.Unmarshal([]byte(cached), &data); err != nil {
			return data, nil
		}
	}

	log.Println("Fetching traffic data from Open Data Hub...")
	resp, err := http.Get(odhURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed %v", err)
	}

	var raw struct {
		Data []model.TrafficEntry `json:"data"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %v", err)
	}

	// save to cache
	jsonData, _ := json.Marshal(raw.Data)
	_ = cache.Set(cacheKey, string(jsonData), cacheTTL)

	log.Println("Fetched and cached fresh traffic data")
	return raw.Data, nil
}