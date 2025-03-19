# Use a newer Go version that satisfies go.mod requirements
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port (adjust if needed)
EXPOSE 8081

# Command to run the application
CMD ["./main"]
