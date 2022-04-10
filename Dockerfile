FROM golang:1.17.0-alpine3.13 AS builder
RUN go env -w GO111MODULE=auto \
  && go env -w GOPROXY=https://goproxy.cn,direct 
WORKDIR /build
COPY ./ .
RUN cd /build && go build -tags netgo -ldflags="-w -s" -o onebotplus cmd/main.go 

FROM golang:1.17.0-alpine3.13
LABEL MAINTAINER=github.com/dezhiShen
WORKDIR /data
RUN apk add -U --repository http://mirrors.ustc.edu.cn/alpine/v3.13/main/ tzdata 
COPY --from=builder /build/onebotplus /usr/bin/onebotplus 
RUN chmod +x /usr/bin/onebotplus
VOLUME /data
ENTRYPOINT ["/usr/bin/onebotplus"]