package design

import (
	. "goa.design/goa/v3/dsl"
)

var resourcePrefix string = "hy_"

//Resources group related API endpoints
// This is request object
// Payload is sending data to server

//-----------------------------------------------------------------------------
// User
//-----------------------------------------------------------------------------
var _ = Service(resourcePrefix+"user", func() {
	Description("The user service returns user data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/user")
	})
	//var _ = Resource(resourcePrefix+"user", func() {
	//
	//	DefaultMedia(User) //Response Media Type
	//	BasePath("/user")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	Method("userList", func() { //just name
		Description("List all users")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(RTUser))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})
	//Action("UserList", func() {
	//	Routing(
	//		GET(""),
	//	)
	//	Description("Retrieve all users.")
	//	Response(OK, CollectionOf(User)) //multiple response
	//	Response(NoContent)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("getUser", func() { //just name
		Description("get user with given user id")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
		})
		Result(RTUser)
		HTTP(func() {
			GET("/{userID}")
			Response(StatusOK)
		})
	})
	//Action("GetUser", func() {
	//	Routing(
	//		GET("/:userID"),
	//	)
	//	Description("Retrieve user with given id.")
	//	Params(func() {
	//		Param("userID", Integer, "User ID", func() {
	//			Minimum(1)
	//		})
	//	})
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("createUser", func() { //just name
		Description("Create new user")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
			Attribute("user_name", String, "User name", fieldUserName)
			Attribute("email", String, "E-mail of user", fieldEmail)
			Attribute("password", String, "Password", fieldPassword)
			Required("user_name", "email", "password")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})
	//Action("CreateUser", func() {
	//	Routing(
	//		POST(""),
	//	)
	//	Description("Create new user")
	//	Payload(UserPayload, func() {
	//		Required("user_name", "email", "password")
	//	})
	//	Response(OK)
	//	Response(Created)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("updateUser", func() { //just name
		Description("Change user properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
			Attribute("user_name", String, "User name", fieldUserName)
			Attribute("email", String, "E-mail of user", fieldEmail)
			Attribute("password", String, "Password", fieldPassword)
			Required("user_name", "email", "password")
		})
		HTTP(func() {
			PUT("/{userID}")
			Response(StatusOK)
		})
	})
	//Action("UpdateUser", func() {
	//	Routing(
	//		PUT("/:userID"),
	//	)
	//	Description("Change user properties")
	//	Params(func() {
	//		Param("userID", Integer, "User ID")
	//	})
	//	Payload(UserPayload, func() {
	//		Required("user_name", "email", "password")
	//	})
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("deleteUser", func() { //just name
		Description("Delete user")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
		})
		HTTP(func() {
			DELETE("/{userID}")
			Response(StatusOK)
		})
	})
	//Action("DeleteUser", func() {
	//	Routing(
	//		DELETE("/:userID"),
	//	)
	//	Description("Delete user ")
	//	Params(func() {
	//		Param("userID", Integer, "User ID")
	//	})
	//	Response(OK)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})
})

//-----------------------------------------------------------------------------
// User Tech
//-----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"usertech", func() {
	Description("The user service returns user data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/user")
	})
	//var _ = Resource(resourcePrefix+"usertech", func() {
	//
	//	DefaultMedia(UserTech) //Response Media Type
	//	BasePath("/user")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	Method("getUserLikeTech", func() { //just name
		Description("get user's favorite techs")
		Error("BadRequest")
		Error("NotFound")
		Error("Unauthorized")
		//query string
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
		})
		Result(CollectionOf(RTUserTech))
		HTTP(func() {
			GET("/{userID}/liketech")
			Response(StatusOK)
		})
	})
	//Action("GetUserLikeTech", func() {
	//	Routing(
	//		GET("/:userID/liketech"),
	//	)
	//	Description("Retrieve user's favorite techs.")
	//	// Params is for get parameter
	//	Params(func() {
	//		Param("userID", Integer, "User ID", func() {
	//			Minimum(1)
	//		})
	//	})
	//	Response(OK, CollectionOf(UserTech)) //multiple response
	//	Response(Unauthorized)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

	Method("getUserDisLikeTech", func() { //just name
		Description("get user's dislike techs")
		Error("BadRequest")
		Error("NotFound")
		Error("Unauthorized")
		//query string
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() { //params
				Minimum(1)
			})
		})
		Result(CollectionOf(RTUserTech))
		HTTP(func() {
			GET("/{userID}/disliketech")
			Response(StatusOK)
		})
	})
	//Action("GetUserDislikeTech", func() {
	//	Routing(
	//		GET("/:userID/disliketech"),
	//	)
	//	Description("Retrieve user's dislike techs.")
	//	// Params is for get parameter
	//	Params(func() {
	//		Param("userID", Integer, "User ID", func() {
	//			Minimum(1)
	//		})
	//	})
	//	Response(OK, CollectionOf(UserTech)) //multiple response
	//	Response(Unauthorized)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})
})

//-----------------------------------------------------------------------------
// User Work History
//-----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"userWorkHistory", func() {
	Description("The user service returns user data")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	HTTP(func() { // HTTP mapping for error responses
		Path("/user")
	})
	//var _ = Resource(resourcePrefix+"userWorkHistory", func() {
	//
	//	DefaultMedia(UserWorkHistory) //Response Media Type
	//	BasePath("/user")
	//
	//	Security(JWT, func() { // Use JWT to auth requests to this endpoint
	//		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	//	})

	Method("getUserWorkHistory", func() { //just name
		Description("get user's work history")
		Error("BadRequest")
		Error("NotFound")
		Error("Unauthorized")
		//query string
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() {
				Minimum(1)
			})
		})
		Result(CollectionOf(RTUserWorkHistory))
		HTTP(func() {
			GET("/{userID}/workhistory")
			Response(StatusOK)
			//cannot use &expr.ResultTypeExpr{UserTypeExpr:(*expr.UserTypeExpr)(0xc0000ffd10),
			// Identifier:"application/vnd.userworkhistory+json; type=collection", ContentType:"",
			// Views:[]*expr.ViewExpr(nil)} (type *expr.ResultTypeExpr)
			// as type int (HTTP status code) or function
			// in service "hy_userWorkHistory" HTTP endpoint "getUserWorkHistory"
		})
	})
	//Action("GetUserWorkHistory", func() {
	//	Routing(
	//		GET("/:userID/workhistory"),
	//	)
	//	Description("Retrieve user's work history.")
	//	// Params is for get parameter
	//	Params(func() {
	//		Param("userID", Integer, "User ID", func() {
	//			Minimum(1)
	//		})
	//	})
	//	Response(OK, CollectionOf(UserWorkHistory)) //multiple response
	//	Response(Unauthorized)
	//	Response(NotFound)
	//	Response(BadRequest, ErrorMedia)
	//})

})
