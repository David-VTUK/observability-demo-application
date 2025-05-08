# Stage 1: Build the Go application
FROM golang:1.23 as builder

# Set architecture-specific build argument
ARG TARGETARCH

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o app ./cmd/main.go

# Stage 2: Create a minimal runtime container
FROM scratch

# Set the working directory inside the container
WORKDIR /

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./app"]