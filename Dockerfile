FROM golang:1.8

#ARG redisHostName=default-redis-server
#ARG mysqlHostName=default-mysql-server

RUN mkdir -p /go/src/github.com/hiromaily/go-goa/ext && \
 mkdir -p /go/src/github.com/hiromaily/go-goa/goa && \
 mkdir -p /go/src/github.com/hiromaily/go-goa/public && \
 mkdir -p /go/src/github.com/hiromaily/go-goa/resources

WORKDIR /go/src/github.com/hiromaily/go-goa
COPY ./ext ./
COPY ./goa ./
COPY ./public ./
COPY ./resources ./
COPY ./Makefile ./

RUN go get -u github.com/goadesign/goa/... && \
go get -u github.com/pilu/fresh
#go get -u github.com/nats-io/nats

RUN go get -d -v ./ext/cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/goa ./ext/cmd/main.go

EXPOSE 8080
CMD ["/go/bin/goa"]
