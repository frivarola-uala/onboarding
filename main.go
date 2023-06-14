package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"onboarding/model"
	"onboarding/repositories"
)

func HandleRequest(ctx context.Context, newContact model.Contact) (string, error) {
	contacts := repositories.NewContactRepository()
	if err := contacts.AddContact(ctx, newContact); err != nil {
		return fmt.Sprintf("Error to add user_id: %s", newContact.Id), err
	}

	return fmt.Sprintf("Success. user_id: %s", newContact.Id), nil
}

func main() {
	lambda.Start(HandleRequest)
}
