package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//This is basic whole API definition
// API defines the microservice endpoint and
// other global properties

// This is the cellar application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("api", func() {
	Title("hiromaily API")
	Description("hiromaily API with goa")
	Contact(func() {
		Name("hiromaily")
		Email("hiromaily@gmail.com")
		URL("https://hiromaily.github.io/")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("goa guide")
		URL("http://goa.design/getting-started.html")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/api")

	Consumes("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
	Consumes("application/xml") // ★
	Consumes("application/json")

	//Consumes("application/json")
	//Produces("application/json")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

//func ObjectCreated(mt *MediaTypeDefinition) func() {
//	return func() {
//		Description("Object created")
//		Status(201)
//		Media(mt)
//	}
//}
