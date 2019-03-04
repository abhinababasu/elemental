FROM golang:latest

MAINTAINER abhinaba@gmail.com

ENV GOPATH=/go

RUN apt-get install -y git && \
    go get github.com/bonggeek/element && \
    go get github.com/bonggeek/wordelement

EXPOSE 8080

WORKDIR /go/src/github.com/bonggeek/wordelement
RUN go build wordelementservice.go

CMD ["/go/src/github.com/bonggeek/wordelement/wordelementservice"]

