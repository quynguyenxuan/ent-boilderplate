  
FROM golang:1.16.6-alpine3.14

RUN apk add build-base
ENV GOPATH /go
ENV GO111MODULE on

RUN wget -O /install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh

RUN chmod +x /install.sh
RUN /install.sh -b $(go env GOPATH)/bin

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENTRYPOINT ["/go/bin/air"]