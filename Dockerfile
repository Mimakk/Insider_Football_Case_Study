# Use official Go image as base
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod ./
RUN go mod download

# Copy entire project
COPY . .

# Build the app
RUN go build -o league-sim main.go

# Expose port
EXPOSE 8080

# Run the app
CMD ["./league-sim"]
