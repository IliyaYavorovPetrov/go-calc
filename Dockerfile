FROM golang:1.21-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN apk update \
    && apk add --no-cache make

RUN make build

EXPOSE 8080

CMD ["make", "run"]
