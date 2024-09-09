package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return handlePost(request)
	case "GET":
		return handleGet(request)
	default:
		return apiResponse(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func handlePost(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return apiResponse(http.StatusBadRequest, "Invalid request body")
	}
	fmt.Printf("Saving person: %+v\n", person)
	return apiResponse(http.StatusOK, "Person created")
}

func handleGet(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Burada tüm kişileri listeleme işlemi yapılacaktır
	people := []Person{
		{FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890", Address: "123 Elm St"},
		{FirstName: "Jane", LastName: "Doe", PhoneNumber: "0987654321", Address: "456 Oak St"},
	}
	responseBody, err := json.Marshal(people)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, "Error marshalling response")
	}
	return apiResponse(http.StatusOK, string(responseBody))
}

func apiResponse(status int, body string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
