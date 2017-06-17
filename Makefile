# Note: tabs by space can't not used for Makefile!
MONGO_PORT=27017

###############################################################################
# Initialization
###############################################################################
init:
	mkdir -p public goa ext/cmd
	touch public/index.html
	#goagen	goagen bootstrap
	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger


###############################################################################
# Docker
###############################################################################
dc:
	docker-compose up


###############################################################################
# PKG Dependencies
###############################################################################
update:
	go get -u github.com/goadesign/goa/...
	go get -u github.com/golang/dep/...
	go get -u github.com/rakyll/hey
	go get -u github.com/davecheney/httpstat
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/gordonklaus/ineffassign
    go get -u github.com/pilu/fresh
	go get -u github.com/alecthomas/gometalinter
	#gometalinter --install

	# this doesn't work
	#go get -u ./ext/cmd/

# dep is dependencies tools
depinit:
	cd ext/cmd/;dep init

dep:
	cd ext/cmd/;dep ensure -update

depcln:
	cd ext/cmd/;rm -rf vendor lock.json manifest.json


###############################################################################
# Golang formatter and detection
###############################################################################
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


###############################################################################
# Goa generation
###############################################################################
gen:
	#goagen won’t be re-generated (by default) if already present
	goagen bootstrap -d github.com/hiromaily/go-goa/goa/design -o goa/

gencln:
	rm -f goa/*.go
	#rm -f goa/hy_*.go goa/{public,swagger,health}.go
	rm -rf goa/app/ goa/client/ goa/swagger/ goa/tool/
	goagen bootstrap -d github.com/hiromaily/go-goa/goa/design -o goa/

aftergen:
	# rewrite package name
	rm -f goa/main.go
	sed -e "1s/main/goa/" ./goa/auth.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/auth.go
	sed -e "1s/main/goa/" ./goa/health.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/health.go
	sed -e "1s/main/goa/" ./goa/public.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/public.go
	sed -e "1s/main/goa/" ./goa/hy_user.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/hy_user.go
	sed -e "1s/main/goa/" ./goa/hy_company.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/hy_company.go

genfull: gencln aftergen

genctl:
	goagen controller -d github.com/hiromaily/go-goa/goa/design -o goa/


###############################################################################
# Build
###############################################################################
run:
	#go run goa/*.go
	go run ext/cmd/*.go

bld:
	#go build -i -v -o ${GOPATH}/bin/go-goa ./goa/
	go build -i -v -o ${GOPATH}/bin/go-goa ./ext/cmd/

bldlinux:
	#GOOS=linux GOARCH=amd64 go build -v -o ${GOPATH}/bin/linux_amd64/$1 ./goa/
	GOOS=linux GOARCH=amd64 go build -v -o ${GOPATH}/bin/linux_amd64/$1 ./ext/cmd/

clibld:
	go build -i -v -o ${GOPATH}/bin/go-goa-cli ./goa/tool/api-cli/*.go


###############################################################################
# Execution
###############################################################################
exec:
	go-goa

cln:
	go clean -n

clnok:
	go clean

cli:
	go-goa-cli company-list hy-company


###############################################################################
# httpie
###############################################################################
http:
	# httpie #brew install httpie
	# jq     #brew install jq

	# Static files
	http localhost:8080/
	http localhost:8080/swagger/swagger.json
	http localhost:8080/swagger-ui/
	http localhost:8080/api/_ah/health

	# Login
	#http POST http://localhost:8080/api/auth/login username=hiro password=xxxxxxxx
	http --body POST http://localhost:8080/api/auth/login username=hiro password=xxxxxxxx

	$(eval TOKEN := $(shell http --body POST http://localhost:8080/api/auth/login username=hiro password=xxxxxxxx | jq '.token' | sed 's/"//g'))

	# User
	http localhost:8080/api/user 'Authorization: Bearer $(TOKEN)'
	http localhost:8080/api/user/1 'Authorization: Bearer $(TOKEN)'
	http POST http://localhost:8080/api/user name=Harry email=test@oo.bb 'Authorization: Bearer $(TOKEN)'
	http PUT http://localhost:8080/api/user/1 name=Harry email=test@oo.bb 'Authorization: Bearer $(TOKEN)'
	http DELETE http://localhost:8080/api/user/1 'Authorization: Bearer $(TOKEN)'

	# Company
	http localhost:8080/api/company 'Authorization: Bearer $(TOKEN)'
	http localhost:8080/api/company/1 'Authorization: Bearer $(TOKEN)'
	http POST http://localhost:8080/api/company name=Google country=America address=California 'Authorization: Bearer $(TOKEN)'
	http PUT http://localhost:8080/api/company/1 name=Google country=America address=California 'Authorization: Bearer $(TOKEN)'
	http DELETE http://localhost:8080/api/company/1 'Authorization: Bearer $(TOKEN)'


###############################################################################
# Curl
###############################################################################
curl:
	# curl
	# Static files
	curl -i localhost:8080/
	curl -i localhost:8080/swagger/swagger.json
	curl -i localhost:8080/swagger-ui/

	# Health check
	curl -i localhost:8080/api/_ah/health

	# Login
	curl -i -H "Content-Type: application/x-www-form-urlencoded" -d "username=hiro&password=xxxxxxxx" -X POST http://localhost:8080/api/auth/login

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


###############################################################################
# Bench
###############################################################################
bench:
	hey -n 20000 -c 50 -m GET http://localhost:8080/api/user


###############################################################################
# HTTP Stat
###############################################################################
httpstat:
	httpstat http://localhost:8080/api/user
