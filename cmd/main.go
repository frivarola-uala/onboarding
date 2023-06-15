package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"onboarding/internal/model"
	"onboarding/internal/repositories"
)

func HandleRequest(ctx context.Context, newContact model.Contact) (string, error) {
	contacts := repositories.NewContactRepository()

	log.Printf("user_id: %s", newContact.Id)

	if err := contacts.AddContact(ctx, newContact); err != nil {
		log.Printf("Error to add user_id: %s", newContact.Id)
		return fmt.Sprintf("Error to add user_id: %s", newContact.Id), err
	}

	return fmt.Sprintf("Success. user_id: %s", newContact.Id), nil
}

func main() {
	lambda.Start(HandleRequest)
}
