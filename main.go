package main

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
)

var (
	topic      *pubsub.Topic
	messagesMu sync.Mutex
	messages   []string
)

const maxMessages = 10

func main() {
	projectID := "develop-aikizoku"
	subID := "pull-test-sub"
	pullMsgs(projectID, subID)
}

func pullMsgs(projectID string, subID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v\n", err)
		return
	}
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		fmt.Printf("Receive: %v\n", err)
		return
	}
}
