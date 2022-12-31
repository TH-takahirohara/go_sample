FROM golang:1.19.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /docker_mysql_sample

EXPOSE 8080

CMD [ "/docker_mysql_sample" ]
