package model

type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type PersonRequest struct {
	HTTPMethod string `json:"httpMethod"`
	Body       string `json:"body"`
}

type PersonResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}
