package bysykkel

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/kellydunn/golang-geo"
)

type result struct {
	Title    string
	Bikes    int
	Locks    int
	Distance int
	Closed   bool
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

func getNearest(userLat float64, userLong float64, stations Stations, availability AvailabilityConfig, status Status) results {

	var r results
	userPoint := geo.NewPoint(userLat, userLong)

	for _, nearStation := range availability.Stations {

		id := nearStation.ID
		stationPoint := geo.NewPoint(stations[id].Center.Latitude, stations[id].Center.Longitude)
		distance := int(userPoint.GreatCircleDistance(stationPoint) * 1000)
		var station result
		var closed bool

		for _, s := range status.Status.StationsClosed {
			if s == id {
				closed = true
			} else {
				closed = false
			}
		}

		station = result{
			Title:    stations[id].Title,
			Bikes:    nearStation.Availability.Bikes,
			Locks:    nearStation.Availability.Locks,
			Distance: distance,
			Closed:   closed,
		}
		r = append(r, station)

	}

	return r

}

func getMessage(r results, i int, t string) string {

	if r[i].Closed {
		return fmt.Sprintf(
			"Station %v, %v meters away, is closed\n", r[i].Title, r[i].Distance)
	} else if r[i].Bikes == 1 {
		return fmt.Sprintf(
			"Station %v, %v meters away, has %v %v\n", r[i].Title, r[i].Distance, r[i].Bikes, t)
	} else {
		return fmt.Sprintf(
			"Station %v, %v meters away, has %v %vs\n", r[i].Title, r[i].Distance, r[i].Bikes, t)
	}

}

// GetNearestBikes gives the user the bikes nearest to his position
func GetNearestBikes(userLat float64, userLong float64, stations Stations, availability AvailabilityConfig, status Status) string {

	var buffer bytes.Buffer
	r := getNearest(userLat, userLong, stations, availability, status)

	sort.Sort(r)
	for i := 0; i < 5; i++ {
		msgText := getMessage(r, i, "bike")
		buffer.WriteString(msgText)
	}

	return buffer.String()

}

// GetNearestLocks gives the user the locks nearest to his position
func GetNearestLocks(userLat float64, userLong float64, stations Stations, availability AvailabilityConfig, status Status) string {

	var buffer bytes.Buffer
	r := getNearest(userLat, userLong, stations, availability, status)

	sort.Sort(r)
	for i := 0; i < 5; i++ {
		msgText := getMessage(r, i, "lock")
		buffer.WriteString(msgText)
	}

	return buffer.String()

}
