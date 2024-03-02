# Use an official Golang runtime as a parent image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container's workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the local package files to the container's workspace
COPY . .

# Build the Go application inside the container
RUN go build -o task

