# monitoring/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o monitoring_service

EXPOSE 8082

CMD ["./monitoring_service"]
