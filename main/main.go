package main

import (
	"context"
)

const consumerCount int = 2

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking

	groups := make(map[int]string)

	groups[1] = "location"
	groups[2] = "battery"

	done := make(chan bool)

	go produce(ctx)

	for i := 1; i <= consumerCount; i++ {
		go consume(ctx, groups[i], done)
	}
	<-done
}
