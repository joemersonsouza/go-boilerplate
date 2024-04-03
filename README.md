# Go REST API Boilerplate with PostgreSQL and Kubernetes
This boilerplate is a starting point for building RESTful APIs in Go using PostgreSQL as the database and Kubernetes for deployment.

# Features
- Structured project layout for REST API development
- PostgreSQL integration with database migrations
- Containerization with Docker
- Kubernetes deployment configuration files included

# Prerequisites
- Go
- Docker
- kubectl (Kubernetes command-line tool) (K3s ou K8s)
- Access to a Kubernetes cluster
- Getting Started

## Run tests
`go test -v ./tests`

## Compose a docker file
`docker compose up`

## Build docker image

`docker build -t go-app:latest --build-arg GO_VERSION=1.22.1 .`

## Run docker image
`docker run -p 8080:8080 go-app:latest`

## Helpful commands

- Removing containers
`docker rm -vf $(docker ps -aq)`

- Removing Images
`docker rmi -f $(docker images -aq)`

# API Endpoints

Some endpoints were created to facilitate the implementation of new features. You can use as start point.

- POST

`curl --location 'localhost:8080/notifications/add' \
--header 'Content-Type: application/json' \
--data '{
    "userId": "8fd9e4d5-bfe0-4b6d-8f85-63d77d13cc43",
    "companyId": "b8af89c8-5ebb-42fb-984d-9bf18619feb7",
    "message": "You have a new sale, please follow the link https://mywebsite.sales/product-123"
}'`

- GET

`
curl --location --request GET 'localhost:8080/notifications/user/8fd9e4d5-bfe0-4b6d-8f85-63d77d13cc43?state=ALL' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21wYW55SWQiOiJhNzcyOTg3NS1iZjAxLTQ5MzQtOTk1NC00YmIwNDIyMjg0YzkiLCJleHAiOjE3MTIyNDUyMzV9.W5YVUxg1wBd9fxfH3VhoEJS1VJqK0Zqn_QTb8odM-Oo' \
--data '{
    "userId": "8fd9e4d5-bfe0-4b6d-8f85-63d77d13cc43",
    "companyId": "b8af89c8-5ebb-42fb-984d-9bf18619feb7",
    "message": "Hello World 2"
}'
`
- PUT

`curl --location --request PUT 'localhost:8080/notifications/e6aa0c19-9d91-4def-8454-ef73608e14a7/read' \
--header 'Content-Type: application/json' \
--data '{
    "userId": "8fd9e4d5-bfe0-4b6d-8f85-63d77d13cc43",
    "companyId": "b8af89c8-5ebb-42fb-984d-9bf18619feb7",
    "message": "Hello World 2"
}'`

- DELETE

`curl --location --request DELETE 'localhost:8080/notifications/user/8fd9e4d5-bfe0-4b6d-8f85-63d77d13cc43' \
--header 'Content-Type: application/json' \
--data '["-- or 1=1"]'`

# Recomended Endpoints

- Get Health status

`curl --location 'localhost:8080/health/status'`

- Generate JWT Token

`curl --location --request POST 'localhost:8080/auth/token' \
--header 'Api-Key: 1bbd59c5-d75a-4729-93dc-deb35b16547c'`


# Contributing
Feel free to contribute to this boilerplate by submitting a pull request with your improvements.

# License
This project is licensed under the MIT License.