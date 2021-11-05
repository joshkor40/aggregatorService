package main

import (
	"math/rand"
	"time"
)

type LastEmptiedReading struct {
	LastEmptied time.Time `json:"last-emptied"`
}

func randate() time.Time {
	days := 1 + rand.Int63n(30)

	min := time.Now().AddDate(0, 0, int(days*-1)).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func NewLastEmptiedReading() *LastEmptiedReading {
	return &LastEmptiedReading{
		LastEmptied: randate(),
	}
}
