# book_service/Dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o book_service
EXPOSE 8000
CMD ["./book_service"]
