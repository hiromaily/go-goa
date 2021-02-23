package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("health", func() {
	Description("health check")

	Method("health", func() { //just name
		NoSecurity()
		HTTP(func() {
			GET("/health")
			Response(StatusOK)
		})
	})
})

//var _ = Resource("health", func() {
//
//	BasePath("/_ah")
//
//	Action("health", func() {
//		Routing(
//			GET("/health"),
//		)
//		Description("Perform health check.")
//		Response(OK, "text/plain")
//	})
//})
