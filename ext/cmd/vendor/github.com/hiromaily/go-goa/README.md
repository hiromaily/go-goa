# go-goa
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


## Licence
[MIT](https://github.com/hiromaily/go-goa/blob/master/LICENSE)

## Author

[hiromaily](https://github.com/hiromaily)
