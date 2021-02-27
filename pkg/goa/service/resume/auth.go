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
	res = &auth.Authorized{}
	s.logger.Print("auth.login")
	return
}
