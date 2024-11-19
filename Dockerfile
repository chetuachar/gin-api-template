FROM golang:1.22 AS builder

WORKDIR /gin-api-template

RUN mkdir src
COPY main.go .
COPY ./src ./src


RUN go mod init gin-api-template


RUN go mod tidy


RUN go build -o gin-api-template .



FROM debian:stable-slim


LABEL description=" " \
    version="1.0"


ENV DEBIAN_FRONTEND=noninteractive \
    TZ=Asia/Kolkata

RUN apt-get update \
    && apt-get install -y procps ca-certificates


WORKDIR /gin-api-template

RUN mkdir log


COPY --from=builder /gin-api-template/gin-api-template .


EXPOSE 8086

HEALTHCHECK --interval=30s --timeout=10s --retries=2 \
  CMD ps aux | grep '[g]in-api-template'


CMD ["./gin-api-template"]