package resolver

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/service"
)

type Mutation struct {
	MessageService service.Message
}

func (m Mutation) SendMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	panic("implement me")
}
