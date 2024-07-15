
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

WORKDIR /app/user-service

RUN go build -o user-service ./cmd


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service/user-service .
CMD ["./user-service"]
