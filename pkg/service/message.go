package service

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/lib/firestore"
	repository "go-zap/pkg/service/internal"
)

type Message interface {
	Save(ctx context.Context, message model.Message) error
	Fetch(ctx context.Context, limit int) ([]model.Message, error)
}

type message struct {
	messageRepository repository.Message
}

func NewMessage(client firestore.Client) Message {
	return &message{
		messageRepository: repository.NewMessage(client),
	}
}

func (m message) Save(ctx context.Context, message model.Message) error {
	return m.messageRepository.Save(ctx, message)
}

func (m message) Fetch(ctx context.Context, limit int) ([]model.Message, error) {
	return m.messageRepository.Fetch(ctx, limit)
}
