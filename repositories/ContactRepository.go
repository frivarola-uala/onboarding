package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"onboarding/model"
	"onboarding/services/storage/nosql"
)

const InitialStatus = "CREATED"

var TableName = "Contacts-federico-onboarding"

type ContactRepository struct {
	dbClient *dynamodb.Client
}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{dbClient: nosql.GetDynamodbClient()}
}

func (r *ContactRepository) AddContact(ctx context.Context, c model.Contact) error {
	c.Status = InitialStatus
	item, err := attributevalue.MarshalMap(c)

	fmt.Printf("item: %+v", item)
	if err != nil {
		log.Fatalf("Error parsing contact. || err: %v", err)
	}
	_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: &TableName,
	})

	if err != nil {
		log.Printf("Couldn't add item to table. || err: %v\n", err)
	}

	return err
}
