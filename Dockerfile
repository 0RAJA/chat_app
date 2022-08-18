FROM golang:alpine AS builder

# 移动到工作目录 /app
WORKDIR /app

#为镜像设置环境变量
ENV Go111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go build -ldflags="-s -w" -o bin/chat src/main.go
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add curl
RUN curl -L https://ghproxy.com/https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/chat .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/src/dao/postgres/migration ./migration
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/config/app ./config

RUN chmod +x wait-for.sh
RUN chmod +x start.sh

EXPOSE 8080

#ENTRYPOINT ["/app/start.sh"]
#CMD ["/app/chat","-path=/app/config"]
