package design

import (
	. "goa.design/goa/v3/dsl"
	"resume/design/types"
)

var _ = Service(resourcePrefix+"user", func() {
	Description("The user service returns user data")

	Security(JWT)
	HTTP(func() {
		Path("/user")
	})

	Method("userList", func() {
		Description("List all users")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTUser))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})

	Method("getUser", func() {
		Description("get user with given user id")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() {
				Minimum(1)
			})
		})
		Result(types.RTUser)
		HTTP(func() {
			GET("/{userID}")
			Response(StatusOK)
		})
	})

	Method("createUser", func() {
		Description("Create new user")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadUser)
			Required("user_name", "email", "password")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})

	Method("updateUser", func() {
		Description("Change user properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadUser)
			Required("user_name", "email", "password")
		})
		HTTP(func() {
			PUT("/{userID}")
			Response(StatusOK)
		})
	})

	Method("deleteUser", func() {
		Description("Delete user")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("userID", Int, "User ID", func() {
				Minimum(1)
			})
		})
		HTTP(func() {
			DELETE("/{userID}")
			Response(StatusOK)
		})
	})
})

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
