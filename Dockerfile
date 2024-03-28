# This stage will be used only to build our binary
ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine as build_base

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . /app

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app 

# This stage will be our final image to be used on the cluster - With a multi-stage build approach we reduce the number of layers and the final image size
FROM alpine:edge

WORKDIR /app
COPY --from=build_base /go-app .

# Define our env vars to be used by the app while running
ENV PSQL_USERNAME="foo"
ENV PSQL_PASSWORD="bar"
ENV PSQL_HOST="localhost"
ENV PSQL_PORT="5432"
ENV PSQL_DBNAME="baz"

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/go-app"]