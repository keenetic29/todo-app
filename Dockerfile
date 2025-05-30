# Билд стадия
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/main.go

# Финальная стадия
FROM alpine:latest
WORKDIR /app

# Копируем бинарник и конфиги
COPY --from=builder /app/todo-app .
COPY --from=builder /app/conf.example.env ./conf.env  

EXPOSE 8080
CMD ["./todo-app"]