package resolver

import (
	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/service"
)

var ErrServiceUnavailable = gqlerror.Errorf("service unavailable")

type Mutation struct {
	MessageService service.Message
}

func (m Mutation) SendMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	message := model.Message{
		Author:  input.Author,
		Message: input.Message,
	}

	err := m.MessageService.Save(ctx, message)
	if err != nil {
		return nil, ErrServiceUnavailable
	}

	return &message, nil
}
