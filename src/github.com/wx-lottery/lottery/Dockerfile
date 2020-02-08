FROM golang:1.13-alpine AS build-env

ADD . /go/src/lottery
WORKDIR /go/src/lottery

ENV APP_NAME="lottery"
ENV GOPROXY="https://goproxy.cn"
RUN go build -o ${APP_NAME}

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN rm -rf /etc/localtime \
 && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENV APP_NAME="lottery"
ENV APP_ROOT="/data/apps/"${APP_NAME}
RUN mkdir -p $APP_ROOT
WORKDIR $APP_ROOT
COPY --from=build-env /go/src/${APP_NAME} $APP_ROOT/

EXPOSE 8080

CMD $APP_ROOT/$APP_NAME
