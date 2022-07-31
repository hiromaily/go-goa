
.PHONY: mod
mod:
	go mod tidy -compat=1.17

.PHONY: gen
gen:
	goa gen resume/design

.PHONY: gen-example
example:
	goa example resume/design

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

.PHONY: gen-all
gen-all: gen-design gen-example replace-example move-example
