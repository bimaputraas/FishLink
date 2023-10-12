FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]

# docker build -t go-fishlink-notification-service .
# docker container create --name go-fishlink-notification-service-container -p 50051:50051 go-fishlink-notification-service
