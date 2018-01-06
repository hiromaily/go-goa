package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Login runs the Login action.
func (c *AuthController) Login(ctx *app.LoginAuthContext) error {
	// AuthController_Login: start_implement

	// Put your logic here

	res := &app.Authorized{}
	return ctx.OK(res)
	// AuthController_Login: end_implement
}
