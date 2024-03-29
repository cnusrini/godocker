FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o webserver .

CMD ["/app/webserver"]
