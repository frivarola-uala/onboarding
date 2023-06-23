package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"onboarding/internal/model"
	"onboarding/internal/repositories"
	"onboarding/internal/services/storage/nosql"
)

var contactsRepository *repositories.DynamoContactRepository

// HandleRequest create a new contact based on the provided data from the payload "/contacts/{id}
// @param {model.Contact} newContact - data of the new contact/**
func HandleRequest(ctx context.Context, newContact model.Contact) (string, error) {

	log.Printf("user_id: %s", newContact.Id)

	if err := contactsRepository.AddContact(ctx, newContact); err != nil {
		log.Printf("Error to add user_id: %s", newContact.Id)
		return fmt.Sprintf("Error to add user_id: %s", newContact.Id), err
	}

	return fmt.Sprintf("Success. user_id: %s", newContact.Id), nil
}

func main() {
	db := nosql.GetDynamodbClient()
	contactsRepository = repositories.NewContactRepository(db)
	lambda.Start(HandleRequest)
}
