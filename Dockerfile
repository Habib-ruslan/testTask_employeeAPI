FROM golang:1.23.4

WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Компилируем приложение
RUN go build -o web ./cmd/web
RUN go build -o cli ./cmd/cli

CMD ["./web"]

EXPOSE 8000