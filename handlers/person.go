package handlers

import "person-service/repo"

// Build repo
type PersonHandler struct {
	repository repo.PersonRepository
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{
		repository: repo.NewDynamoDBRepository(),
	}
}
