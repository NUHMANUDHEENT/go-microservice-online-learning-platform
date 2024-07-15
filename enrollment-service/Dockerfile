FROM golang:1.22-alpine as builder

WORKDIR /app


COPY . .

WORKDIR /app/enrollment-service

RUN go build -o enrollment-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/enrollment-service/enrollment-service .
CMD ["./enrollment-service"]
