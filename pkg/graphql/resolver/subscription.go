package resolver

import (
	"context"

	"go-zap/pkg/graphql/model"
	"go-zap/pkg/service"
)

type Subscription struct {
	MessageService service.Message
}

func (s Subscription) Messages(ctx context.Context) (<-chan *model.Message, error) {
	panic("implement me")
}
