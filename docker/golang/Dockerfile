FROM golang:1.10

#RUN echo $GOPATH => /go

#ARG redisHostName=default-redis-server
#ARG mysqlHostName=default-mysql-server

RUN apt-get -y update && apt-get install -y git

RUN go get -u github.com/goadesign/goa/... && \
go get -u github.com/hiromaily/fresh
