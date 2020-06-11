package main

import (
	"context"
	"go-zap/pkg/lib/firestore"
	"go-zap/pkg/service"
	"net/http"

	"go-zap/pkg/graphql"
	"go-zap/pkg/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "go-whatsapp-2166d")
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./website/build"))
	playgroundHandler := playground.Handler("GraphQL", "/graphql")
	graphQLHandler := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &resolver.GraphQL{
			MessageService: service.NewMessage(client),
		},
	}))

	server.Handle("/", fileServer)
	server.Handle("/graphql", graphQLHandler)
	server.Handle("/graphql/explorer", playgroundHandler)

	log.Info("Listen http://localhost:8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
