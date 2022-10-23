package design

import (
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

var _ = Service(resourcePrefix+"tech", func() {
	Description("The company service returns company data")

	Security(JWT)
	HTTP(func() {
		Path("/tech")
	})

	Method("techList", func() {
		Description("List all techs")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTTech))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})

	Method("getTech", func() {
		Description("get tech with given tech id")
		Error("BadRequest")
		Error("NotFound")
		//query string
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID")
		})
		Result(types.RTCompany)
		HTTP(func() {
			GET("/{techID}")
			Response(StatusOK)
		})
	})

	Method("createTech", func() {
		Description("Create new tech")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadTech)
			Required("name")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})

	Method("updateTech", func() {
		Description("Change tech properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID")
			Extend(types.PayloadTech)
			Required("name")
		})
		HTTP(func() {
			PUT("/{techID}")
			Response(StatusOK)
		})
	})

	Method("deleteTech", func() {
		Description("Delete tech")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("techID", Int, "Tech ID")
		})
		HTTP(func() {
			DELETE("/{techID}")
			Response(StatusOK)
		})
	})
})
