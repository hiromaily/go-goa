package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// reference
// https://pkg.go.dev/goa.design/goa/v3
// https://pkg.go.dev/goa.design/goa/v3/dsl
// https://goa.design/learn/upgrading/

// https://goa.design/learn/getting-started/
// https://goa.design/ja/learn/getting-started/

// API defines a network service API. It provides the API name, description and other global
// properties. There may only be one API declaration in a given design package.
var _ = API("resume-api", func() {
	// Title sets the API title. It is used by the generated OpenAPI specification.
	Title("resume API")
	// Description sets the expression description.
	// Description may appear in API, Docs, Type or Attribute.
	Description("go-goa API with goa framework")
	// Contact sets the API contact information.
	Contact(func() {
		Name("hiromaily")
		Email("hiromaily@gmail.com")
		URL("https://hiromaily.github.io/")
	})
	// License sets the API license information.
	License(func() {
		Name("MIT")
		URL("https://github.com/goadesign/goa/blob/master/LICENSE")
	})
	// Docs provides external documentation URLs. It is used by the generated
	Docs(func() {
		Description("goa README")
		URL("https://github.com/goadesign/goa")
	})
	// CORS policy
	// https://github.com/goadesign/plugins/tree/v3/cors
	// Sets CORS response headers for requests with Origin header matching the regular expression ".*domain.*"
	cors.Origin("/swagger.goa.design/", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		cors.MaxAge(600)
		cors.Credentials()
	})

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server exposes as well as host details.
	// Not defining a server in a design has the same effect as defining a single
	// server that exposes all of the services defined in the design in a single
	// host listening on "locahost" and using port 80 for HTTP endpoints and 8080
	// for GRPC endpoints.
	Server("resume-api-server", func() {
		Host("localhost", func() {
			URI("http://localhost:8080/api")
			URI("grpc://localhost:9090")
		})
		// List the services hosted by this server.
		Services("auth", resourcePrefix+"company", resourcePrefix+"companybranch", "health",
			resourcePrefix+"tech", resourcePrefix+"user", resourcePrefix+"usertech", resourcePrefix+"userWorkHistory")
	})
	HTTP(func() {
		Consumes("application/json", "application/xml")
	})
})
