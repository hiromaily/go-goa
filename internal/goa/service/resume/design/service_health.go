package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("health", func() {
	Description("health check")

	HTTP(func() {
		// BasePath
		Path(baseAPIDir)
	})

	Method("health", func() {
		NoSecurity()
		HTTP(func() {
			GET("/health")
			Response(StatusOK)
		})
	})
})
