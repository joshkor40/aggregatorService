package main 

var literToGallonConversion float64 = .264172

type VolumeReading struct {
	Liters  float64 `json:"liters"`
	Gallons float64 `json:"gallons"`
}

// VolumeReading request section

func NewVolumeReading(liters float64) *VolumeReading {
	return &VolumeReading{
		Liters:  liters,
		Gallons: liters * literToGallonConversion,
	}
}
