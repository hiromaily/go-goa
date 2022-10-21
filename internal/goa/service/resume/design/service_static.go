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
	Files("/assets/{*filepath}", "assets/")
	Files("/swagger-ui/*filepath", "resources/swagger-ui/dist/")
	//Files("/swagger.json", "goa/swagger/swagger.json")
	Files("/openapi.json", "../gen/http/openapi.json")
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
