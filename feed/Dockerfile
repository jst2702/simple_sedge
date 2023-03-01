# Build stage
FROM golang:1.19-alpine3.16 AS builder

RUN apk add --no-cache make bash

WORKDIR /app
COPY . .
RUN go install

RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY db/migrations ./migrations

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]