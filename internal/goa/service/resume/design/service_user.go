package design

import (
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

// ----------------------------------------------------------------------------
// User Service
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"user", func() {
	Description("The user service returns user data")

	Security(JWT)
	HTTP(func() {
		// BasePath
		Path("/user")
	})

	Error("NotFound")
	Error("BadRequest")

	Method("userList", func() {
		Description("List all users")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTUser))

		HTTP(func() {
			GET("")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("getUser", func() {
		Description("Get user by given user id")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Required("user_id")
		})
		Result(types.RTUser)

		HTTP(func() {
			GET("/{user_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("createUser", func() {
		Description("Create new user")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadUser)
			Required("user_name", "email", "password")
		})
		Result(types.RTUser)
		//Result(types.RTUser, func() {
		//	View("id")
		//	Required("id")
		//})

		HTTP(func() {
			POST("")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("updateUser", func() {
		Description("Update user data")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadUser)
			Required("user_id")
		})

		HTTP(func() {
			PUT("/{user_id}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("deleteUser", func() {
		Description("Delete user")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Required("user_id")
		})

		HTTP(func() {
			DELETE("/{user_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
})

// ----------------------------------------------------------------------------
// User Tech Service
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"usertech", func() {
	Description("The user service returns user data")

	Security(JWT)
	HTTP(func() {
		Path("/user")
	})

	Method("getUserLikeTech", func() {
		Description("get user's favorite techs")
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
		Result(CollectionOf(types.RTUserTech))
		HTTP(func() {
			GET("/{userID}/liketech")
			Response(StatusOK)
		})
	})

	Method("getUserDisLikeTech", func() {
		Description("get user's dislike techs")
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
		Result(CollectionOf(types.RTUserTech))
		HTTP(func() {
			GET("/{userID}/disliketech")
			Response(StatusOK)
		})
	})
})

var _ = Service(resourcePrefix+"userWorkHistory", func() {
	Description("The user service returns user data")

	Security(JWT)
	HTTP(func() {
		Path("/user")
	})

	Method("getUserWorkHistory", func() {
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
		Result(CollectionOf(types.RTUserWorkHistory))
		HTTP(func() {
			GET("/{userID}/workhistory")
			Response(StatusOK)
		})
	})
})
