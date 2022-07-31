package design

import (
	. "goa.design/goa/v3/dsl"
	"resume/design/types"
)

// https://hackmd.io/@vWEf3Ct8QVu2_w2RDTRIaQ/ByyoCY5Hr

var JWT = JWTSecurity("jwt", func() {
	Scope("api:access", "API access") // Define "api:access" scope
})

var _ = Service("auth", func() {
	Description("The auth service performs login with JWT")

	Security(JWT)
	Error("Unauthorized")
	HTTP(func() { // HTTP mapping for error responses
		// Use HTTP status 401 for 'Unauthorized' errors.
		Response("Unauthorized", StatusUnauthorized)
	})

	Method("login", func() {
		NoSecurity()
		Payload(func() {
			Extend(types.PayloadLogin)
			Required("email", "password")
		})
		Result(types.RTAuthorized)
		HTTP(func() {
			POST("/login")

			//FIXME: response token for JWT
			//Response(StatusOK, func() {
			//	//Header("auth:Authorization", String, "Auth token", func() {
			//	//	Pattern("^Bearer [^ ]+$")
			//	//})
			//})
		})
	})

	//Action("Login", func() {
	//	Description("user login")
	//	NoSecurity() // Override the need for auth
	//	Routing(POST("login"))
	//
	//	Payload(LoginPayload, func() {
	//		Required("email", "password")
	//	})
	//
	//	Response(OK, func() {
	//		Headers(func() {
	//			Header("Authorization", String, "Generated JWT")
	//		})
	//		Media(Authorized)
	//	})
	//	Response(Unauthorized)
	//})
})
