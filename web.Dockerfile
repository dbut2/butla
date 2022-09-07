FROM golang:1.18-alpine

WORKDIR /app

COPY ./vendor ./vendor
COPY ./config ./config
COPY ./internal ./internal
COPY ./go.mod go.mod
COPY ./go.sum go.sum
COPY ./main.go ./main.go

RUN go build -o web .

ENV PORT=8080

EXPOSE ${PORT}

ENV ENV=local

CMD ./web --config-path config/${ENV}.yaml
