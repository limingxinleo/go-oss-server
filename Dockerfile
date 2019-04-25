FROM golang:1.12-alpine

LABEL maintainer="limx <715557344@qq.com>"
ENV GO111MODULE=on

RUN apk add git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/limingxinleo/go-oss-server

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN go build

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["./go-oss-server"]