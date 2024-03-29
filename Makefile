CURRENTDIR=`pwd`
modVer=$(shell cat go.mod | head -n 3 | tail -n 1 | awk '{print $2}' | cut -d'.' -f2)
currentVer=$(shell go version | awk '{print $3}' | sed -e "s/go//" | cut -d'.' -f2)

PROJECT_ROOT=${GOPATH}/src/github.com/hiromaily/go-goa
TOMLPATH=${PROJECT_ROOT}/configs/settings.toml

ENDPOINT=localhost:8090/api

###############################################################################
# Initialization
###############################################################################
# expects MacOS user here because of using `brew` command
.PHONY: setup-tools
setup-tools:
	brew install httpie
	brew install jq
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go install github.com/rakyll/hey@latest
	go install github.com/davecheney/httpstat@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	go install github.com/cosmtrek/air@latest
	#go get -u github.com/volatiletech/null/v8

.PHONY: sqlboiler
sqlboiler:
	sqlboiler --wipe mysql


###############################################################################
# linter and formatter
###############################################################################
.PHONY: check-ver
check-ver:
	#echo $(modVer)
	#echo $(currentVer)
	@if [ ${currentVer} -lt ${modVer} ]; then\
		echo go version ${modVer}++ is required but your go version is ${currentVer};\
	fi

.PHONY: imports
imports:
	./scripts/imports.sh

.PHONY: lint
lint:
	golangci-lint run

#.PHONY: lint-fix
#lint-fix: imports
#	golangci-lint run --fix

.PHONY: cache-clean
cache-clean:
	golangci-lint cache clean

###############################################################################
# Goa generation
# -
###############################################################################
.PHONY: gen-design
gen-design:
	make -C internal/goa/service/resume gen

.PHONY: gen-example
gen-example:
	make -C internal/goa/service/resume example


###############################################################################
# Build on local
###############################################################################

.PHONY: build
build:
	go build -v -o ${GOPATH}/bin/goa-file-server ./cmd/fileserver/server/...
	go build -v -o ${GOPATH}/bin/goa-server ./cmd/resume/server/...
	go build -v -o ${GOPATH}/bin/goa-client ./cmd/resume/cli/...

.PHONY: build-ci
build-ci:
	go build -v -o goa-file-server ./cmd/fileserver/server/...
	go build -v -o goa-server ./cmd/resume/server/...
	go build -v -o goa-client ./cmd/resume/cli/...

.PHONY: run
run:
	goa-server -conf ./configs/settings.toml
	#go run -race ./cmd/resume/server/... -conf ./configs/settings.toml

.PHONY: run-file
run-file:
	goa-file-server

.PHONY: rerun
rerun: build run


###############################################################################
# Docker
###############################################################################
docker-build:
	docker compose build --no-cache api-server
	docker compose build --no-cache file-server


###############################################################################
# TEST API
###############################################################################
.PHONY: test
test:
	./scripts/api-test.sh

###############################################################################
# check API by httpie
###############################################################################
.PHONY: chk
chk: http-login
	http $(ENDPOINT)/user/1/workhistory 'Authorization: Bearer $(TOKEN)'

# Login
# when displaying response body
#http --body POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password
#$(eval TOKEN := $(shell http --body POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password | jq '.token' | sed 's/"//g'))
# when displaying response headers
#http --headers POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password
#$(eval TOKEN := $(shell http --headers POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password | head -n 2 | tail -n 1 | sed -e "s/Authorization: //g"))

# FXIME: somehow this task run by running other task
.PHONY: http-login
http-login:
	#$(eval TOKEN := $(shell http --headers POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password | head -n 2 | tail -n 1 | sed -e "s/Authorization: //g"))
    $(eval TOKEN := $(shell http --body POST $(ENDPOINT)/auth/login email=hiroki@goa.com password=password | jq '.token' | sed 's/"//g'))

.PHONY: static
static:
	# Static files
	http localhost:8080/assets/index.html
	http localhost:8080/openapi.json
	http localhost:8080/openapi3.json

.PHONY: http-health
http-health:
	http $(ENDPOINT)/health

.PHONY: http-user
http-user: http-login
	# User
	http $(ENDPOINT)/user 'Authorization: Bearer $(TOKEN)'
	http $(ENDPOINT)/user/1 'Authorization: Bearer $(TOKEN)'
	http POST $(ENDPOINT)/user user_name=harry email=newuser01@foo.com password=secret123 'Authorization: Bearer $(TOKEN)'
	http PUT $(ENDPOINT)/user/5 email=updateduser01@bar.com password=secret456 'Authorization: Bearer $(TOKEN)'
	http DELETE $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'
	http $(ENDPOINT)/user/5 'Authorization: Bearer $(TOKEN)'

