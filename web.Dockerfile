FROM golang:1.18-alpine

ENV ENV=local
ENV PORT=8080

EXPOSE ${PORT}

WORKDIR /app

COPY ./vendor ./vendor
COPY ./config ./config
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod go.mod
COPY ./go.sum go.sum
COPY ./cmd/web/web.go ./cmd/web/web.go

RUN go build -o web cmd/web/web.go

CMD ./web
