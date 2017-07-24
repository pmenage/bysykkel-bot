package bysykkel

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/kellydunn/golang-geo"
)

type result struct {
	Title    string
	Bikes    int
	Locks    int
	Distance int
}

type results []result

func (slice results) Len() int {
	return len(slice)
}

func (slice results) Less(i, j int) bool {
	return slice[i].Distance < slice[j].Distance
}

func (slice results) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// GetNearestBikes gives the user the bikes nearest to his position
func GetNearestBikes(userLat float64, userLong float64, stations StationsConfig, availability AvailabilityConfig) string {

	var buffer bytes.Buffer
	var r results
	userPoint := geo.NewPoint(userLat, userLong)

	for _, station := range stations.Stations {
		if math.Abs(userLat-station.Center.Latitude) < 0.005 && math.Abs(userLong-station.Center.Longitude) < 0.005 {
			log.Printf("User is at %v and %v", userLat, userLong)
			for _, nearStation := range availability.Stations {
				if nearStation.ID == station.ID {

					stationPoint := geo.NewPoint(station.Center.Latitude, station.Center.Longitude)

					distance := int(userPoint.GreatCircleDistance(stationPoint) * 1000)
					fmt.Printf("\n\n Distance is: %v\n\n\n", distance)

					station := result{
						Title:    station.Title,
						Bikes:    nearStation.Availability.Bikes,
						Locks:    nearStation.Availability.Locks,
						Distance: distance,
					}

					r = append(r, station)

				}
			}
		}
	}

	sort.Sort(r)
	for i := 0; i < 5; i++ {
		msgText := fmt.Sprintf(
			"Station %v, at distance %v, has %v bikes and %v locks\n",
			r[i].Title,
			r[i].Distance,
			r[i].Bikes,
			r[i].Locks)
		buffer.WriteString(msgText)
	}

	if buffer.String() == "" {
		buffer.WriteString("There are no stations near you. Please try again later.")
	}
	return buffer.String()
}
