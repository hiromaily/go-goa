package design

import (
	"fmt"
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
		Path(fmt.Sprintf("%s/user", baseAPIDir))
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Unauthorized")

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
			Response("Unauthorized", StatusUnauthorized)
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
			Response("Unauthorized", StatusUnauthorized)
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
			Response("Unauthorized", StatusUnauthorized)
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
			Response("Unauthorized", StatusUnauthorized)
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
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})

// ----------------------------------------------------------------------------
// User Tech Service
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"usertech", func() {
	Description("The usertech service returns user's tech data")

	Security(JWT)
	HTTP(func() {
		Path(fmt.Sprintf("%s/user", baseAPIDir))
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Unauthorized")

	Method("getUserLikeTech", func() {
		Description("get user's favorite techs")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Required("user_id")
		})
		Result(CollectionOf(types.RTUserTech))

		HTTP(func() {
			GET("/{user_id}/liketech")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("getUserDisLikeTech", func() {
		Description("get user's dislike techs")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Required("user_id")
		})
		Result(CollectionOf(types.RTUserTech))

		HTTP(func() {
			GET("/{user_id}/disliketech")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})

// ----------------------------------------------------------------------------
// User Work History Service
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"userWorkHistory", func() {
	Description("The user work history service returns user working history data")

	Security(JWT)
	HTTP(func() {
		Path(fmt.Sprintf("%s/user", baseAPIDir))
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Unauthorized")

	Method("getUserWorkHistory", func() {
		Description("get user's work history")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("user_id", Int, "User ID", func() {
				Minimum(1)
			})
			Required("user_id")
		})
		Result(CollectionOf(types.RTUserWorkHistory))

		HTTP(func() {
			GET("/{user_id}/workhistory")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
