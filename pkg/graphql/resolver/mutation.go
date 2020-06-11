package resolver

import (
	"context"
	"time"

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
		AuthorID:  input.AuthorID,
		Author:    input.Author,
		Content:   input.Message,
		CreatedAt: time.Now(),
	}

	err := m.MessageService.Save(ctx, message)
	if err != nil {
		return nil, ErrServiceUnavailable
	}

	return &message, nil
}
