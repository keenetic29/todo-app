# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app ./cmd/main.go

# Runtime stage
FROM alpine:latest
WORKDIR /
COPY --from=builder /todo-app /todo-app
# Создаем папку configs если ее нет
RUN mkdir -p /configs
EXPOSE 8080
ENTRYPOINT ["/todo-app"]