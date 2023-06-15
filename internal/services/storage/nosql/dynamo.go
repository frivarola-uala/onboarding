package nosql

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfigMod "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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

func GetDynamodbLocalClient() *dynamodb.Client {
	cfg, err := awsConfigMod.LoadDefaultConfig(context.TODO(),
		awsConfigMod.WithRegion("us-east-1"),
		awsConfigMod.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		awsConfigMod.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func GetDynamodbClient() *dynamodb.Client {
	awsConfig = getAwsConfig()
	region := awsConfig.Region

	dynamodbClient = dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
		opt.Region = region
	})

	return dynamodbClient
}
