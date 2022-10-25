package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = Service("static", func() {
	Description("The static service returns static files")
	// Sets CORS response headers for requests with Origin header matching the regular expression ".*domain.*"
	cors.Origin("*", func() {
		cors.Methods("GET", "OPTIONS")
	})
	// path, actual stored path
	Files("/assets/{*filepath}", "docs/")
	//Files("/openapi.json", "../gen/http/openapi.json")
	//Files("/swagger-ui/*filepath", "resources/swagger-ui/dist/")
	//Files("/swagger.json", "goa/swagger/swagger.json")
})
