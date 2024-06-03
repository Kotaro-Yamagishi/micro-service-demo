# order_service/Dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o order_service
EXPOSE 8002
CMD ["./order_service"]
