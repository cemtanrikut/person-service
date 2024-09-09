package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"person-service/model"
	"person-service/repo"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Build repo
type PersonHandler struct {
	Repository repo.PersonRepository
}

func NewPersonHandler() *PersonHandler {
	// AWS session creating
	sess := session.Must(session.NewSession())

	// DynamoDB client creates with sess
	dynamoClient := dynamodb.New(sess)

	// repo.NewDynamoDBRepository func gets DynamoDB client example
	return &PersonHandler{
		Repository: repo.NewDynamoDBRepository(dynamoClient),
	}
}

func (h *PersonHandler) HandleRequest(ctx context.Context, request model.PersonRequest) (model.PersonResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return h.createPerson(ctx, request)
	default:
		return model.PersonResponse{StatusCode: http.StatusBadRequest}, nil
	}
}

func (h *PersonHandler) createPerson(ctx context.Context, request model.PersonRequest) (model.PersonResponse, error) {
	var person model.Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return model.PersonResponse{StatusCode: http.StatusBadRequest}, err
	}

	err = h.Repository.CreatePerson(ctx, person)
	if err != nil {
		return model.PersonResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return model.PersonResponse{StatusCode: http.StatusCreated}, nil
}
