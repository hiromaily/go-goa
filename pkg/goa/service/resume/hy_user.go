package resumeapi

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/pointer"
	"github.com/hiromaily/go-goa/pkg/repository"
	hyuser "resume/gen/hy_user"
)

// hy_user service example implementation.
// The example methods log the requests and return zero values.
type hyUsersrvc struct {
	jwt      jwts.JWTer
	userRepo repository.UserRepository
}

// NewHyUser returns the hy_user service implementation.
func NewHyUser(jwt jwts.JWTer, userRepo repository.UserRepository) hyuser.Service {
	return &hyUsersrvc{jwt, userRepo}
}

// JWTAuth implements the authorization logic for service "hy_user" for the
// "jwt" security scheme.
func (s *hyUsersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	log.Info().
		Str("token", token).
		Strs("scheme.Scopes", scheme.Scopes).
		Msg("hyUser.JWTAuth")

	if err := s.jwt.ValidateToken(token); err != nil {
		return ctx, err
	}

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
	return ctx, nil
}

// UserList returns all users
func (s *hyUsersrvc) UserList(ctx context.Context, p *hyuser.UserListPayload) (res hyuser.UserCollection, view string, err error) {
	log.Info().Msg("hyUser.UserList")

	users, err := s.userRepo.UserList()
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call userRepo.UserList()")
	}
	if len(users) == 0 {
		return nil, "", hyuser.MakeNotFound(errors.New("user not found"))
	}
	res = users
	view = "default"

	return
}

// GetUser returns user by given UserID
func (s *hyUsersrvc) GetUser(ctx context.Context, p *hyuser.GetUserPayload) (res *hyuser.User, view string, err error) {
	log.Info().Msg("hyUser.getUser")

	user, err := s.userRepo.GetUser(p.UserID)
	if err != nil {
		// Not Found
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, "", hyuser.MakeNotFound(errors.New("user not found"))
		}
		return nil, "", errors.Wrapf(err, "fail to call userRepo.GetUser(%d)", p.UserID)
	}
	res = user
	view = "default"

	return
}

// CreateUser creates new user
func (s *hyUsersrvc) CreateUser(ctx context.Context, p *hyuser.CreateUserPayload) (res *hyuser.User, view string, err error) {
	log.Info().Msg("hyUser.createUser")

	userID, err := s.userRepo.InsertUser(p.UserName, p.Email, p.Password)
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call InsertUser()")
	}
	res = &hyuser.User{
		ID: pointer.Int(userID),
	}
	view = "id"

	return
}

// Change user properties
func (s *hyUsersrvc) UpdateUser(ctx context.Context, p *hyuser.UpdateUserPayload) (err error) {
	//	svc := &m.User{Db: c.ctx.Db}
	//	err := svc.UpdateUser(ctx.UserID, ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.UserID{ID: &ctx.UserID}
	//	return ctx.OKId(res)
	log.Info().Msg("hyUser.updateUser")
	return
}

// Delete user
func (s *hyUsersrvc) DeleteUser(ctx context.Context, p *hyuser.DeleteUserPayload) (err error) {
	//	svc := &m.User{Db: c.ctx.Db}
	//	err := svc.DeleteUser(ctx.UserID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.UserID{ID: &ctx.UserID}
	//	return ctx.OKId(res)
	log.Info().Msg("hyUser.deleteUser")
	return
}
