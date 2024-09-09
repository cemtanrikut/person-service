package main

import (
	"person-service/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := handlers.NewPersonHandler()
	lambda.Start(handler.HandleRequest)
}
