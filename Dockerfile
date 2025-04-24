# Use official Go image
FROM golang:1.24-alpine

# Create working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8081
EXPOSE 8081

# Run the application
CMD ["./main"]
