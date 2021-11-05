package main

import (
	"math/rand"
	"time"
)

type CondenserReading struct {
	CurrentVolume VolumeReading `json:"current-volume"`
	Humidity      float64       `json:"humidity"`
	LastEmptied   time.Time     `json:"last-emptied"`
	Uptime        float64       `json:"uptime"`
	Location      LatLonReading `json:"location"`
}

// HumidityReading requests

func NewCondenserReading() *CondenserReading {
	liters := (rand.Float64() * 5) + 10 // generate random liter number
	return &CondenserReading{
		CurrentVolume: *NewVolumeReading(liters),
		Humidity:      NewHumidityReading().Humidity,
		LastEmptied:   NewLastEmptiedReading().LastEmptied,
		Uptime:        NewUptimeReading().Uptime,
		Location:      *NewLatLonReading(),
	}
}
