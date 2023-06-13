package nosql

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfigMod "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

var awsConfig aws.Config
var dynamodbClient *dynamodb.Client

func getAwsConfig() aws.Config {
	var err error
	awsConfig, err = awsConfigMod.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load AWS config, %v", err)
	}

	return awsConfig
}

func GetDynamodbClient() *dynamodb.Client {
	awsConfig = getAwsConfig()
	region := awsConfig.Region

	dynamodbClient = dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
		opt.Region = region
	})

	return dynamodbClient
}
