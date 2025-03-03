# Dockerfile
FROM golang:1.23-alpine3.21 AS builder

RUN apk update && apk add --no-cache git bash build-base
WORKDIR /app

COPY go.mod go.sum ./
# next line is temp
COPY .env ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:3.21
RUN apk add --no-cache bash curl postgresql-client
WORKDIR /root

COPY --from=builder /app/main .
# next line is temp
COPY --from=builder /app/.env .
COPY script/database.sh /root
COPY script/migrate.sh /root
COPY script/start.sh /root
RUN chmod +x /root/start.sh

COPY internal/migration/postgres ./internal/migration/postgres

EXPOSE 7700
ENTRYPOINT ["/root/start.sh"]