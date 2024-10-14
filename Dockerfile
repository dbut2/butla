FROM golang:alpine AS builder

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY config.yaml config.yaml
RUN go build -o /bin/butla main.go

FROM alpine AS final

COPY --from=builder /bin/butla butla

EXPOSE 8080

CMD ["./butla"]