.PHONY: http-tech
http-tech: http-login
	# Tech
	http $(ENDPOINT)/tech 'Authorization: Bearer $(TOKEN)'
	http $(ENDPOINT)/tech/1 'Authorization: Bearer $(TOKEN)'
	http POST $(ENDPOINT)/tech tech_name='New Tech' 'Authorization: Bearer $(TOKEN)'
	http PUT $(ENDPOINT)/tech/133 tech_name='Old Tech' 'Authorization: Bearer $(TOKEN)'
	http DELETE $(ENDPOINT)/tech/133 'Authorization: Bearer $(TOKEN)'

.PHONY: http-usertech
http-usertech: http-login
	http $(ENDPOINT)/user/1/liketech 'Authorization: Bearer $(TOKEN)'
	http $(ENDPOINT)/user/1/disliketech 'Authorization: Bearer $(TOKEN)'

.PHONY: http-userworkhistory
http-userworkhistory: http-login
	http $(ENDPOINT)/user/1/workhistory 'Authorization: Bearer $(TOKEN)'

.PHONY: http-company
http-company: http-login
	# Company
	http $(ENDPOINT)/company 'Authorization: Bearer $(TOKEN)'
	http $(ENDPOINT)/company/1 'Authorization: Bearer $(TOKEN)'
	http POST $(ENDPOINT)/company company_name=Google country_id:=230 address=California 'Authorization: Bearer $(TOKEN)'
	http PUT $(ENDPOINT)/company/7 company_name=Google3 address=Tokyo country_id:=10 'Authorization: Bearer $(TOKEN)'
	http DELETE $(ENDPOINT)/company/7 'Authorization: Bearer $(TOKEN)'


###############################################################################
# Curl [WIP]
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
# Docker
###############################################################################
#dcgobld:
#	#goplus:1.11.5
#	docker build -t hirokiy/goplus:1.11.5 -f ./docker/golang/Dockerfile .
#	docker push hirokiy/goplus:1.11.5
#
#dcins:
#	docker exec -it gogoa_webserver_1 bash
#
#dctest:
#	docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go"
#	#docker-compose exec webserver /bin/sh -c "go test -v ext/cmd/*.go -f /go/src/github.com/hiromaily/go-goa/resources/tomls/docker.toml"
#
#dcpush:
#	docker push hirokiy/go-goa:1.0


###############################################################################
# Heroku
###############################################################################
#heroku_init:
#	heroku plugins:install heroku-container-registry
#	heroku container:login
#	heroku create goa-web
#	#heroku container:push web
#	#cd docker/mysql;heroku container:push mysql
#
#heroku_after_change_app_name:
#	#git remote remove heroku [Important!!]
#
#heroku_settings:
#	heroku apps
#	#goa-web
#
#heroku_open:
#	open -a goa-web
#	#https://goa-web.herokuapp.com/
#
#heroku_remove:
#	heroku apps:destroy


###############################################################################
# Build Image for Heroku
###############################################################################
#heroku_build_docker1:
#	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
#	docker push hirokiy/goapack_base:latest
#
#heroku_build_docker2:
#	docker build --no-cache -t hirokiy/goapack:latest -f ./docker/Dockerfile.heroku .
#
#heroku_build_multistage:
#	#It works!
#	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.multistage.heroku .
#
#heroku_bldfull:
#	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
#	docker build --no-cache -t hirokiy/goapack:latest -f ./docker/Dockerfile.heroku .
#
#heroku_exec_docker:
#	docker run -p 8080:8080 --name goapack hirokiy/goapack:latest
#	docker stop goapack
#
#heroku_upd:
#	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.heroku .
#	docker push registry.heroku.com/goa-web/web
#
#heroku_updfull:
#	docker build --no-cache -t hirokiy/goapack_base:latest -f ./docker/Dockerfile.base.heroku .
#	docker build --no-cache -t registry.heroku.com/goa-web/web -f ./docker/Dockerfile.heroku .
#	docker push registry.heroku.com/goa-web/web


###############################################################################
# Build Image for GCP Kubernetes
###############################################################################
#dcp_build:
#	docker-compose build --no-cache webserver
#	#docker build --no-cache -t hirokiy/go-goa-mysql:latest -f ./docker/mysql/Dockerfile .
#
#dcp_push:
#	docker push hirokiy/go-goa:1.0
#	#docker push hirokiy/go-goa-mysql:latest


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
