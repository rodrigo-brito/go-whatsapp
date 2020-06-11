package firestore

import (
	"context"

	"cloud.google.com/go/firestore"

	firebase "firebase.google.com/go"
)

type Client interface {
	Collection(path string) *firestore.CollectionRef
}

func NewClient(ctx context.Context, projectID string) (Client, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID})
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}
