FROM golang:1.23.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Установка PostgreSQL клиента для pg_isready
RUN apt-get update && \
    apt-get install -y postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./.env ./
COPY ./wait-for-db.sh ./

RUN chmod +x ./wait-for-db.sh

RUN swag init -g ./pkg/routes/routes.go -o ./docs

RUN go build -o app ./cmd/service/main.go

EXPOSE 8080

CMD ["./wait-for-db.sh", "./app"]