FROM golang:latest
MAINTAINER "gxin0426@126.com"
WORKDIR /opt/kube-dashboard
ADD . /opt/kube-dashboard
RUN go build .
EXPOSE 8877
ENTRYPOINT ["./dashboard"]