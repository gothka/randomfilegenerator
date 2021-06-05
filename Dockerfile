FROM golang:1.16.5-alpine

RUN mkdir /app

ADD main.go /app

WORKDIR /app

RUN go env -w GO111MODULE=off

RUN go build -o main .

CMD ["/app/main"]