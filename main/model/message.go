package model

type Message struct {
	Operation   Operation
	Payload     Vehicle
}

type Operation int

const (
	StateEvent Operation = iota
	BatteryEvent
	LocationEvent
)
