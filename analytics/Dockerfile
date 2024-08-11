# analytics/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o analytics_service

EXPOSE 8081
EXPOSE 50053

CMD ["./analytics_service"]
