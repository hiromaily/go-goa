package design

import (
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

var JWT = JWTSecurity("jwt", func() {
	Scope("api:access", "API access") // Define "api:access" scope
})

var _ = Service("auth", func() {
	Description("The auth service performs login with JWT")

	HTTP(func() { // HTTP mapping for error responses
		// base path
		Path("/auth")
	})

	Method("login", func() {
		Description("Login and return jwt token")

		Payload(func() {
			Extend(types.PayloadLogin)
			Required("email", "password")
		})
		Result(types.RTAuthorized)
		// Use HTTP status 401 for 'Unauthorized' errors.

		// FIXME: somehow MakeUnauthorized() is gone
		// Note: this type of expression can't generate auth.MakeUnauthorized() function
		//Error("unauthorized", String, "Credentials are invalid")
		Error("Unauthorized")

		HTTP(func() {
			POST("/login")
			//Response(StatusOK)
			Response(StatusOK, func() {
				// mapping <variable-name:header-name>
				Header("token:Authorization", String, "JWT token", func() {
					Pattern("^Bearer [^ ]+$")
				})
			})
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
