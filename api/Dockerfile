FROM golang:1.20.2
RUN mkdir -p /go/src/cozyLiv

WORKDIR /go/src/cozyLiv/api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD . /go/src/cozyLiv

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]