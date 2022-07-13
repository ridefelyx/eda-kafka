package model

import "time"

type Message struct {
	Operation   Operation
	Payload     Vehicle
	PublishedAt time.Time
}

type Operation int

const (
	StateEvent Operation = iota
	BatteryEvent
	LocationEvent
)
