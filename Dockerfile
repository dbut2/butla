FROM golang:alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

COPY ./main.go ./main.go

RUN go build -o /build/server .

FROM alpine

WORKDIR /app

COPY --from=builder /build/server server

CMD ["./server"]
