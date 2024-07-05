# Указываем базовый образ
FROM golang:1.20-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN go build -o /go-distributed-cache

# Создаем минималистичный финальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /root/

# Копируем скомпилированное приложение из builder
COPY --from=builder /go-distributed-cache .

# Экспонируем порт
EXPOSE 8080

# Команда для запуска приложения
CMD ["./go-distributed-cache"]