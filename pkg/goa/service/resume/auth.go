package resumeapi

import (
	"context"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"

	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/repository"
	"github.com/pkg/errors"
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

func (s *authsrvc) Login(ctx context.Context, p *auth.LoginPayload) (res *auth.Authorized, err error) {
	log.Info().Msg("auth.Login()")

	userID, err := s.userRepo.Login(p.Email, p.Password)
	if err != nil {
		// FIXME: why filed:fault has `fault`??
		return nil, auth.MakeUnauthorized(errors.New("login error"))
	}

	// Generate JWT
	signedToken, err := s.jwt.CreateBasicToken(time.Now().Add(time.Minute*60).Unix(), strconv.Itoa(userID), p.Email)
	if err != nil {
		return nil, auth.MakeUnauthorized(errors.New("fail to generate jwt"))
	}

	// TODO: Set auth header for client retrieval
	//ctx..ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Response
	res = &auth.Authorized{
		Token: signedToken,
		ID:    userID,
	}

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
	return
}
