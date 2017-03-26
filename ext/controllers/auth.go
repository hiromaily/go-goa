package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	"github.com/hiromaily/go-goa/ext/libs/jwt"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	ctx *c.Ctx
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, ctx *c.Ctx) *AuthController {
	return &AuthController{
		Controller: service.NewController("AuthController"),
		ctx:        ctx,
	}
}

// Login runs the login action.
func (c *AuthController) Login(ctx *app.LoginAuthContext) error {
	// AuthController_Login: start_implement
	fmt.Println("[auth.go][Login]", ctx)

	// Login
	svc := &m.User{Db: c.ctx.Db}
	err := svc.Login(ctx.Payload.Username, ctx.Payload.Password)
	if err != nil {
		return ctx.Unauthorized()
	}

	// Generate JWT
	signedToken, err := jwt.GenerateToken(c.ctx)
	if err != nil {
		return ctx.Unauthorized()
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	// AuthController_Login: end_implement
	res := &app.Authorized{Token: signedToken}
	return ctx.OK(res)
}