
PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/go-goa
TOMLPATH=${PROJECT_ROOT}/configs/settings.toml
TOMLPATH2=${PROJECT_ROOT}/configs/docker.toml

###############################################################################
# Initialization
###############################################################################
.PHONY: setup-tools
setup-tools:
	go get -u github.com/rakyll/hey
	go get -u github.com/davecheney/httpstat
	go get -u goa.design/goa/v3/...@v3
	go get -u goa.design/plugins/v3/cors/dsl

.PHONY: setup-resume-service
setup-resume-service:
	cd internal/goa/service/resume && \
	go mod init resume && \
	go get -u goa.design/goa/v3/...@v3 && \
	go get -u goa.design/plugins/v3/cors/dsl

#.PHONY: setup-swagger
#setup-swagger:
#	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger
#	git submodule add https://github.com/swagger-api/swagger-ui.git resources/swagger-ui
#	#after this, `make genfull`
#
#	#
#	cd resources/swagger-ui/dist/;sed -e "s|http://petstore.swagger.io/v2/swagger.json|/swagger.json|g" index.html > goa.html

###############################################################################
# PKG Dependencies
###############################################################################
.PHONY: update-service
update-service:
	cd pkg/goa/design/resume && \
	go get -u -d -v ./...


###############################################################################
# linter and formatter
###############################################################################
.PHONY: lint-all
lint-all: imports lint

.PHONY: imports
imports:
	./scripts/imports.sh

.PHONY: lint
lint:
	golangci-lint run --fix


###############################################################################
# Goa generation
###############################################################################
.PHONY: gen-design
gen-design:
	cd internal/goa/service/resume && goa gen resume/design

.PHONY: gen-example
gen-example:
	cd internal/goa/service/resume && goa example resume/design

.PHONY: replace-example
replace-example:
	grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'
	grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server-cli/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'
	#resumeapi "github.com/hiromaily/go-goa/pkg/goa/service/resume"
	#resumeapi "resume"

.PHONY: move-example
move-example:
	mv internal/goa/service/resume/cmd/resume_api_server/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/server/
	mv internal/goa/service/resume/cmd/resume_api_server-cli/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/cli/
	mv internal/goa/service/resume/auth.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/health.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_company.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_companybranch.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_tech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_user.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_user_work_history.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/hy_usertech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	mv internal/goa/service/resume/public.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
	rm -rf internal/goa/service/resume/cmd

.PHONY: gen-all
gen-all: gen-design gen-example replace-example move-example


###############################################################################
# Build on local
###############################################################################
run-server:
	go run -race ./cmd/resume-api/server/...

bld-server:
	go build -race -v -o ${GOPATH}/bin/goa-server ./cmd/resume-api/server/...

#bldlinux:
#	GOOS=linux GOARCH=amd64 go build -race -v -o ${GOPATH}/bin/linux_amd64/$1 ./ext/cmd/


###############################################################################
# Goa generation (It's better to exexute `make genfull` regularly (old style)
###############################################################################
#gen-old:
#	#goagen won’t be re-generated (by default) if already present
#	goagen bootstrap -d github.com/hiromaily/go-goa/goa/design -o goa/
#
#gencln-old:
#	rm -f goa/*.go
#	#rm -f goa/hy_*.go goa/{public,swagger,health}.go
#	rm -rf goa/app/ goa/client/ goa/swagger/ goa/tool/
#	goagen bootstrap -d github.com/hiromaily/go-goa/goa/design -o goa/
#
#aftergen-old:
#	# rewrite package name
#	rm -f goa/main.go
#
#	sed -e "1s/main/goa/" ./goa/auth.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/auth.go
#
#	sed -e "1s/main/goa/" ./goa/health.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/health.go
#
#	sed -e "1s/main/goa/" ./goa/public.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/public.go
#
#	sed -e "1s/main/goa/" ./goa/hy_user.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_user.go
#
#	sed -e "1s/main/goa/" ./goa/hy_tech.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_tech.go
#
#	sed -e "1s/main/goa/" ./goa/hy_company.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_company.go
#
#	sed -e "1s/main/goa/" ./goa/hy_companybranch.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_companybranch.go
#
#	sed -e "1s/main/goa/" ./goa/hy_usertech.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_usertech.go
#
#	sed -e "1s/main/goa/" ./goa/hy_user_work_history.go >> ./resources/tmp/tmp.go
#	mv -f ./resources/tmp/tmp.go ./goa/hy_user_work_history.go
#
#genfull: gencln aftergen
#
#genctl:
#	goagen controller -d github.com/hiromaily/go-goa/goa/design -o goa/

