package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Tech
//-----------------------------------------------------------------------------
var _ = Service(resourcePrefix+"tech", func() {
	Description("The company service returns company data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/tech")
	})
	//var _ = Resource(resourcePrefix+"tech", func() {
	//
	//	DefaultMedia(Tech)
	//	BasePath("/tech")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	// Method defines a single service method.
	Method("techList", func() { //just name
		Description("List all techs")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(RTTech))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
	//Action("TechList", func() {
	//	Routing(
	//		GET(""),
	//	)
	//	Description("List all techs")
	//	//NoSecurity()
	//
	//	Response(OK, CollectionOf(Tech)) //multiple response
	//	Response(NoContent)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("getTech", func() { //just name
		Description("get tech with given tech id")
		Error("BadRequest")
		Error("NotFound")
		//query string
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID") //params
		})
		Result(RTCompany)
		HTTP(func() {
			GET("/{techID}")
			Response(StatusOK)
		})
	})
	//Action("GetTech", func() {
	//	Routing(
	//		GET("/:techID"),
	//	)
	//	Description("Retrieve tech with given tech id")
	//	Params(func() {
	//		Param("techID", Integer, "Tech ID")
	//	})
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("createTech", func() { //just name
		Description("Create new tech")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("name", String, "Tech Name", fieldTechName)
			Required("name")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})
	//Action("CreateTech", func() {
	//	Routing(
	//		POST(""),
	//	)
	//	Description("Create new tech")
	//	Payload(TechPayload, func() {
	//		//TODO:required value in media_type is given priority over this part...
	//		Required("name")
	//	})
	//
	//	//no response template named "Created" in resource "hy_tech" action "CreateTech"
	//	//=>it should be defined in api_definition.go
	//	//Response(Created, "^/tech/[0-9]+$")
	//	Response(OK)
	//	Response(Created)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("updateTech", func() { //just name
		Description("Change tech properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID") //params
			Attribute("name", String, "Tech Name", fieldTechName)
			Required("name")
		})
		HTTP(func() {
			PUT("/{techID}")
			Response(StatusOK)
		})
	})
	//Action("UpdateTech", func() {
	//	Routing(
	//		PUT("/:techID"),
	//	)
	//	Description("Change tech properties")
	//
	//	Params(func() {
	//		Param("techID", Integer, "Tech ID")
	//	})
	//	Payload(TechPayload, func() {
	//		Required("name")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("deleteTech", func() { //just name
		Description("Delete tech")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID") //params
		})
		HTTP(func() {
			DELETE("/{techID}")
			Response(StatusOK)
		})
	})
	//Action("DeleteTech", func() {
	//	Routing(
	//		DELETE("/:techID"),
	//	)
	//	Description("Delete tech")
	//
	//	Params(func() {
	//		Param("techID", Integer, "Tech ID")
	//	})
	//
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})
})
