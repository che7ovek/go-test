# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o botApp ./cmd

RUN chmod +x /app/botApp

#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/botApp /app

CMD [ "/app/botApp" ]