package hogehoge

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	projectID := "my-project-id"
	topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %w", err)
		return
	}
	defer client.Close()

	var wg sync.WaitGroup
	var totalErrors uint64
	t := client.Topic(topicID)

	for i := 0; i < 1; i++ {
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Message " + strconv.Itoa(i)),
		})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			// The Get method blocks until a server-generated ID or
			// an error is returned for the published message.
			id, err := res.Get(ctx)
			if err != nil {
				// Error handling code can be added here.
				fmt.Printf("Failed to publish: %v", err)
				atomic.AddUint64(&totalErrors, 1)
				return
			}
			fmt.Printf("Published message %d; msg ID: %v\n", i, id)
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		fmt.Printf("%d of %d messages did not publish successfully", totalErrors, 1)
		return
	}
	return
}
