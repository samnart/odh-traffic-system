FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o traffic-service

EXPOSE 8083
CMD ["./traffic-service"]