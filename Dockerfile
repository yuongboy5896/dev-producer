FROM golang:alpine as builder

MAINTAINER lwnmengjing

ENV GOPROXY https://goproxy.cn/

WORKDIR /go/release
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata

COPY go.mod ./go.mod
RUN go mod tidy
COPY . .
RUN pwd && ls

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o dev-producer .

FROM alpine

COPY --from=builder /go/release/dev-producer /
COPY  --from=builder /go/release/config /config
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8090

CMD ["/dev-producer"]