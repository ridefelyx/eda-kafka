package main

import (
	"POC_Kafka/main/model"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"time"
)

func produceDoc(ctx context.Context) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", broker1Address, topicVehicles, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	operation := model.StateEvent
	for {
		newMessage := model.Message{
			Operation: operation,
			Payload: model.Vehicle{
				ID:                uuid.New().String(),
				LicensePlate:      randomString(8),
				Latitude:          float64(rand.Intn(90-(-90)) + (-90)),
				Longitude:         float64(rand.Intn(180-(-180)) + (-180)),
				BatteryPercentage: int64(rand.Intn(100)),
				State:             model.Active,
			},
			PublishedAt: time.Now(),
		}

		stringJson, _ := json.Marshal(newMessage)

		err := conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
		if err != nil {
			return 
		}
		_, err = conn.WriteMessages(
			kafka.Message{Value: stringJson},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		if operation < 2 {
			operation++
		} else {
			operation = 0
		}
	}
}