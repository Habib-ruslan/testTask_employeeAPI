FROM golang:1.23.4

WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Компилируем приложение
RUN go build -o main .

CMD ["./main"]

EXPOSE 8000