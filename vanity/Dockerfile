FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o /bin/vanity-server

FROM alpine:latest

COPY --from=builder /bin/vanity-server .

EXPOSE 8080

CMD ["./vanity-server"]
