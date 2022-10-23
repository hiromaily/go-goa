package design

import (
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

// goa メモ
// https://hackmd.io/@vWEf3Ct8QVu2_w2RDTRIaQ/ByyoCY5Hr

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
		Payload(func() {
			Extend(types.PayloadLogin)
			Required("email", "password")
		})
		Result(types.RTAuthorized)
		// Use HTTP status 401 for 'Unauthorized' errors.
		Error("Unauthorized")

		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			//Response(StatusOK, func() {
			//	Header("Authorization", String, func() {
			//		Pattern("^Bearer [^ ]+$")
			//	})
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
