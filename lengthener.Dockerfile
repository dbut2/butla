FROM golang:1.18-alpine

ENV ENV=local
ENV PORT=8080

EXPOSE ${PORT}

WORKDIR /app

COPY ./vendor ./vendor
COPY ./configs ./configs
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod go.mod
COPY ./go.sum go.sum
COPY ./cmd/lengthener/lengthener.go ./cmd/lengthener/lengthener.go

RUN go build -o lengthener cmd/lengthener/lengthener.go

CMD ./lengthener
