# log_ingestion/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o log_ingestion_service

EXPOSE 8080

CMD ["./log_ingestion_service"]
