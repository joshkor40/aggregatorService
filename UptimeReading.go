package main

import "time"


type UptimeReading struct {
	Uptime float64 `json:"uptime"`
}

func NewUptimeReading() *UptimeReading {
  deltaSeconds := time.Now().Unix() - randTime().Unix()
	uptimeHours := float64(deltaSeconds / 60 / 60) // generate random uptime hours
	return &UptimeReading{
		Uptime:  uptimeHours,
	}
}