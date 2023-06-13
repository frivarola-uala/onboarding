package model

import (
	"context"
)

type Contact struct {
	Id string `json:"id"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Status string `json:"status"`
}

type ContactRepository interface {
	AddContact(ctx context.Context, c Contact) error
}
