FROM golang:1.20.4-alpine

# Set working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod go.sum ./

# Download go modules
RUN go mod download

# Copy the entire project
# This command copies your entire project into the Docker image.
COPY . .

# Build the Go app
RUN go build -o http_go cmd/http/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the PORT environment variable
ENV HTTP_PORT 8080

CMD ["./http_go"]