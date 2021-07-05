FROM golang:1.16 AS builder

RUN mkdir -p /app
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o /web-shortener ./cmd/server/main.go

### we are using another container to use for application
FROM scratch
#RUN mkdir -p /app
COPY --from=builder /app /app
WORKDIR /app
CMD ["/web-shortener","-fileConfig=./config.yaml"]