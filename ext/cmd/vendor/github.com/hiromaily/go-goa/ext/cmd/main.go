//go:generate goagen bootstrap -d github.com/hiromaily/go-goa/goa/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hiromaily/go-goa/goa/app"
	g "github.com/hiromaily/go-goa/goa"

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
	c := g.NewHealthController(service)
	app.MountHealthController(service, c)
	// Mount "hy_company" controller
	c2 := g.NewHyCompanyController(service)
	app.MountHyCompanyController(service, c2)
	// Mount "hy_user" controller
	c3 := g.NewHyUserController(service)
	app.MountHyUserController(service, c3)
	// Mount "public" controller
	c4 := g.NewPublicController(service)
	app.MountPublicController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
