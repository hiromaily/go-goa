package design

import (
	. "goa.design/goa/v3/dsl"
	//cors "goa.design/plugins/v3/cors/dsl"
)

// Reference
// - https://pkg.go.dev/goa.design/goa/v3
// - https://pkg.go.dev/goa.design/goa/v3/dsl [important]
// - https://goa.design/learn/upgrading/

var (
	resourcePrefix = "hy_"
	domain         = "localhost"
)

// API
var _ = API("resume", func() {
	Title("resume API")
	Description("go-goa API with goa framework")
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
		Description("goa README")
		URL("https://github.com/goadesign/goa")
	})
	// CORS policy: https://github.com/goadesign/plugins/tree/v3/cors
	// - Sets CORS response headers for requests with Origin header matching the regular expression ".*domain.*"
	//cors.Origin(fmt.Sprintf("/%s/", domain), func() {
	//	cors.Headers("*")
	//	cors.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
	//	cors.MaxAge(600)
	//	cors.Credentials()
	//})

	// Server
	// this name can be used as directory name after generated in cmd
	Server("resume", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:9090")
		})
		// List the services hosted by this server.
		Services("auth", resourcePrefix+"company", "health", resourcePrefix+"tech",
			resourcePrefix+"user", resourcePrefix+"usertech", resourcePrefix+"userWorkHistory")
	})
	// HTTP: https://pkg.go.dev/goa.design/goa/v3/dsl#HTTP
	HTTP(func() {
		// MIME type support
		Consumes("application/json", "application/xml")
	})
})
