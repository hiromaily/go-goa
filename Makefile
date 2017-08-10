# Note: tabs by space can't not used for Makefile!

PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/go-goa
TOMLPATH=${PROJECT_ROOT}/resources/tomls/settings.toml
TOMLPATH2=${PROJECT_ROOT}/resources/tomls/docker.toml

###############################################################################
# Initialization
###############################################################################
init_local:
	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger
	git submodule add https://github.com/swagger-api/swagger-ui.git resources/swagger-ui
	#after this, `make genfull`

###############################################################################
# Docker
###############################################################################
dcgobld:
	#goplus:1.8
	docker build -t hirokiy/goplus:1.8 -f ./docker/golang/Dockerfile .
	docker push hirokiy/goplus:1.8

dcup:
	docker-compose build
	docker-compose up

dcbld:
	docker-compose build --no-cache

dcins:
	docker exec -it gogoa_webserver_1 bash

dctest:
	docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go
	#docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go -f /go/src/github.com/hiromaily/go-goa/resources/tomls/docker.toml"

dcpush:
	docker push hirokiy/go-goa:1.0

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
	#go get -d -v ./ext/cmd/

# dep is dependencies tools
depinit:
	cd ext/cmd/;dep init

dep:
	cd ext/cmd/;dep ensure -update

depcln:
	cd ext/cmd/;rm -rf vendor lock.json manifest.json

#godep:
#	rm -rf Godeps
#	rm -rf ./vendor
#	godep save ./...


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
	golint ./... | grep -v '^goa\/' || true
	#golint ./... | grep -v '^vendor\/' || true

chk:
	go fmt `go list ./... | grep -v '/vendor/'`
	go vet `go list ./... | grep -v '/vendor/'`
	go fix `go list ./... | grep -v '/vendor/'`
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .


###############################################################################
# Goa generation (It's better to exexute `make genfull` regularly
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
	sed -e "1s/main/goa/" ./goa/hy_companybranch.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/hy_companybranch.go
	sed -e "1s/main/goa/" ./goa/hy_usertech.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/hy_usertech.go
	sed -e "1s/main/goa/" ./goa/hy_user_work_history.go >> ./resources/tmp/tmp.go
	mv -f ./resources/tmp/tmp.go ./goa/hy_user_work_history.go

genfull: gencln aftergen

genctl:
	goagen controller -d github.com/hiromaily/go-goa/goa/design -o goa/


updgoa:
	go get -u github.com/goadesign/goa/...

#link_swagger:
#	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger

# this command should be executed regularly
updall: updgoa gencln aftergen


###############################################################################
# Build for local
###############################################################################
run:
	go run ext/cmd/main.go
	#go run goa/*.go

bld:
	go build -i -v -o ${GOPATH}/bin/go-goa ./ext/cmd/
	#go build -i -v -o ${GOPATH}/bin/go-goa ./goa/

bldlinux:
	GOOS=linux GOARCH=amd64 go build -v -o ${GOPATH}/bin/linux_amd64/$1 ./ext/cmd/
	#GOOS=linux GOARCH=amd64 go build -v -o ${GOPATH}/bin/linux_amd64/$1 ./goa/

clibld:
	go build -i -v -o ${GOPATH}/bin/go-goa-cli ./goa/tool/api-cli/*.go


###############################################################################
# Execution for local
###############################################################################
exec:
	go-goa


###############################################################################
# CLI
###############################################################################
cli:
	go-goa-cli company-list hy-company


###############################################################################
# maintenance
###############################################################################
cln:
	go clean -n

clnok:
	go clean


###############################################################################
# test
###############################################################################
gotest:
	go test -v ext/cmd/*.go
	#go test -v ext/cmd/*.go -f ${TOMLPATH}
	#go test -v ext/cmd/main_test.go -health 10

gotest1:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestUserAPIOnTable"

gotest2:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestCompanyAPIOnTable"

gotest3:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestCompanyBranchAPIOnTable"

gotest4:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestGetUserTechOnTable"

gotest5:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestGetUserWorkHistoryOnTable"


###############################################################################
# Heroku
###############################################################################
heroku_init:
	heroku plugins:install heroku-container-registry
	heroku container:login
	heroku create goa-web
	heroku container:push web
	cd docker/mysql;heroku container:push mysql


heroku_after_change_app_name:
    #git remote remove heroku [Important!!]
	#registry.heroku.com/go-goa/web
    docker build -t hirokiy/go-goa:1.0 -f ./Dockerfile .
	docker build -t hirokiy/goa-mysql:1.0 -f ./docker/mysql/Dockerfile .

	#docker tag hirokiy/go-goa:1.0 registry.heroku.com/goa-web/web
	#docker tag registry.heroku.com/safe-inlet-49884/mysql registry.heroku.com/goa-web/mysql

	docker push registry.heroku.com/goa-web/web
	docker push registry.heroku.com/goa-web/mysql

heroku_settings:
	heroku apps
	#goa-web

heroku_open:
	open -a goa-web
	#https://goa-web.herokuapp.com/

heroku_remove:
    heroku apps:destroy

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
	#http POST http://localhost:8080/api/auth/login email=aaaa@test.jp password=password
	http --body POST http://localhost:8080/api/auth/login email=aaaa@test.jp password=password

	$(eval TOKEN := $(shell http --body POST http://localhost:8080/api/auth/login email=aaaa@test.jp password=password | jq '.token' | sed 's/"//g'))

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
