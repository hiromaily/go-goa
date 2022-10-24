package resumeapi

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/repository"
	auth "resume/gen/auth"
)

// auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct {
	jwt      jwts.JWTer
	userRepo repository.UserRepository
}

// NewAuth returns the auth service implementation.
func NewAuth(jwt jwts.JWTer, userRepo repository.UserRepository) auth.Service {
	return &authsrvc{jwt, userRepo}
}

// TODO: how to handle user has arleady logged in?
func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.Authorized, err error) {
	log.Info().Msg("auth.Login()")

	userID, err := s.userRepo.Login(p.Email, p.Password)
	if err != nil {
		return nil, auth.MakeUnauthorized(errors.New("login error"))
	}

	// Generate JWT
	signedToken, err := s.jwt.CreateBasicToken(time.Now().Add(time.Minute*60).Unix(), strconv.Itoa(userID), p.Email)
	if err != nil {
		return nil, auth.MakeUnauthorized(errors.New("fail to generate jwt"))
	}

	// Response
	res = &auth.Authorized{
		Token: signedToken,
		ID:    userID,
	}
	return
}
