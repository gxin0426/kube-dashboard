FROM golang:latest
MAINTAINER "gxin0426@126.com"
RUN go build .
EXPOSE 8877
ENTRYPOINT ["./dashboard"]