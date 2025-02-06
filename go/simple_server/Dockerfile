
# Use the official Go image to build the app
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Install dependencies (and create go.mod file if not already done)
RUN go mod tidy

# Copy the source code into the container (after go.mod)
COPY . .

# Build the Go app
RUN go build -o simple_server .

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./simple_server"]

