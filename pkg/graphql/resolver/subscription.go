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
		<-ctx.Done()
		s.PubSub.Unsub(subscription, topicName)
	}()

	go func() {
		for message := range subscription {
			messages <- message.(*model.Message)
		}
	}()

	return messages, nil
}
