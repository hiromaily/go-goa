#!/bin/sh

# gen design, example
cd internal/goa/service/resume && goa gen resume/design &&\
goa example resume/design -o ./example

# mv
mkdir tmp

mv internal/goa/service/resume/design/example/cmd/resume_api_server/http.go cmd/resume-api/server/
mv internal/goa/service/resume/design/example/cmd/resume_api_server/main.go cmd/resume-api/server/
mv internal/goa/service/resume/design/example/cmd/resume_api_server-cli/http.go cmd/resume-api/cli/
mv internal/goa/service/resume/design/example/cmd/resume_api_server-cli/main.go cmd/resume-api/cli/


#sed -e "1s/main/goa/" ./internal/goa/service/resume/design/example/cmd/resume_api_server/http.go >> ./tmp/tmp.go
#mv -f ./tmp/tmp.go ./cmd/resume-api/server/http.go

#resume/example/gen/
#=>
#github.com/hiromaily/go-goa/pkg/goa/service

#resume/gen/auth/
#=>
#github.com/hiromaily/go-goa/internal/goa/service/resume/gen/auth
