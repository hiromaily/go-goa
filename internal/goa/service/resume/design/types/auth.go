package types

import (
	. "goa.design/goa/v3/dsl"
)

// PayloadLogin defines the data structure used in the login request body
var PayloadLogin = Type("PayloadLogin", func() {
	Attribute("email", String, "E-mail of user", fieldEmail)
	Attribute("password", String, "Password", fieldPassword)
})

// RTAuthorized is the authority resource result type
var RTAuthorized = ResultType("application/vnd.authorized+json", func() {
	Description("An authorized response")

	Attributes(func() {
		Attribute("token", String, "JWT token", func() {
			Example("token.string")
		})
		Attribute("id", Int, "User ID", fieldID)
		Required("token", "id")
	})

	// Note: View could be omitted if defined Attributes is fine
	View("default", func() {
		Attribute("token")
		Attribute("id")
	})
})
