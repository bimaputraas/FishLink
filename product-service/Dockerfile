FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 50051

ENTRYPOINT ["/app/binary"]

# docker build product-service -t product-service
# docker container create --name go-fishlink-product-service-container -p 50051:50051 go-fishlink-product-service
