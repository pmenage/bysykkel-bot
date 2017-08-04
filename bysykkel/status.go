package bysykkel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Status has the status for all stations
type Status struct {
	Status StationsStatus `json:"status"`
}

// StationsStatus has status
type StationsStatus struct {
	AllStationsClosed bool  `json:"all_stations_closed"`
	StationsClosed    []int `json:"stations_closed"`
}

// GetStatus gets the status of the stations
func GetStatus(key string) Status {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://oslobysykkel.no/api/v1/status", nil)
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

	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		panic(err)
	}

	return status

}
