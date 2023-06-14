package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"onboarding/model"
	"onboarding/repositories"
)

func HandleGetRequest(ctx context.Context, id string) (*model.Contact, error) {
	contacts := repositories.NewContactRepository()
	contact, err := contacts.GetContact(ctx, id)

	if err != nil {
		log.Printf("Error to get user_id: %s", id)
		return nil, err
	}

	return contact, err
}

func main() {
	lambda.Start(HandleGetRequest)
}
