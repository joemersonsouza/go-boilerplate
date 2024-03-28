# Compose a docker file
`docker compose up`

# Build docker image

`docker build -t notification:latest .`

# Run docker image
`docker run -p 8080:8080 notification:latest`

# Helpful commands

- Removing containers
`docker rm -vf $(docker ps -aq)`

- Removing Images
`docker rmi -f $(docker images -aq)`