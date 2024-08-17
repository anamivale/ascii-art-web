# Dockerfile

# Use an official Golang runtime as a parent image
FROM golang:1.22.2-alpine

LABEL maintainer github.com/anamivale

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any needed dependencies
RUN go mod download

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the app by default when the container starts
CMD ["go", "run", "main.go"]
