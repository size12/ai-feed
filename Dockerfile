FROM golang:alpine

ADD . .

RUN go build -o service cmd/feed/main.go
CMD ["./service"]