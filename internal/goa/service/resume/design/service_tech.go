package design

import (
	"fmt"
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

// ----------------------------------------------------------------------------
// Service tech
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"tech", func() {
	Description("The tech service returns tech data")

	Security(JWT)
	HTTP(func() {
		// BasePath
		Path(fmt.Sprintf("%s/tech", baseAPIDir))
	})

	Error("NotFound")
	Error("BadRequest")

	Method("techList", func() {
		Description("List all techs")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTTech))

		HTTP(func() {
			GET("")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("getTech", func() {
		Description("returns tech by given tech id")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("tech_id", Int, "Tech ID", func() {
				Minimum(1)
			})
			Required("tech_id")
		})
		Result(types.RTTech)

		HTTP(func() {
			GET("/{tech_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("createTech", func() {
		Description("Create new tech")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadTech)
			Required("tech_name")
		})
		Result(types.RTTech)

		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("updateTech", func() {
		Description("Update tech data")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("tech_id", Int, "Tech ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadTech)
			Required("tech_id", "tech_name")
		})

		HTTP(func() {
			PUT("/{tech_id}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("deleteTech", func() {
		Description("Delete tech")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("tech_id", Int, "Tech ID", func() {
				Minimum(1)
			})
			Required("tech_id")
		})
		HTTP(func() {
			DELETE("/{tech_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
})
