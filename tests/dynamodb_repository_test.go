package tests

import (
	"context"
	"person-service/model"
	"person-service/repo"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBClient struct {
	mock.Mock
	dynamodbiface.DynamoDBAPI
}

func (m *MockDynamoDBClient) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	args := m.Called(ctx, input)
	output, _ := args.Get(0).(*dynamodb.PutItemOutput)
	return output, nil // args.Error(1) gets panic
}

func TestCreatePersonDynamoDB(t *testing.T) {
	mockDynamoDB := new(MockDynamoDBClient)
	repository := repo.NewDynamoDBRepository(mockDynamoDB)

	person := model.Person{FirstName: "Cem", LastName: "Tanrikut", PhoneNumber: "9876543210", Address: "3527 RB"}

	mockDynamoDB.On("PutItemWithContext", mock.Anything, mock.Anything).Return(nil)

	err := repository.CreatePerson(context.TODO(), person)

	assert.Nil(t, err)
	mockDynamoDB.AssertExpectations(t)
}
