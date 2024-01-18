# Used the official golang:latest base image instead of ubuntu:latest to have a smaller image and ensure a consistent development environment.
FROM golang:latest

WORKDIR /app

# Removed unnecessary installation of git, go, and wget since they are already included in the golang image.

COPY . .

# Build the application
RUN go build -o main .

#Simplified the CMD instruction for running the application.
CMD ["./main"]
