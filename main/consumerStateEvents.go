package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func consumeStateEvents(ctx context.Context) {

	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()
	batchSize := int(10e6) // 10MB

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker1Address},
		Topic:     topicVehicles,
		GroupID:   "location", // if groupID is set SetOffsetAt will throw error
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
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
