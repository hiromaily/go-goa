package design

import (
	. "goa.design/goa/v3/dsl"
)

// https://hackmd.io/@vWEf3Ct8QVu2_w2RDTRIaQ/ByyoCY5Hr

// Additional file for auth when login
// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	//Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

var PayloadLogin = Type("PayloadLogin", func() {
	Attribute("email", String, "E-mail of user", fieldEmail)
	Attribute("password", String, "Password", fieldPassword)
})

// Resource to Service
// Action to Method

// Service defines a group of remotely accessible methods that are hosted
// together. The service DSL makes it possible to define the methods, their
// input and output as well as the errors they may return independently of the
// underlying transport (HTTP or gRPC). The transport specific DSLs defined by
// the HTTP and GRPC functions define the mapping between the input, output and
// error type attributes and the transport data (e.g. HTTP headers, HTTP bodies
// or gRPC messages).
var _ = Service("auth", func() {
	Description("The auth service performs login with JWT")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	Error("Unauthorized")
	//var _ = Resource("auth", func() {
	//	BasePath("/auth")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	// Method defines a single service method.
	Method("login", func() { // just name
		NoSecurity()
		//Note: embed Type causes wrong generated code
		//Payload(PayloadLogin, func() {
		//	Required("email", "password")
		//})
		Payload(func() {
			Attribute("email", String, "E-mail of user", fieldEmail)
			Attribute("password", String, "Password", fieldPassword)
			Required("email", "password")
		})
		Result(RTAuthorized)
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			POST("/login")
			Response(StatusOK)
			//FIXME: response token for JWT
			//Response(StatusOK, func() {
			//	Headers(func() {
			//		Header("Authorization", String, "Generated JWT")
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
	//	//Response(BadRequest)
	//})
})
