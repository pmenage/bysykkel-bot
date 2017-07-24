package bysykkel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// AvailabilityConfig gives the bikes and locks at a given stations
type AvailabilityConfig struct {
	Stations    []station `json:"stations"`
	UpdatedAt   time.Time `json:"updated_at"`
	RefreshRate float64   `json:"refresh_rate"`
}

type station struct {
	ID           int          `json:"id"`
	Availability availability `json:"availability"`
}

type availability struct {
	Bikes int `json:"bikes"`
	Locks int `json:"locks"`
}

// GetStationsAvailability gets the availabilities near you
func GetStationsAvailability(key string) AvailabilityConfig {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://oslobysykkel.no/api/v1/stations/availability", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Client-Identifier", key)
	resp, err := netClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var c AvailabilityConfig
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	for _, station := range c.Stations {
		log.Printf("Station number %v has %v bikes and %v locks\n", station.ID, station.Availability.Bikes, station.Availability.Locks)
	}

	return c

}
