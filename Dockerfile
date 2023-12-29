# Build stage
FROM golang:1.18-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN apk update \
    && apk add --no-cache make

RUN make build

# Run stage
FROM alpine:3.17

WORKDIR /app

COPY --from=builder app/bin/main /app/bin/main

EXPOSE 8080

CMD ["./main"]