

FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o check-saldo-server ./cmd/check-saldo-server

FROM alpine

WORKDIR /app

COPY --from=builder /app/check-saldo-server /usr/bin/

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY .env .

EXPOSE 8005

CMD /wait && check-saldo-server