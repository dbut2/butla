FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY ./vendor ./vendor
COPY ./pkg ./pkg
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./web ./web

RUN go build -o app ./web

FROM alpine

ENV PORT=8080
EXPOSE ${PORT}

WORKDIR /app

COPY --from=builder /app/app ./web

CMD ["./web"]
