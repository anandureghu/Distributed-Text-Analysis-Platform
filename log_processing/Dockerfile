# log_processing/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o log_processing_service

EXPOSE 50051

CMD ["./log_processing_service"]
