package types

import (
	. "goa.design/goa/v3/dsl"
)

// PayloadUserTech defines the data structure used in the create user-tech request body
var PayloadUserTech = Type("PayloadUserTech", func() {
	Attribute("tech_name", String, "Tech name", fieldTechName)
})

// RTUserTech is the user tech resource result type.
var RTUserTech = ResultType("application/vnd.usertech+json", func() {
	Description("A user information")

	ContentType("application/json")

	Reference(PayloadUserTech)

	Attributes(func() {
		Attribute("id", Int, "Tech ID", fieldID)
		Attribute("tech_name")

		// Required adds a "required" validation to the attribute.
		//Required("tech_name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("tech_name")
	})

	View("techName", func() {
		Description("only tech name")
		Attribute("tech_name")
	})
})
