//go:generate goagen bootstrap -d github.com/hiromaily/go-goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hiromaily/go-goa/app"
)

func main() {
	// Create service
	service := goa.New("api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "health" controller
	c := NewHealthController(service)
	app.MountHealthController(service, c)
	// Mount "hy_company" controller
	c2 := NewHyCompanyController(service)
	app.MountHyCompanyController(service, c2)
	// Mount "hy_user" controller
	c3 := NewHyUserController(service)
	app.MountHyUserController(service, c3)
	// Mount "js" controller
	c4 := NewJsController(service)
	app.MountJsController(service, c4)
	// Mount "public" controller
	c5 := NewPublicController(service)
	app.MountPublicController(service, c5)
	// Mount "swagger" controller
	c6 := NewSwaggerController(service)
	app.MountSwaggerController(service, c6)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
