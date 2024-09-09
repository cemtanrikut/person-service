package repo

import (
	"context"
	"person-service/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Interfaces
type PersonRepository interface {
	CreatePerson(ctx context.Context, person model.Person) error
}

type DynamoDBRepository struct {
	Client dynamodbiface.DynamoDBAPI // Using the interface type here
}

func NewDynamoDBRepository(client dynamodbiface.DynamoDBAPI) *DynamoDBRepository {
	return &DynamoDBRepository{
		Client: client,
	}
}

func (repo *DynamoDBRepository) CreatePerson(ctx context.Context, person model.Person) error {
	av, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Persons"),
		Item:      av,
	}

	_, err = repo.Client.PutItemWithContext(ctx, input)
	return err
}
