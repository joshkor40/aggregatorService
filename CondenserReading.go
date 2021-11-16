package main

import (
	"math/rand"
	"time"
)

type CondenserReading struct {
	CondenserId   string        `json:"condenserId"`
	CurrentVolume VolumeReading `json:"currentVolume"`
	Humidity      float64       `json:"humidity"`
	LastEmptied   time.Time     `json:"lastEmptied"`
	Uptime        float64       `json:"uptime"`
	Location      LatLonReading `json:"location"`
}

// HumidityReading requests

func NewCondenserReading(condenserId string) *CondenserReading {
	liters := (rand.Float64() * 5) + 10 // generate random liter number
	return &CondenserReading{
		CondenserId:   condenserId,
		CurrentVolume: *NewVolumeReading(liters),
		Humidity:      NewHumidityReading().Humidity,
		LastEmptied:   NewLastEmptiedReading().LastEmptied,
		Uptime:        NewUptimeReading().Uptime,
		Location:      *NewLatLonReading(),
	}
}
