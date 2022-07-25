package model

type Vehicle struct {
	ID                string  `json:"id"`
	LicensePlate      string  `json:"license_plate"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	BatteryPercentage int64   `json:"battery_percentage"`
}
