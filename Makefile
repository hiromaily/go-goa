# Note: tabs by space can't not used for Makefile!
MONGO_PORT=27017


update:
	go get -u github.com/goadesign/goa/...
	go get -u github.com/golang/dep/...

dep:
	dep init

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
	goagen bootstrap -d github.com/hiromaily/go-goa/design

genn:
	rm -f hy_*.go {public,swagger,health,js}.go
	goagen bootstrap -d github.com/hiromaily/go-goa/design

genfull:
	rm -rf app/ client/ swagger/ tool/
	rm -f hy_*.go {main,public,swagger,health,js}.go
	goagen bootstrap -d github.com/hiromaily/go-goa/design

run:
	go run *.go

bld:
	go build -i -v -o ${GOPATH}/bin/go-goa .

exec:
	go-goa
