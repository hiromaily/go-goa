package resumeapi

import (
	"context"
	"log"
	auth "resume/gen/auth"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	logger *log.Logger
}

// NewAuth returns the auth service implementation.
func NewAuth(logger *log.Logger) auth.Service {
	return &authsrvc{logger}
}

// Login implements login.
func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.Authorized, err error) {
	//	// Login
	//	svc := &m.User{Db: c.ctx.Db}
	//	id, err := svc.Login(ctx.Payload.Email, ctx.Payload.Password)
	//	if err != nil {
	//		//return ctx.Unauthorized()
	//		return err
	//	} else if id == 0 {
	//		return ctx.Unauthorized()
	//	}
	//
	//	// Generate JWT
	//	signedToken, err := jwt.GenerateToken(c.ctx, id)
	//	if err != nil {
	//		return ctx.Unauthorized()
	//	}
	//
	//	// Set auth header for client retrieval
	//	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	//
	//	// Send response
	//	// AuthController_Login: end_implement
	//	res := &app.Authorized{Token: signedToken, ID: id}
	//	return ctx.OK(res)

	res = &auth.Authorized{}
	s.logger.Print("auth.login")
	return
}
