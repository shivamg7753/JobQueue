# Start from the official Golang image for build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd && go build -o /jobqueue

# Use a minimal image for running
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /jobqueue .
EXPOSE 8080
CMD ["./jobqueue"]
