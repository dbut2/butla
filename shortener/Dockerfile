FROM golang:alpine AS builder

WORKDIR /app

COPY ./vendor ./vendor
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o shortener ./cmd/shortener

FROM alpine

ENV PORT=8080
EXPOSE ${PORT}

WORKDIR /app

COPY --from=builder /app/shortener ./shortener

CMD ["./shortener"]
