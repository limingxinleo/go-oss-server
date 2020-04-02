FROM golang:1.14 as builder

LABEL maintainer="limx <715557344@qq.com>"
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE=on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/src/github.com/limingxinleo/go-oss-server

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/limingxinleo/go-oss-server/app /

EXPOSE 8080

ENTRYPOINT ["/app"]
