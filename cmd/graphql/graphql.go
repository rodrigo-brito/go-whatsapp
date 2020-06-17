package main

import (
	"context"
	"net/http"
	"time"

	"go-zap/pkg/graphql"
	"go-zap/pkg/graphql/resolver"
	"go-zap/pkg/lib/firestore"
	"go-zap/pkg/service"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cskr/pubsub"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	pubSub := pubsub.New(0)
	client, err := firestore.NewClient(ctx, "go-whatsapp-2166d")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir("./website/build"))
	playgroundHandler := playground.Handler("GraphQL", "/graphql")
	graphQLHandler := handler.New(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &resolver.GraphQL{
			MessageService: service.NewMessage(client, pubSub),
			PubSub:         pubSub,
		},
	}))

	graphQLHandler.SetQueryCache(lru.New(100))
	graphQLHandler.Use(extension.AutomaticPersistedQuery{Cache: lru.New(100)})
	graphQLHandler.AddTransport(transport.Options{})
	graphQLHandler.AddTransport(transport.GET{})
	graphQLHandler.AddTransport(transport.POST{})
	graphQLHandler.AroundOperations(func(ctx context.Context, next gqlgen.OperationHandler) gqlgen.ResponseHandler {
		gqlgen.GetOperationContext(ctx).DisableIntrospection = false
		return next(ctx)
	})

	graphQLHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.AllowAll().Handler)
	router.Handle("/*", fileServer)
	router.Handle("/graphql", graphQLHandler)
	router.Handle("/graphql/explorer", playgroundHandler)

	log.Info("Listen http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
