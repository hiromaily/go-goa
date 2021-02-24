package service

import (
	"context"
	"fmt"
	"log"
	hyusertech "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_usertech"

	"goa.design/goa/v3/security"
)

// hy_usertech service example implementation.
// The example methods log the requests and return zero values.
type hyUsertechsrvc struct {
	logger *log.Logger
}

// NewHyUsertech returns the hy_usertech service implementation.
func NewHyUsertech(logger *log.Logger) hyusertech.Service {
	return &hyUsertechsrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_usertech" for the
// "jwt" security scheme.
func (s *hyUsertechsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// get user's favorite techs
func (s *hyUsertechsrvc) GetUserLikeTech(ctx context.Context, p *hyusertech.GetUserLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	view = "default"
	s.logger.Print("hyUsertech.getUserLikeTech")
	return
}

// get user's dislike techs
func (s *hyUsertechsrvc) GetUserDisLikeTech(ctx context.Context, p *hyusertech.GetUserDisLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	view = "default"
	s.logger.Print("hyUsertech.getUserDisLikeTech")
	return
}
