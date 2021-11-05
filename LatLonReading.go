package main

import (
	"math/rand"
)

type LatLonReading struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func NewLatLonReading() *LatLonReading {
	northBoundary := 42.0
	southBoundary := 32.5
	eastBoundary := -114.13333333333334
	westBoundary := -124.40944444444445

	randLon := westBoundary + (rand.Float64() * (eastBoundary - westBoundary))
	randLat := southBoundary + (rand.Float64() * (northBoundary - southBoundary))

	return &LatLonReading{
		Latitude:  randLat,
		Longitude: randLon,
	}
}
