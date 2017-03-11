# Note: tabs by space can't not used for Makefile!
MONGO_PORT=27017

init:
	mkdir public
	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/swagger ./public/swagger
	touch public/index.html

update:
	go get -u github.com/goadesign/goa/...
	go get -u github.com/golang/dep/...

dep:
	#dep init
	dep ensure -update

fmt:
	go fmt `go list ./... | grep -v '/vendor/'`

vet:
	go vet `go list ./... | grep -v '/vendor/'`

fix:
	go fix `go list ./... | grep -v '/vendor/'`

lint:
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .

chk:
	go fmt `go list ./... | grep -v '/vendor/'`
	go vet `go list ./... | grep -v '/vendor/'`
	go fix `go list ./... | grep -v '/vendor/'`
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .

gen:
	#goagen won’t be re-generated (by default) if already present
	goagen bootstrap -d github.com/hiromaily/go-goa/design

genn:
	rm -f hy_*.go {public,swagger,health}.go
	goagen bootstrap -d github.com/hiromaily/go-goa/design

genfull:
	rm -rf app/ client/ swagger/ tool/
	rm -f hy_*.go {main,public,swagger,health}.go
	goagen bootstrap -d github.com/hiromaily/go-goa/design

run:
	go run *.go

bld:
	go build -i -v -o ${GOPATH}/bin/go-goa .

clibld:
	go build -i -v -o ${GOPATH}/bin/go-goa-cli ./tool/api-cli/*.go

exec:
	go-goa

cli:
	go-goa-cli company-list hy-company

curlall:
	curl -i localhost:8080/
	curl -i localhost:8080/swagger.json
	curl -i localhost:8080/api/_ah/health
	curl -i localhost:8080/api/company
	curl -i localhost:8080/api/company/1
	curl -i localhost:8080/api/user
	curl -i localhost:8080/api/user/1
