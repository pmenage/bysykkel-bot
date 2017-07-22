package bysykkel

import (
	"log"
	"math"
)

// GetNearBikes gives the user the bikes nearest to his position
func GetNearBikes(userLat float64, userLong float64, stations StationsConfig, availability AvailabilityConfig) {

	for _, station := range stations.Stations {
		if math.Abs(userLat-station.Center.Latitude) < 0.01 && math.Abs(userLong-station.Center.Longitude) < 0.01 {
			log.Printf("User is at %v and %v", userLat, userLong)
		}
	}
}
