#!/bin/sh

# gen design, example
cd internal/goa/service/resume
goa gen resume/design
goa example resume/design

# replace
grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'
grep -l '\"resume\"' ./internal/goa/service/resume/cmd/resume_api_server-cli/*.go | xargs sed -i.bak -e 's|\"resume\"|\"github.com/hiromaily/go-goa/pkg/goa/service/resume\"|g'

# mv
mv cmd/resume_api_server/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/server/
mv cmd/resume_api_server-cli/*.go ${GOPATH}/src/github.com/hiromaily/go-goa/cmd/resume-api/cli/
mv auth.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv health.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_company.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_companybranch.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_tech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_user.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_user_work_history.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv hy_usertech.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
mv public.go ${GOPATH}/src/github.com/hiromaily/go-goa/pkg/goa/service/resume/
rm -rf cmd

cd ${GOPATH}/src/github.com/hiromaily/go-goa

#replace `resumeapi "resume"`
#     to `github.com/hiromaily/go-goa/pkg/goa/service/resume`
#in main.go


#sed -e "1s/main/goa/" ./internal/goa/service/resume/design/example/cmd/resume_api_server/http.go >> ./tmp/tmp.go
#mv -f ./tmp/tmp.go ./cmd/resume-api/server/http.go

#resume/example/gen/
#=>
#github.com/hiromaily/go-goa/pkg/goa/service

#resume/gen/auth/
#=>
#github.com/hiromaily/go-goa/internal/goa/service/resume/gen/auth
