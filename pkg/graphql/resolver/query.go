package resolver

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/service"
)

type Query struct {
	MessageService service.Message
}

func (q Query) Messages(ctx context.Context, limit int) ([]*model.Message, error) {
	panic("implement me")
}
