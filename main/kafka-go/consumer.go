package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func consume(ctx context.Context, groupID string) {
	fmt.Printf("%s consumer started \n\n", groupID)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topicVehicles,
		GroupID: groupID,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Printf("received event for %s - at %s: \n%s \n\n", groupID, msg.Time, string(msg.Value))
	}
}


func consumeInTimeFrame(ctx context.Context) {
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	batchSize := int(10e6) // 10MB

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topicVehicles,
		//GroupID:   "location", // if groupID is set SetOffsetAt will throw error
		Partition: 0,
		MinBytes:  batchSize,
		MaxBytes:  batchSize,
	})

	err := r.SetOffsetAt(context.Background(), startTime)
	if err != nil {
		fmt.Println("cannot set offset if groupID is set")
		return
	}

	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			break
		}
		if m.Time.After(endTime) {
			//break
		}
		fmt.Printf("received event at offset %v - at %s: \n%s \n\n", m.Offset, m.Time, string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

