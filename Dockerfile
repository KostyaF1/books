FROM golang:latest

RUN mkdir -p /go/src/books
WORKDIR /go/src/books
COPY . .

ENV GOPATH=/go                              \
    PATH=/go/bin:$PATH
#    #PGDATA=/usr/local/var/postgres pg_ctl=reload \
    DRIVER=postgres                              \
    DBURL="postgresql://postgres:1111@localhost:5432/books?sslmode=disable"

RUN go get -v ./... && go build ./server/main.go


CMD ./main

EXPOSE 8080

