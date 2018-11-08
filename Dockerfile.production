FROM golang:1.10.4 as builder
# install dep
RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/app

ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

ADD src src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o websocket_server src/*


FROM alpine:3.7

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root

COPY --from=builder /go/src/app/websocket_server .

CMD ["./websocket_server"]


