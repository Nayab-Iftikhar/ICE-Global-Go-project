# Build stage
FROM golang:1.22 AS build

WORKDIR /app

# Copy Go Modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary for Linux/amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Final image
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
COPY wait-for-mysql.sh /wait-for-mysql.sh
RUN chmod +x /wait-for-mysql.sh
RUN apk add --no-cache mysql-client
CMD ["/wait-for-mysql.sh", "mysql", "./main"]
