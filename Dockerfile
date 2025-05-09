FROM golang:1.23.0 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./.env ./

RUN swag init -g ./pkg/routes/routes.go -o ./docs

RUN go build -o app ./cmd/service/main.go

EXPOSE 8080

CMD ["./app"]