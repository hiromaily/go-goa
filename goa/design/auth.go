package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

// TODO:This code may be better to move to resources.go
// Resource jwt uses the JWTSecurity security scheme.
var _ = Resource("auth", func() {
	Description("This resource uses JWT to secure its endpoints")
	BasePath("/auth")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("Login", func() {
		Description("user login")
		NoSecurity() // Override the need for auth
		Routing(POST("login"))
		Payload(func() {
			Member("email")
			Member("password")
			Required("email", "password")
		})
		Response(OK, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
			Media(Authorized)
		})
		Response(Unauthorized)
	})
})
