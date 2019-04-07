package models

import (
	"encoding/json"
	"time"
)

// Location is a latiture longitude pair
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// CurrentLocation reprecents the current location for a player in a game
type CurrentLocation struct {
	Timestamp time.Time `json:"timestamp"`
	Game      int       `json:"game"`
	Player    string    `json:"player"`
	Location  Location  `json:"location"`
}

// NewCurrentLocation unmarshals a CurrentLocation from Json
func NewCurrentLocation(data []byte) (CurrentLocation, error) {
	var c CurrentLocation
	err := json.Unmarshal(data, &c)
	return c, err
}
