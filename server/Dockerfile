FROM golang:latest

RUN mkdir /go-server

WORKDIR /go-server

COPY . /go-server/

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "."]