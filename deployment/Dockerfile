ARG BASE_IMAGE=dbut2/shortener
FROM ${BASE_IMAGE}

WORKDIR /app

ARG ENV=local
COPY ./configs/${ENV}.yaml ./config.yaml

CMD ["./lengthener"]
