FROM golang:1.18-alpine

ENV ENV=local
ENV PORT=8081

EXPOSE ${PORT}

WORKDIR /app

COPY ./vendor ./vendor
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod go.mod
COPY ./go.sum go.sum

RUN go build -o admin cmd/admin/admin.go

CMD ./admin
