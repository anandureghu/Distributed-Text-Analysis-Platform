# data_storage/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o data_storage_service

EXPOSE 50052

CMD ["./data_storage_service"]
