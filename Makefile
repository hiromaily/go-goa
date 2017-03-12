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

bldlinux:
	GOOS=linux GOARCH=amd64 go build -v -o ${GOPATH}/bin/linux_amd64/$1 .

clibld:
	go build -i -v -o ${GOPATH}/bin/go-goa-cli ./tool/api-cli/*.go

exec:
	go-goa

cln:
	go clean -n

clnok:
	go clean

cli:
	go-goa-cli company-list hy-company

curlall:
	# curl
	# Static files
	curl -i localhost:8080/
	curl -i localhost:8080/swagger/swagger.json

	# Health check
	curl -i localhost:8080/api/_ah/health

	# User
	curl -i localhost:8080/api/user
	curl -i localhost:8080/api/user/1

	curl -i -H "Content-Type: application/x-www-form-urlencoded" -d "name=Harry&email=aa@bb.cc" -X POST http://localhost:8080/api/user
	curl -i --data-urlencode "name=harry" --data-urlencode "email=aa@bb.cc" http://localhost:8080/api/user
	curl -i -H "Content-Type: application/json" -d '{"name":"harry","email":"aa@bb.cc"}' -X POST http://localhost:8080/api/user

	curl -i -H "Content-Type: application/x-www-form-urlencoded" -d "name=harry&email=aa@bb.cc" -X PUT http://localhost:8080/api/user/1
	curl -i -X DELETE http://localhost:8080/api/user/1

	# Company
	curl -i localhost:8080/api/company
	curl -i localhost:8080/api/company/1

	curl -i -H "Content-Type: application/x-www-form-urlencoded" -d "name=Google&country=America&address=California" -X POST http://localhost:8080/api/company
	curl -i -H "Content-Type: application/x-www-form-urlencoded" -d "name=Google&country=America&address=California" -X PUT http://localhost:8080/api/company/1
	curl -i -X DELETE http://localhost:8080/api/company/1

httpieall:
    # httpie
	# Static files
	http localhost:8080/
	http localhost:8080/swagger/swagger.json
	http localhost:8080/api/_ah/health

    # User
	http localhost:8080/api/user
	http localhost:8080/api/user/1
    http POST http://localhost:8080/api/user name=Harry email=test@oo.bb
    http PUT http://localhost:8080/api/user/1 name=Harry email=test@oo.bb
    http DELETE http://localhost:8080/api/user/1

    # Company
	http localhost:8080/api/company
	http localhost:8080/api/company/1
    http POST http://localhost:8080/api/company name=Google country=America address=California
    http PUT http://localhost:8080/api/company/1 name=Google country=America address=California
    http DELETE http://localhost:8080/api/comany/1
