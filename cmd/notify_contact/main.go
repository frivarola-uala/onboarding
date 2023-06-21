package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"onboarding/internal/model"
	"onboarding/internal/services/publisher"
)

const ContactsTopicName = "arn:aws:sns:us-east-1:620097380428:ContactsTopicFede"

// HandleRequest push message to SNS Topic triggered by DynamoDB"/contacts/{id}
// @param {model.Contact} newContact - data of the new contact/**
func HandleRequest(ctx context.Context, e events.DynamoDBEvent) (string, error) {

	log.Printf("event: %v", e)
	var c model.Contact
	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		// Print new values for attributes name and age
		c.Id = record.Change.NewImage["id"].String()
		c.FirstName = record.Change.NewImage["first_name"].String()
		c.LastName = record.Change.NewImage["last_name"].String()
	}

	log.Printf("contact: %v", c)
	publisher.PublishMessage(c, ContactsTopicName)

	return fmt.Sprintf("Success."), nil
}

func main() {
	publisher.InitializeSNSClient()
	lambda.Start(HandleRequest)
}
