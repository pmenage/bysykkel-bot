package bysykkel

import (
	"bytes"
	"fmt"
	"log"
	"math"
)

// GetNearestBikes gives the user the bikes nearest to his position
func GetNearestBikes(userLat float64, userLong float64, stations StationsConfig, availability AvailabilityConfig) string {

	var buffer bytes.Buffer

	for _, station := range stations.Stations {
		if math.Abs(userLat-station.Center.Latitude) < 0.01 && math.Abs(userLong-station.Center.Longitude) < 0.01 {
			log.Printf("User is at %v and %v", userLat, userLong)
			for _, nearStation := range availability.Stations {
				if nearStation.ID == station.ID {
					msgText := fmt.Sprintf(
						"Station %v has %v bikes and %v locks\n",
						station.Title,
						nearStation.Availability.Bikes,
						nearStation.Availability.Locks)
					buffer.WriteString(msgText)
				}
			}
		}
	}
	return buffer.String()
}
