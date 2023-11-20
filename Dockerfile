# # Start from a base image with Go installed
# FROM golang:1.20 as builder

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go modules and sum files
# COPY go.mod go.sum ./

# # Download Go module dependencies
# RUN go mod download

# # Copy the rest of the source code
# COPY . .

# # Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# # Use a small base image to create a minimal final image
# FROM alpine:latest  

# # Add certificates for HTTPS communication
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# # Copy the pre-built binary file from the previous stage
# COPY --from=builder /app/main .

# # Expose the port the app runs on
# EXPOSE 8080

# # Command to run the executable
# CMD ["./main"]

# Start from a base image with Go 1.20 installed
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a small base image to create a minimal final image
FROM alpine:latest

# Add certificates for HTTPS communication
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
