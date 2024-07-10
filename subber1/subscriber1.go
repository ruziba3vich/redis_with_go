package main

import (
	"context"// Initialize a Redis client

	"fmt"

	"github.com/go-redis/redis/v8"
)

// Initialize a Redis client
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	subscribe("mychannel")
}

func subscribe(channel string) {
	pubsub := rdb.Subscribe(ctx, channel)
	defer pubsub.Close()

	ch := pubsub.Channel()

	fmt.Println("Subscriber 1 is listening...")
	for msg := range ch {
		fmt.Printf("Subscriber 1 received message: %s from channel: %s\n", msg.Payload, msg.Channel)
	}
}
