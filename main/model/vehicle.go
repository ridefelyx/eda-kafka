package model

type Vehicle struct {
	ID                string
	LicensePlate      string
	Latitude          float64
	Longitude         float64
	BatteryPercentage int64
	State             State
}

type State int

const (
	Active State = iota
	Damaged
	InWorkshop
	InTransport
)
