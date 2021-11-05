package main

import (
	"math/rand"
)

type HumidityReading struct {
	Humidity float64 `json:"humidity"`
}

// HumidityReading requests

func NewHumidityReading() *HumidityReading {
	return &HumidityReading{
		Humidity: (rand.Float64() * 50) + 35,
	}
}