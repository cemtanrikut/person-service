# Person Service

Person Service is a serverless microservice designed to manage personal information. It allows users to create new person entries and list existing ones through a REST API.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

```bash
go version go1.16 linux/amd64
AWS CLI v2
AWS account and configured AWS CLI
```

## Installing

A step by step series of examples how to get a development env running:

Clone the repository to your local machine:

```bash
git clone https://github.com/cemtanrikut/person-service.git
cd person-service
```

Install the necessary dependencies:

```bash
make deps
```

## Running the tests

How to run the automated tests for this system:

```bash
make test
```

## Deployment

How to deploy this on a live system:

```bash
make build
```

## Built with

- Go - The primary programming language
- AWS Lambda - Serverless compute service
- DynamoDB - NoSQL database service

## License

This project is licensed under the MIT License