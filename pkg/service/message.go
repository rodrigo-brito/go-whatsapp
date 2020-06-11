package service

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/lib/firestore"
	repository "go-zap/pkg/service/internal"

	"github.com/cskr/pubsub"
)

const topicName = "messages"

type Message interface {
	Save(ctx context.Context, message model.Message) error
	Fetch(ctx context.Context, limit int) ([]*model.Message, error)
}

type message struct {
	messageRepository repository.Message
	pubSub            *pubsub.PubSub
}

func NewMessage(client firestore.Client, pubSub *pubsub.PubSub) Message {
	return &message{
		messageRepository: repository.NewMessage(client),
		pubSub:            pubSub,
	}
}

func (m message) Save(ctx context.Context, message model.Message) error {
	err := m.messageRepository.Save(ctx, message)
	if err != nil {
		return err
	}
	m.pubSub.TryPub(&message, topicName)
	return nil
}

func (m message) Fetch(ctx context.Context, limit int) ([]*model.Message, error) {
	return m.messageRepository.Fetch(ctx, limit)
}
