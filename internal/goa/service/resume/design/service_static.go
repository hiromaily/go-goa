package design

import (
	. "goa.design/goa/v3/dsl"
	// cors "goa.design/plugins/v3/cors/dsl"
)

var _ = Service("static", func() {
	Description("The static service returns static files")

	// Sets CORS response headers for requests with Origin header matching the regular expression ".*domain.*"
	//cors.Origin("*", func() {
	//	cors.Methods("GET", "OPTIONS")
	//})

	// URL path, actual stored file path in server
	Files("/{*filepath}", "./docs/")

	Files("/openapi.json", "./internal/goa/service/resume/gen/http/openapi.json")
	Files("/openapi3.json", "./internal/goa/service/resume/gen/http/openapi3.json")
})
