package main

import (
	"POC_Kafka/main/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"math/rand"
	"strconv"
	"time"
)

// the topicVehicles and broker address are initialized as constants
const (
	topicVehicles  = "vehicles"
	broker1Address = "localhost:9092"
	broker2Address = "localhost:9093"
)

func produce(ctx context.Context) {
	//conn, err := kafka.Dial("tcp", broker1Address)
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer conn.Close()
	//
	//controller, err := conn.Controller()
	//if err != nil {
	//	panic(err.Error())
	//}
	//var controllerConn *kafka.Conn
	//controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer controllerConn.Close()
	//
	//
	//topicConfigs := []kafka.TopicConfig{
	//	{
	//		Topic:             topicVehicles,
	//		NumPartitions:     1,
	//		ReplicationFactor: 1,
	//	},
	//}
	//
	//err = controllerConn.CreateTopics(topicConfigs...)
	//if err != nil {
	//	panic(err.Error())
	//}
		i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topicVehicles,
	})

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
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: stringJson,
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		if operation < 2 {
			operation++
		} else {
			operation = 0
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++

		time.Sleep(time.Second * 3)
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
