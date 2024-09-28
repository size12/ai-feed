FROM golang:alpine

ADD . /build

WORKDIR /build

RUN go build -o service cmd/feed/main.go
CMD ["./service"]