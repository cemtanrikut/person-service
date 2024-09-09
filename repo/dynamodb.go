package repo

import (
	"context"
	"person-service/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Interfaces
type PersonRepository interface {
	CreatePerson(ctx context.Context, person model.Person) error
}

type DynamoDBRepository struct {
	client *dynamodb.DynamoDB
}

// Build new dynamo repo
func NewDynamoDBRepository() *DynamoDBRepository {
	sess := session.Must(session.NewSession())
	return &DynamoDBRepository{
		client: dynamodb.New(sess),
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

	_, err = repo.client.PutItemWithContext(ctx, input)
	return err
}
