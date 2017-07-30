package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var resourcePrefix string = "hy_"

//Resources group related API endpoints
// This is request object
// Payload is sending data to server

var _ = Resource(resourcePrefix+"user", func() {

	DefaultMedia(User) //Response Media Type
	BasePath("/user")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	//Actions define a single API endpoint
	// This name should be unique
	Action("UserList", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users.")
		//Responses define the shape and status code
		//Response(OK, CollectionOf(User))
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
				View("tiny")
			}))
			//	Headers(func() {                // Headers list the response HTTP headers
			//		Header("X-Request-Id")  // Header syntax is identical to Attribute's
			//      Required("X-Request-Id")
			//	})
		})

		//HTTP Request Header
		//Headers(func() {                  // Headers describe relevant action headers
		//    Header("Authorization", String)
		//    Header("X-Account", Integer)
		//    Required("Authorization", "X-Account")
		//})
	})

	Action("GetUser", func() {
		Routing(
			GET("/:userID"),
		)
		Description("Retrieve user with given id.")
		Params(func() {
			Param("userID", Integer, "User ID", func() {
				Minimum(1)
			})
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("CreateUser", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		//Payload(func() {
		//	Member("name")
		//	Required("name")
		//})
		Payload(UserPayload, func() {
			Required("name", "email")
		})

		Response(Created, "/user/[0-9]+") //[??]
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateUser", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user properties")
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Payload(UserPayload)

		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteUser", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource(resourcePrefix+"company", func() {

	DefaultMedia(Company)
	BasePath("/company")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	// Parent sets the resource parent
	//Parent("user")  //TODO: how does it work??

	Action("CompanyList", func() {
		Routing(
			GET(""),
		)
		Description("List all companies")
		//NoSecurity()
		Response(OK, func() {
			Media(CollectionOf(Company, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("GetCompany", func() {
		Routing(
			GET("/:companyID"),
		)
		Description("Retrieve company with given id")
		Params(func() {
			Param("companyID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	//TODO: is it for websocket??
	//Action("watch", func() {
	//	Routing(
	//		GET("/:bottleID/watch"),
	//	)
	//	Scheme("ws")
	//	Description("Retrieve bottle with given id")
	//	Params(func() {
	//		Param("bottleID", Integer)
	//	})
	//	Response(SwitchingProtocols)
	//	Response(BadRequest, ErrorMedia)
	//})

	Action("CreateCompany", func() {
		Routing(
			POST(""),
		)
		Description("Record new company")
		Payload(CompanyPayload, func() {
			Required("name", "country", "address")
		})
		Response(Created, "^/user/[0-9]+/company/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateCompany", func() {
		Routing(
			PUT("/:companyID"),
		)
		Params(func() {
			Param("companyID", Integer)
		})
		Payload(CompanyPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	//TODO:somohow confliction error occur
	//=> type miss
	Action("DeleteCompany", func() {
		Routing(
			DELETE("/:companyID"),
		)
		Params(func() {
			Param("companyID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/*filepath", "public/")
	Files("/swagger-ui/*filepath", "resources/swagger-ui/dist/")
	Files("/swagger.json", "goa/swagger/swagger.json")
})

//var _ = Resource("public", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/", "public/index.html")
//})

//var _ = Resource("js", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/js/*filepath", "public/js")
//})

//var _ = Resource("swagger", func() {
//	Origin("*", func() {
//		Methods("GET", "OPTIONS")
//	})
//	Files("/swagger.json", "public/swagger/swagger.json")
//})
