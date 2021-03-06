// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Message struct {
	// Author's UUID
	AuthorID string `json:"authorID"`
	// Author of message
	Author string `json:"author"`
	// Message content
	Content string `json:"content"`
	// Date of creation
	CreatedAt time.Time `json:"createdAt"`
}

type MessageInput struct {
	// Author's UUID
	AuthorID string `json:"authorID"`
	// Author's name
	Author string `json:"author"`
	// Message content
	Message string `json:"message"`
}
