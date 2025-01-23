FROM golang:alpine AS builder

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
RUN go build -o /bin/butla main.go

FROM alpine AS final

WORKDIR /app

COPY --from=builder /bin/butla /bin/butla

EXPOSE 8080

CMD ["/bin/butla"]
