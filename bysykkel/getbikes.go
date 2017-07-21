package bysykkel

import "log"

// GetNearBikes gives the user the bikes nearest to his position
func GetNearBikes(userLat string, userLong string, stations StationsConfig, availability AvailabilityConfig) {

	log.Printf("User is at %v and %v", userLat, userLong)

}
