# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY .env .
COPY wait-for.sh .
COPY start.sh .
COPY migrations ./migration

EXPOSE 4000

 CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
