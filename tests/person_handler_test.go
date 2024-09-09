package tests

import (
	"context"
	"encoding/json"
	"person-service/handlers"
	"person-service/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePerson(ctx context.Context, person model.Person) error {
	args := m.Called(ctx, person)
	return args.Error(0)
}

func TestCreatePerson(t *testing.T) {
	mockRepo := new(MockRepository)
	handler := handlers.PersonHandler{Repository: mockRepo}
	person := model.Person{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "1234567890",
		Address:     "1234 Main St",
	}

	body, _ := json.Marshal(person)
	request := model.PersonRequest{
		HTTPMethod: "POST",
		Body:       string(body),
	}

	mockRepo.On("CreatePerson", mock.Anything, person).Return(nil)

	response, err := handler.HandleRequest(context.Background(), request)

	assert.Nil(t, err)
	assert.Equal(t, 201, response.StatusCode)
	mockRepo.AssertExpectations(t)
}
