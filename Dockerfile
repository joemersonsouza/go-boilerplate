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
COPY . ./

# Build
RUN go build -o go-app .

# This stage will be our final image to be used on the cluster - With a multi-stage build approach we reduce the number of layers and the final image size
FROM scratch

WORKDIR /app
COPY --from=build_base /app/go-app .

# Define our env vars to be used by the app while running
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=admin
ENV DB_NAME=notificationdb
ENV DB_PASSWORD=admin

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
ENTRYPOINT [ "/app/go-app" ]