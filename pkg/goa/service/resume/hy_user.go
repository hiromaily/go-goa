package resumeapi

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
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

// List all users
func (s *hyUsersrvc) UserList(ctx context.Context, p *hyuser.UserListPayload) (res hyuser.UserCollection, view string, err error) {
	log.Info().Msg("hyUser.UserList")

	users, err := s.userRepo.UserList()
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call userRepo.UserList()")
	}
	if len(users) == 0 {
		return nil, "", hyuser.MakeNoContent(errors.New("no user"))
	}
	res = users

	view = "default"
	return
}

// get user with given user id
func (s *hyUsersrvc) GetUser(ctx context.Context, p *hyuser.GetUserPayload) (res *hyuser.User, view string, err error) {
	//	user := &app.User{}
	//
	//	svc := &m.User{Db: c.ctx.Db}
	//	err := svc.GetUser(ctx.UserID, user)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if user.ID == nil {
	//		//404
	//		return ctx.NotFound()
	//	}
	//
	//	//res := &app.User{}
	//	return ctx.OK(user)
	res = &hyuser.User{}
	view = "default"
	log.Info().Msg("hyUser.getUser")
	return
}

// Create new user
func (s *hyUsersrvc) CreateUser(ctx context.Context, p *hyuser.CreateUserPayload) (err error) {
	//	svc := &m.User{Db: c.ctx.Db}
	//	userID, err := svc.InsertUser(ctx.Payload) //*CreateUserHyUserPayload
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.UserID{ID: &userID}
	//	return ctx.OKId(res)
	log.Info().Msg("hyUser.createUser")
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
