FROM golang:1.17-alpine3.14 AS builder
COPY ./app/* /app/
WORKDIR /app
RUN go build \
    -ldflags "-s -w -X main.BuildAt=`date +%FT%T%z`" \
    -o redis-analysis

FROM alpine:3.14
COPY --from=builder /app/redis-analysis /usr/local/sbin/redis-analysis

