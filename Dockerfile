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

# Copy swagger UI
COPY swagger-ui/ /app/swagger-ui/

# Serve static files (simple HTTP)
RUN apk add --no-cache busybox-extras
CMD busybox httpd -f -p 8081 -h /app/swagger-ui
