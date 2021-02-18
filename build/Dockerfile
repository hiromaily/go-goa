FROM hirokiy/goplus:1.11.5

#RUN echo $GOPATH => /go

#ARG redisHostName=default-redis-server
#ARG mysqlHostName=default-mysql-server

#RUN apt-get -y update && apt-get install -y git


#RUN go get -u github.com/goadesign/goa/... && \
#go get -u github.com/hiromaily/fresh

RUN go get -u github.com/hiromaily/go-goa/...


#submodule for swaggger-ui
WORKDIR /go/src/github.com/hiromaily/go-goa/resources/swagger-ui
RUN git submodule init && git submodule update

WORKDIR /go/src/github.com/hiromaily/go-goa/resources/swagger-ui/dist
RUN sed -e "s|http://petstore.swagger.io/v2/swagger.json|/swagger.json|g" index.html > goa.html

RUN mkdir -p /go/src/github.com/hiromaily/go-goa/tmp/log && mkdir -p /var/log/go
#RUN mkdir -p /go/src/github.com/hiromaily/go-goa/ext && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/goa && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/public && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/resources

WORKDIR /go/src/github.com/hiromaily/go-goa
# go get can retrieve all data from git repo
#COPY ./ext ./
#COPY ./goa ./
#COPY ./public ./
#COPY ./resources ./
#COPY ./Makefile ./
#COPY ./runner.conf ./

# Swagger -> it doesn't work...
#RUN ln -s /go/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger
#RUN ln -s /go/src/github.com/hiromaily/go-goa/goa/swagger /go/src/github.com/hiromaily/go-goa/public/swagger

# Docker use Godeps because these are for TravisCI
#RUN	rm -rf Godeps ./vendor
#RUN go get -d -v -u ./ext/cmd/
RUN	rm -rf frontend_workspace


RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/go-goa ./ext/cmd/main.go

#EXPOSE 80
CMD ["/go/bin/go-goa", "-f", "/go/src/github.com/hiromaily/go-goa/resources/tomls/docker.toml"]
