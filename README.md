# go-goa

[![Build Status](https://travis-ci.org/hiromaily/go-goa.svg?branch=master)](https://travis-ci.org/hiromaily/go-goa)
[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/go-goa)](https://goreportcard.com/report/github.com/hiromaily/go-goa)
[![codebeat badge](https://codebeat.co/badges/f2ee2ed0-5588-46f9-a47e-d50633a06739)](https://codebeat.co/projects/github-com-hiromaily-go-goa-master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f207ca57e48e456389341fc41bb06951)](https://www.codacy.com/app/hiromaily2/go-goa?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hiromaily/go-goa&amp;utm_campaign=Badge_Grade)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/go-goa/master/LICENSE)

go-goa is sameple for how to use goa framework

## About goa framework
* [goa](https://goa.design/)
* [Getting Started with goa](https://goa.design/learn/guide/)
* [The goa API Design Language](https://goa.design/design/overview/)
* [github.com/goadesign/goa/design](https://goa.design/reference/goa/design/)
* [github.com/goadesign/goa/design/apidsl](https://goa.design/reference/goa/design/apidsl/)


## Install
First, please check Makefile.

```bash
# Run Docker
dcup:
	docker-compose build
	docker-compose up
```

## Directory structure
| Directory NAME  | Description                        |
|:----------------|:-----------------------------------|
| docker          | docker resources                   |
| ext             | extension source code for logic    |
| goa             | generated files by goa             |
| public          | static files                       |
| resources       | swagger-ui, toml files and so on   |



## Usage
Please check Makefile.

```bash
# updating goa framework is required to keep latest.
updgoa:
	go get -u github.com/goadesign/goa/...  

# After modifying goa/design/* files
genfull: gencln aftergen
# ==> this regenerate goa related files. Don't worry, your logic files are in ext/...

# Integration Test
gotest:
	go test -v ext/cmd/*.go

```


## dependency management tool
[dep](https://github.com/golang/dep)

#### related files
* lock.json
* manifest.json
* vendor/


## API List
* ctrl=Health action=Health route=GET /api/_ah/health
* ctrl=Public files=goa/swagger/swagger.json route=GET /swagger.json
* ctrl=Public files=public/ route=GET /*filepath
* ctrl=Public files=resources/swagger-ui/dist/ route=GET /swagger-ui/*filepath
* ctrl=Public files=public/index.html route=GET /
* ctrl=Public files=resources/swagger-ui/dist/index.html route=GET /swagger-ui/
* ctrl=Auth action=Login route=POST /api/auth/login
* ctrl=HyUser action=CreateUser route=POST /api/user security=jwt
* ctrl=HyUser action=DeleteUser route=DELETE /api/user/:userID security=jwt
* ctrl=HyUser action=GetUser route=GET /api/user/:userID security=jwt
* ctrl=HyUser action=UpdateUser route=PUT /api/user/:userID security=jwt
* ctrl=HyUser action=UserList route=GET /api/user security=jwt
* ctrl=HyCompany action=CompanyList route=GET /api/company security=jwt
* ctrl=HyCompany action=CreateCompany route=POST /api/company security=jwt
* ctrl=HyCompany action=DeleteCompany route=DELETE /api/company/:companyID security=jwt
* ctrl=HyCompany action=GetCompanyGroup route=GET /api/company/:companyID security=jwt
* ctrl=HyCompany action=UpdateCompany route=PUT /api/company/:companyID security=jwt
* ctrl=HyCompanybranch action=CreateCompanyBranch route=POST /api/company/branch/:ID security=jwt
* ctrl=HyCompanybranch action=DeleteCompanyBranch route=DELETE /api/company/branch/:ID security=jwt
* ctrl=HyCompanybranch action=GetCompanyBranch route=GET /api/company/branch/:ID security=jwt
* ctrl=HyCompanybranch action=UpdateCompanyBranch route=PUT /api/company/branch/:ID security=jwt



## Performance
To eveluate performance, [hey](https://github.com/rakyll/hey) has been used.
Execute `Make bench`

```
# March 12 2017
hey -n 20000 -c 50 -m GET http://localhost:8080/api/user

Summary:
  Total:	4.1092 secs
  Slowest:	0.0939 secs
  Fastest:	0.0002 secs
  Average:	0.0101 secs
  Requests/sec:	4867.1221
  Total data:	60000 bytes
  Size/request:	3 bytes

Status code distribution:
  [200]	20000 responses
```

## Licence
[MIT](https://github.com/hiromaily/go-goa/blob/master/LICENSE)

## Author

[hiromaily](https://github.com/hiromaily)
