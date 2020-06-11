//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"go-zap/pkg/graphql"
	"go-zap/pkg/service"

	"github.com/cskr/pubsub"
)

type GraphQL struct {
	MessageService service.Message
	PubSub         *pubsub.PubSub
}

func (g GraphQL) Mutation() graphql.MutationResolver {
	return &Mutation{
		MessageService: g.MessageService,
	}
}

func (g GraphQL) Query() graphql.QueryResolver {
	return &Query{
		MessageService: g.MessageService,
	}
}

func (g GraphQL) Subscription() graphql.SubscriptionResolver {
	return &Subscription{
		PubSub: g.PubSub,
	}
}
