# user_service/Dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o user_service
EXPOSE 8001
CMD ["./user_service"]
