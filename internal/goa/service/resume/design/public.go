package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

//-----------------------------------------------------------------------------
// Public
//-----------------------------------------------------------------------------

var _ = Service("public", func() {
	Description("The company service returns company data")
	//var _ = Resource("public", func() {
	// Sets CORS response headers for requests with Origin header matching the regular expression ".*domain.*"
	cors.Origin("*", func() {
		cors.Methods("GET", "OPTIONS")
	})
	//Origin("*", func() {
	//	Methods("GET", "OPTIONS")
	//})
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
