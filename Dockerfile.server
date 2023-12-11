FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY ./server ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main /app/

COPY ./server/.env .

EXPOSE 4000

ENV ENV_FILE_LOCATION /app/.env

CMD ["./main"]
