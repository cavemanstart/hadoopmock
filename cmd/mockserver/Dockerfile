FROM aslan-spock-register.qiniu.io/qa/pcdn-base:1017110257 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ARG GOC_BUILD
ARG NAMESPACE_BUILD

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN GOPROXY=https://goproxy.cn,direct go mod download
COPY . .

FROM aslan-spock-register.qiniu.io/qa/alpine-node-wkhtmltopdf:18.12.1-0.12.6-full

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/hadoopmock /app/hadoopmock

CMD ["./hadoopmock", "-f", "etc/config.yaml"]