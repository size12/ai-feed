FROM golang:alpine

ADD ../ai-feed2 .

RUN go build -o service cmd/feed/main.go
CMD ["./service"]