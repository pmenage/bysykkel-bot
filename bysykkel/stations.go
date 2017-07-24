package bysykkel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// StationsConfig has the configuration of the stations
type StationsConfig struct {
	Stations []stationConfig `json:"stations"`
}

type stationConfig struct {
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
func GetStations(key string) StationsConfig {

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

	var c StationsConfig
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	return c

}
