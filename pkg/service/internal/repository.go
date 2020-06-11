package repository

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/lib/firestore"
)

const collectionName = "messages"

type Message interface {
	Save(ctx context.Context, message model.Message) error
	Fetch(ctx context.Context, limit int) ([]*model.Message, error)
}

type message struct {
	client firestore.Client
}

func NewMessage(client firestore.Client) Message {
	return &message{
		client: client,
	}
}

func (m message) Save(ctx context.Context, message model.Message) error {
	_, _, err := m.client.Collection(collectionName).Add(ctx, message)
	return err
}

func (m message) Fetch(ctx context.Context, limit int) ([]*model.Message, error) {
	messages := make([]*model.Message, 0)

	results, err := m.client.
		Collection(collectionName).
		Limit(limit).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		message := new(model.Message)
		err = result.DataTo(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
