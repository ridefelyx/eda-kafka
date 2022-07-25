package main

import (
	"POC_Kafka/main/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"time"
)

// the topicVehicles and broker address are initialized as constants
const (
	topicVehicles  = "topic-vehicles"
	broker1Address = "localhost:19091"
)

func produce() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", broker1Address, topicVehicles, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	operation := model.StateEvent
	i := 0
	for {
		newMessage := model.Message{
			Operation: operation,
			Payload: model.Vehicle{
				ID:                uuid.New().String(),
				LicensePlate:      randomString(8),
				Latitude:          float64(rand.Intn(90-(-90)) + (-90)),
				Longitude:         float64(rand.Intn(180-(-180)) + (-180)),
				BatteryPercentage: int64(rand.Intn(100)),
			},
		}

		stringJson, _ := json.Marshal(newMessage)

		err := conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
		if err != nil {
			return 
		}
		_, err = conn.WriteMessages(
			kafka.Message{
				Key: []byte(newMessage.Payload.ID),
				Value: stringJson,
				Time: time.Now(),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		if operation < 2 {
			operation++
		} else {
			operation = 0
		}
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second*5)
	}
}