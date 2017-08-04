package bysykkel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Stations contains all the stations config according to ID
type Stations map[int]StationConfig

// Config has the configuration of the stations
type Config struct {
	Stations []StationConfig `json:"stations"`
}

// StationConfig contains the configuration of a station
type StationConfig struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Subtitle      string        `json:"subtitle"`
	NumberOfLocks int           `json:"number_of_locks"`
	Center        coordinates   `json:"center"`
	Bounds        []coordinates `json:"bounds"`
}

type coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetStations gets the stations near you
func GetStations(key string) Stations {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://oslobysykkel.no/api/v1/stations", nil)
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

	var c Config
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	stations := make(Stations)
	for _, config := range c.Stations {
		stations[config.ID] = config
	}

	return stations

}
