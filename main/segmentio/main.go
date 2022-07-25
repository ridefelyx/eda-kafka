package main

import (
	"context"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking

	groups := make(map[int]string)

	groups[0] = "location"
	groups[1] = "battery"

	done := make(chan bool)

	go produce()

	for i := 0; i < len(groups); i++ {
		go consume(ctx, groups[i])
		//go consumeInTimeFrame(ctx)
	}
	<-done
}
