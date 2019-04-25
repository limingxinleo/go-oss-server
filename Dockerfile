FROM golang:1.12-alpine as builder

LABEL maintainer="limx <715557344@qq.com>"
ENV GO111MODULE=on

RUN apk add git

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/src/github.com/limingxinleo/go-oss-server

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /go/src/github.com/limingxinleo/go-oss-server /

EXPOSE 8080

ENTRYPOINT ["/app"]
