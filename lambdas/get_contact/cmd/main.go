package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"net/url"
	"onboarding/internal/model"
	"onboarding/internal/repositories"
	"onboarding/internal/services/storage/nosql"
	"strconv"
)

const IdParam = "id"

var contactsRepository *repositories.DynamoContactRepository

// HandleRequest retrieves a model.Contact based on the provided ID from the URL path "/contacts/{id}
// @param {string} id - The unique identifier of the user/**
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	paramUrlEncode, _ := request.PathParameters[IdParam]
	id, err := url.QueryUnescape(paramUrlEncode)

	if err != nil {
		return buildResponse(http.StatusBadRequest, err)
	}

	log.Printf("user_id: %s", id)

	if err = validateInput(id); err != nil {
		return buildResponse(http.StatusBadRequest, err)
	}

	contact, err := contactsRepository.GetContact(ctx, model.GetContactQuery{Id: id})

	if err != nil {
		log.Printf("Error to get user_id: %s", id)
		return buildResponse(http.StatusBadRequest, err)
	}

	return buildResponse(http.StatusOK, contact)
}

func validateInput(id string) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	if _, err := strconv.Atoi(id); err != nil {
		return fmt.Errorf("id should be numeric")
	}

	return nil
}

func buildResponse(status int, payload interface{}) events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	bytes, err := json.Marshal(payload)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Body = err.Error()
	}

	response.StatusCode = status
	response.Body = string(bytes)

	return response
}

func main() {
	db := nosql.GetDynamodbClient()
	contactsRepository = repositories.NewContactRepository(db)
	lambda.Start(HandleRequest)
}
