# Use the official Golang image as the base image
FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build ./cmd/analytics

# Use slim alpine image for production
FROM alpine:3.18 as production
WORKDIR /app
COPY --from=builder /app/analytics .
EXPOSE 3000

# Run the Go program when the container starts
CMD ["./analytics"]
