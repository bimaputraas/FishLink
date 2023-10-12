FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]

# docker build -t go-fishlink-mainapi .
# docker run --name go-fishlink-mainapi-container -p 8080:8080 go-fishlink-mainapi