# this command should be executed regularly
#updall: updgoa gencln aftergen

#link_swagger:
#	ln -s ${GOPATH}/src/github.com/hiromaily/go-goa/goa/swagger ./public/swagger



###############################################################################
# Docker
###############################################################################
dcgobld:
	#goplus:1.11.5
	docker build -t hirokiy/goplus:1.11.5 -f ./docker/golang/Dockerfile .
	docker push hirokiy/goplus:1.11.5

dcup:
	docker-compose build
	docker-compose up

dcbld:
	docker-compose build

dcbldn:
	docker-compose build --no-cache

dcins:
	docker exec -it gogoa_webserver_1 bash

dctest:
	docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go"
	#docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go -f /go/src/github.com/hiromaily/go-goa/resources/tomls/docker.toml"

dcpush:
	docker push hirokiy/go-goa:1.0


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

gotest6:
	go test -v ext/cmd/*.go -run "TestLoginOnTable|TestTechAPIOnTable"


###############################################################################
# Heroku
###############################################################################
heroku_init:
	heroku plugins:install heroku-container-registry
	heroku container:login
	heroku create goa-web
	#heroku container:push web
	#cd docker/mysql;heroku container:push mysql

heroku_after_change_app_name:
	#git remote remove heroku [Important!!]

heroku_settings:
	heroku apps
	#goa-web

heroku_open:
	open -a goa-web
	#https://goa-web.herokuapp.com/

heroku_remove:
	heroku apps:destroy


###############################################################################
# Build Image for Heroku
###############################################################################
heroku_build_docker1:
	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
	docker push hirokiy/goapack_base:latest

heroku_build_docker2:
	docker build --no-cache -t hirokiy/goapack:latest -f ./docker/Dockerfile.heroku .

heroku_build_multistage:
	#It works!
	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.multistage.heroku .

heroku_bldfull:
	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
	docker build --no-cache -t hirokiy/goapack:latest -f ./docker/Dockerfile.heroku .

heroku_exec_docker:
	docker run -p 8080:8080 --name goapack hirokiy/goapack:latest
	docker stop goapack

heroku_upd:
	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.heroku .
	docker push registry.heroku.com/goa-web/web

heroku_updfull:
	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.heroku .
	docker push registry.heroku.com/goa-web/web


###############################################################################
# Deploy Heroku on Travis
###############################################################################
travis_init:
	gem install travis

travis_enc:
	travis encrypt $(heroku auth:token) --add deploy.api_key

###############################################################################
# Build Image for GCP Kubernetes
###############################################################################
dcp_build:
	docker-compose build --no-cache webserver
	#docker build --no-cache -t hirokiy/go-goa-mysql:latest -f ./docker/mysql/Dockerfile .

dcp_push:
	docker push hirokiy/go-goa:1.0
	#docker push hirokiy/go-goa-mysql:latest


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
	#http POST http://localhost:8080/api/auth/login email=hiroki@goa.com password=password
	http --body POST http://localhost:8080/api/auth/login email=hiroki@goa.com password=password

	$(eval TOKEN := $(shell http --body POST http://localhost:8080/api/auth/login email=hiroki@goa.com password=password | jq '.token' | sed 's/"//g'))

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
