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
	repository repo.PersonRepository
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{
		repository: repo.NewDynamoDBRepository(),
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

	err = h.repository.CreatePerson(ctx, person)
	if err != nil {
		return model.PersonResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return model.PersonResponse{StatusCode: http.StatusCreated}, nil
}
