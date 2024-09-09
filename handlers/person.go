package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"person-service/model"
	"person-service/repo"
)

// Build repo
type PersonHandler struct {
	Repository repo.PersonRepository
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{
		Repository: repo.NewDynamoDBRepository(),
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
