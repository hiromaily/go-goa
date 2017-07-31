package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Additional file for health check
//
var _ = Resource("health", func() {

	BasePath("/_ah")

	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})
