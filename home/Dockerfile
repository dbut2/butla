FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.mod
COPY main.go main.go

RUN go build -o /app/main

FROM alpine AS final

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
