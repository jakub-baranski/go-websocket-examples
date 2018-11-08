FROM golang:1.10.4 as builder
# install dep
RUN go get github.com/golang/dep/cmd/dep

# Live reload utility for gin
RUN go get github.com/codegangsta/gin

WORKDIR /go/src/app

ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

ADD src src

CMD ["go", "run", "src/main.go"]
CMD ["gin", "--path", "src", "--port", "8080", "run", "main.go"]


