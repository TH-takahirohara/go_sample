# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

# Go modulesのダウンロード
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# コンパイルを行い、成果物であるバイナリファイルをルートに置く
RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
