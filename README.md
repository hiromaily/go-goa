# go-goa

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

## dependency management tool
[dep](https://github.com/golang/dep)

#### related files
* lock.json
* manifest.json
* vendor/

## Usage
#### Which files should be changed for adding logic code?
* ./hy_company.go
* ./hy_user.go


## API List
* ctrl=Health action=Health route=GET /api/_ah/health

* ctrl=Public files=public/ route=GET /*filepath
* ctrl=Public files=public/index.html route=GET /
* ctrl=Public files=swagger-ui/dist/ route=GET /swagger-ui/*filepath
* ctrl=Public files=swagger-ui/dist/index.html route=GET /swagger-ui/

* ctrl=HyUser action=UserList route=GET /api/user
* ctrl=HyUser action=GetUser route=GET /api/user/:userID
* ctrl=HyUser action=CreateUser route=POST /api/user
* ctrl=HyUser action=UpdateUser route=PUT /api/user/:userID
* ctrl=HyUser action=DeleteUser route=DELETE /api/user/:userID

* ctrl=HyCompany action=CompanyList route=GET /api/company
* ctrl=HyCompany action=GetCompany route=GET /api/company/:companyID
* ctrl=HyCompany action=CreateCompany route=POST /api/company
* ctrl=HyCompany action=UpdateCompany route=PUT /api/company/:companyID
* ctrl=HyCompany action=DeleteCompany route=DELETE /api/company/:companyID

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
