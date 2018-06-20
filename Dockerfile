FROM golang:alpine AS build-machine

RUN apk --update add git openssh ca-certificates iptables && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY . /go/src/github.com/katsew/kodama

WORKDIR /go/src/github.com/katsew/kodama

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install

ENTRYPOINT ["/go/bin/kodama"]
CMD ["http", "server"]