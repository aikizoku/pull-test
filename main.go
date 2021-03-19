package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	projectID := "develop-aikizoku"
	subID := "pull-test-sub"
	pullMsgs(ctx, projectID, subID)
}

func pullMsgs(ctx context.Context, projectID string, subID string) {
	// PubSubClientを作成
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v\n", err)
		return
	}

	// Subscription(Pull型)を指定
	sub := client.Subscription(subID)

	// レシーブ開始
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// 受け取ったら書き込む
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		fmt.Printf("Receive: %v\n", err)
		return
	}
}
