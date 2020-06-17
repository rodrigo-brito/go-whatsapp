package resolver

import (
	"context"

	"go-zap/pkg/graphql/model"

	"github.com/cskr/pubsub"
)

const topicName = "messages"

type Subscription struct {
	PubSub *pubsub.PubSub
}

func (s Subscription) Messages(ctx context.Context) (<-chan *model.Message, error) {
	subscription := s.PubSub.Sub(topicName)
	messages := make(chan *model.Message)

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.PubSub.Unsub(subscription, topicName)
			case message := <-subscription:
				if message != nil {
					messages <- message.(*model.Message)
				}
			}
		}
	}()

	return messages, nil
}
