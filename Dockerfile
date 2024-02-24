# # Start from a Golang base image
# FROM golang:1.17-alpine AS builder

# # Set the current working directory inside the container
# WORKDIR /app

# # # Copy the Go modules manifests
# # COPY go.mod .
# # COPY go.sum .

# # # Download dependencies
# # RUN go mod download

# # Copy the source code into the container
# COPY . .

# # Build the Go app
# RUN go build -o main main.go

# # Start a new stage from scratch
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the pre-built binary from the previous stage
# COPY --from=builder /app/main .

# # Expose port 8080
# EXPOSE 8080

# # Command to run the executable
# CMD ["./app/main"]


FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
# Expose port 8080
EXPOSE 8080
CMD ["./out/dist"]