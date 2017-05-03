FROM golang:alpine

ADD src/main.go /tmp

RUN apk add --no-cache git bzr rpm xzi glide && \
    go build -o /ferrarin /tmp/main.go && \
    rm -rf /go /usr/local/go /tmp/main.go

ENTRYPOINT ["/ferrarin"]

