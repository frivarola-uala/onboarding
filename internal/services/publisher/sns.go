package publisher

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

type Message struct {
	Default string `json:"default"`
}

var snsClient *sns.SNS

func InitializeSNSClient() {
	//	var cfg *aws.Config
	//	*cfg, _ = awsConfigMod.LoadDefaultConfig(ctx)
	if snsClient == nil {
		mySession := session.Must(session.NewSession())
		snsClient = sns.New(mySession)
	}
}

func PublishMessage(payload interface{}, topicName string) error {

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		log.Printf("error parsing message. %v", err)
	}

	message := Message{
		Default: string(jsonStr),
	}

	messageBytes, _ := json.Marshal(message)

	req, errPublish := snsClient.PublishRequest(&sns.PublishInput{
		TopicArn:         aws.String(topicName),
		Message:          aws.String(string(messageBytes)),
		MessageStructure: aws.String("json"),
	})

	if errPublish != nil {
		log.Printf("error in publish request. %v", errPublish.String())
	}

	err = req.Send()
	if err != nil {
		log.Fatal(err)
	}

	return err
}
