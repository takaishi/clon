FROM golang:1.11.1 as build

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
ENV CGO_ENABLED 0

RUN mkdir -p /go/{src,bin,pkg}

ADD . /go/src/github.com/takaishi/clon
WORKDIR /go/src/github.com/takaishi/clon
RUN go get
RUN go build

FROM alpine:latest as app
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /
COPY --from=build /go/src/github.com/takaishi/clon/clon /clon

ENTRYPOINT ["/clon", "--config", "/etc/clon.yml"]