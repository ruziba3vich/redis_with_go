package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// Initialize a Redis client
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	for {
		var message string
		fmt.Scan(&message)
		err := publish("mychannel", message)
		if err != nil {
			log.Fatalf("couldn't send message: %v", err)
		}
	}
}

func publish(channel, message string) error {
	err := rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	fmt.Printf("Published message: %s to channel: %s\n", message, channel)
	return nil
}
