package repositories

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"onboarding/internal/model"
)

const (
	InitialStatus = "CREATED"
	TableName     = "Contacts-federico-onboarding"
)

type DynamoContactRepository struct {
	dbClient *dynamodb.Client
}

func NewContactRepository(db *dynamodb.Client) *DynamoContactRepository {
	return &DynamoContactRepository{dbClient: db}
}

func (r *DynamoContactRepository) AddContact(ctx context.Context, c model.Contact) error {
	c.Status = InitialStatus

	item, err := attributevalue.MarshalMap(c)
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

func (r *DynamoContactRepository) GetContact(ctx context.Context, q model.GetContactQuery) (*model.Contact, error) {

	if q.Status == "" {
		q.Status = InitialStatus
	}

	key, err := attributevalue.MarshalMap(q)

	if err != nil {
		log.Printf("Error in marshal %v. Error: %v\n", q.Id, err)
	}

	result, err := r.dbClient.GetItem(ctx,
		&dynamodb.GetItemInput{Key: key, TableName: aws.String(TableName)})

	var contact model.Contact

	if err != nil {
		log.Printf("Couldn't get info about %v. || error: %v\n", q.Id, err)
	} else {
		err = attributevalue.UnmarshalMap(result.Item, &contact)
		if err != nil {
			log.Printf("Couldn't unmarshal result. || error: %v\n", err)
		}
	}

	return &contact, err
}
