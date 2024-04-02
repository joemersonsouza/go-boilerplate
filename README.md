# Compose a docker file
`docker compose up`

# Build docker image

`docker build -t go-app:latest --build-arg GO_VERSION=1.22.1 .`

# Run docker image
`docker run -p 8080:8080 go-app:latest`

# Helpful commands

- Removing containers
`docker rm -vf $(docker ps -aq)`

- Removing Images
`docker rmi -f $(docker images -aq)`