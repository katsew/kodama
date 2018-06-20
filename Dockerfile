FROM golang:alpine

RUN apk --update add git openssh ca-certificates iptables && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY . /go/src/github.com/katsew/kodama

WORKDIR /go/src/github.com/katsew/kodama

RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install

ENTRYPOINT ["kodama"]
CMD ["http", "server"]