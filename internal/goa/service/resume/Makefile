
.PHONY: install
install:
	go install goa.design/goa/v3/cmd/goa@v3
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


.PHONY: gen
gen:
	goa gen resume/design

.PHONY: gen-example
gen-example:
	goa example resume/design

.PHONY: build-example
build-example:
	go build ./cmd/resume
	go build ./cmd/resume-cli

all: gen gen-example build-example

# This can display all endpoints
# good to know after design modified
.PHONY: run-server
run-server:
	./resume

.PHONY: check-endpoint
check-endpoint:
	curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{}"
	curl http://localhost:8080/company
	curl http://localhost:8080/company/1
	curl -X POST http://localhost:8080/company -H "Content-Type: application/json" -d "{}"
	curl -X PUT http://localhost:8080/company/1 -H "Content-Type: application/json" -d "{}"
	curl -X DELETE http://localhost:8080/company/1

.PHONY: clean
clean:
	rm -rf resume resume-cli
	rm -rf gen cmd
	rm -rf auth.go health.go hy_company.go hy_companybranch.go hy_tech.go hy_user.go
	rm -rf hy_user_work_history.go hy_usertech.go public.go static.go


#.PHONY: replace-example
#replace-example:
#	grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'
#	grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server-cli/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'
#	#resumeapi "github.com/hiromaily/go-goa/pkg/goa/service/resume"
#	#resumeapi "resume"

#.PHONY: move-example
#move-example:
#	mv cmd/resume_api_server/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/server/
#	mv cmd/resume_api_server-cli/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/cli/
#	mv auth.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv health.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_company.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_companybranch.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_tech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_user.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_user_work_history.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv hy_usertech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	mv public.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
#	rm -rf internal/goa/service/resume/cmd

#.PHONY: gen-all
#gen-all: gen-design gen-example replace-example move-example
