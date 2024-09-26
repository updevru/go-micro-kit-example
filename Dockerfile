## Контейнер для сборки Go
FROM golang:1.21  AS builder
RUN mkdir -p /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main

## Контейнер в котором будет находиться программа
FROM alpine:latest
LABEL maintainer="Ladygin Sergey <sladygin@updev.ru>"

EXPOSE 8080
EXPOSE 8081

RUN apk update && \
    apk add --no-cache tzdata

COPY --from=builder /app/main /app/main

WORKDIR /app
ENTRYPOINT ["./main"]