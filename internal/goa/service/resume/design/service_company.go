package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Company
//-----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"company", func() {
	Description("The company service returns company data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/company")
	})
	//var _ = Resource(resourcePrefix+"company", func() {
	//
	//	DefaultMedia(Company)
	//	BasePath("/company")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	// Method defines a single service method.
	Method("companyList", func() { //just name
		Description("List all companies")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(RTCompany))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
	//Action("CompanyList", func() {
	//	Routing(
	//		GET(""),
	//	)
	//	Description("List all companies")
	//	//NoSecurity()
	//
	//	Response(OK, CollectionOf(Company)) //multiple response
	//	Response(NoContent)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("getCompanyGroup", func() { //just name
		Description("Retrieve company with given company_id")
		Error("NoContent")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID") //params
			Attribute("hq_flg", String, "Head Quarters flag", func() {
				Enum("1", "0")
			})
		})
		Result(CollectionOf(RTCompany))
		HTTP(func() {
			GET("/{companyID}")
			Response(StatusOK)
		})
	})
	//Action("GetCompanyGroup", func() {
	//	Routing(
	//		GET("/:companyID"),
	//	)
	//	Description("Retrieve company with given company_id")
	//	Params(func() {
	//		Param("companyID", Integer, "Company ID")
	//		//query string
	//		Param("hq_flg", func() {
	//			Enum("1", "0")
	//		})
	//	})
	//	Response(OK, CollectionOf(Company))
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("createCompany", func() { //just name
		Description("Create new company")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("name", String, "Company Name", fieldCompanyName)
			Attribute("country_id", Int, "Country's ID", fieldCountryID)
			Attribute("address", String, "Address of company", fieldAddress)
			Required("name", "country_id", "address")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})
	//Action("CreateCompany", func() {
	//	Routing(
	//		POST(""),
	//	)
	//	Description("Create new company")
	//	Payload(CompanyPayload, func() {
	//		//TODO:required value in media_type is given priority over this part...
	//		Required("name", "country_id", "address")
	//	})
	//
	//	//no response template named "Created" in resource "hy_company" action "CreateCompany"
	//	//=>it should be defined in api_definition.go
	//	//Response(Created, "^/user/[0-9]+/company/[0-9]+$")
	//	Response(OK)
	//	Response(Created)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("updateCompany", func() { //just name
		Description("Change company properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID") //params
			Attribute("name", String, "Company Name", fieldCompanyName)
			Attribute("country_id", Int, "Country's ID", fieldCountryID)
			Attribute("address", String, "Address of company", fieldAddress)
			Required("name", "country_id", "address")
		})
		HTTP(func() {
			PUT("/{companyID}")
			Response(StatusOK)
		})
	})
	//Action("UpdateCompany", func() {
	//	Routing(
	//		PUT("/:companyID"),
	//	)
	//	Description("Change company properties")
	//
	//	Params(func() {
	//		Param("companyID", Integer, "Company ID")
	//	})
	//	Payload(CompanyPayload, func() {
	//		Required("name", "country_id", "address")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("deleteCompany", func() { //just name
		Description("Delete company")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID") //params
		})
		HTTP(func() {
			DELETE("/{companyID}")
			Response(StatusOK)
		})
	})
	//Action("DeleteCompany", func() {
	//	Routing(
	//		DELETE("/:companyID"),
	//	)
	//	Description("Delete company")
	//
	//	Params(func() {
	//		Param("companyID", Integer, "Company ID")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})
})

//-----------------------------------------------------------------------------
// Company branch
//-----------------------------------------------------------------------------
var _ = Service(resourcePrefix+"companybranch", func() {
	Description("The company branch service returns company branch data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/company/branch")
	})
	//var _ = Resource(resourcePrefix+"companybranch", func() {
	//
	//	DefaultMedia(Company)
	//	BasePath("/company/branch")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	Method("getCompanyBranch", func() { //just name
		Description("Get company branch with given id")
		Error("NotFound")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyDetailID", Int, "Company detail ID")
		})
		HTTP(func() {
			GET("/{companyDetailID}")
			Response(StatusOK)
		})
	})
	//Action("GetCompanyBranch", func() {
	//	Routing(
	//		GET("/:ID"),
	//	)
	//	Description("Retrieve company branch with given id")
	//	Params(func() {
	//		//Param("companyID", Integer, "Company ID")
	//		Param("ID", Integer, "Company detail ID")
	//	})
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("createCompanyBranch", func() { //just name
		Description("Create new company branch")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID") //params
			Attribute("country_id", Int, "Country's ID", fieldCountryID)
			Attribute("address", String, "Address of company", fieldAddress)
			Required("country_id", "address")
		})
		HTTP(func() {
			//TODO:somehow POST path should be unique
			POST("/{companyID}")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})
	//Action("CreateCompanyBranch", func() {
	//	//TODO:somehow POST path should be unique
	//	Routing(
	//		//POST("/:companyID"),
	//		POST("/:ID"),
	//	)
	//	Description("Create new company branch")
	//	Params(func() {
	//		//Param("companyID", Integer, "Company ID")
	//		Param("ID", Integer, "Company ID") //TODO:Though this name is ID, but companyID is set.
	//	})
	//	Payload(CompanyTinyPayload, func() {
	//		//TODO:required value in media_type is given priority over this part...
	//		Required("country_id", "address")
	//	})
	//
	//	//no response template named "Created" in resource "hy_company" action "CreateCompany"
	//	//=>it should be defined in api_definition.go
	//	//Response(Created, "^/user/[0-9]+/company/[0-9]+$")
	//	Response(OK)
	//	Response(Created)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("updateCompanyBranch", func() { //just name
		Description("Change company branch properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyDetailID", Int, "Company detail ID") //params
			Attribute("country_id", Int, "Country's ID", fieldCountryID)
			Attribute("address", String, "Address of company", fieldAddress)
			Required("country_id", "address")
		})
		HTTP(func() {
			PUT("/{companyDetailID}")
			Response(StatusOK)
		})
	})
	//Action("UpdateCompanyBranch", func() {
	//	Routing(
	//		PUT("/:ID"),
	//	)
	//	Description("Change company branch properties")
	//
	//	Params(func() {
	//		//Param("companyID", Integer, "Company ID")
	//		Param("ID", Integer, "Company detail ID")
	//	})
	//	Payload(CompanyTinyPayload, func() {
	//		Required("country_id", "address")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("deleteCompanyBranch", func() { //just name
		Description("Delete company branch")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyDetailID", Int, "Company detail ID")
		})
		HTTP(func() {
			DELETE("/{companyDetailID}")
			Response(StatusOK)
		})
	})
	//Action("DeleteCompanyBranch", func() {
	//	Routing(
	//		DELETE("/:ID"),
	//	)
	//	Description("Delete company branch")
	//
	//	Params(func() {
	//		//Param("companyID", Integer, "Company ID")
	//		Param("ID", Integer, "Company detail ID")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})
})
