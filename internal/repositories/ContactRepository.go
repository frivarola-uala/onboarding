package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"onboarding/internal/model"
	"onboarding/internal/services/storage/nosql"
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

	log.Printf("contact: %v", &c)
	item, err := attributevalue.MarshalMap(c)
	log.Printf("item: %v", item)

	if err != nil {
		log.Fatalf("Error parsing contact. || err: %v", err)
	}
	_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(TableName),
	})

	if err != nil {
		log.Printf("Couldn't add item to table. || err: %v\n", err)
	}

	return err
}

func (r *ContactRepository) GetContact(ctx context.Context, id string) (*model.Contact, error) {

	key, err := attributevalue.MarshalMap(id)

	if err != nil {
		log.Printf("Error in marshal %v. Error: %v\n", id, err)
	}

	response, err := r.dbClient.GetItem(ctx,
		&dynamodb.GetItemInput{Key: key, TableName: &TableName})

	var contact model.Contact

	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", id, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &contact)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	return &contact, err
}
