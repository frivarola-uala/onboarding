package model

import (
	"context"
)

type Contact struct {
	Id string `dynamodbav:"id,omitempty" json:"id,omitempty"`

	FirstName string `dynamodbav:"first_name" json:"first_name"`

	LastName string `dynamodbav:"last_name" json:"last_name"`

	Status string `dynamodbav:"status" json:"status"`
}

type ContactRepository interface {
	AddContact(ctx context.Context, c Contact) error
	GetContact(ctx context.Context, id string) (Contact, error)
}
