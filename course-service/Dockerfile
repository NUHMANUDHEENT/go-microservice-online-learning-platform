FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

WORKDIR /app/course-service

RUN go build -o course-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/course-service/course-service .
CMD ["./course-service"]
