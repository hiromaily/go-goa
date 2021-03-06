package resumeapi

import (
	"context"
	"fmt"
	"log"
	hyuser "resume/gen/hy_user"

	"goa.design/goa/v3/security"
)

// hy_user service example implementation.
// The example methods log the requests and return zero values.
type hyUsersrvc struct {
	logger *log.Logger
}

// NewHyUser returns the hy_user service implementation.
func NewHyUser(logger *log.Logger) hyuser.Service {
	return &hyUsersrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_user" for the
// "jwt" security scheme.
func (s *hyUsersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// List all users
func (s *hyUsersrvc) UserList(ctx context.Context, p *hyuser.UserListPayload) (res hyuser.UserCollection, view string, err error) {
	//	//type User struct {
	//	//	Email     string `form:"email" json:"email" xml:"email"`
	//	//	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	//	//	// User ID
	//	//	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	//	//}
	//	var users []*app.User
	//
	//	svc := &m.User{Db: c.ctx.Db}
	//	err := svc.UserList(&users)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(users) == 0 {
	//		return ctx.NoContent()
	//	}
	//
	//	//type UserCollection []*User
	//	res := app.UserCollection(users)
	//	return ctx.OK(res)

	view = "default"
	s.logger.Print("hyUser.userList")
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
	s.logger.Print("hyUser.getUser")
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

	s.logger.Print("hyUser.createUser")
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

	s.logger.Print("hyUser.updateUser")
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

	s.logger.Print("hyUser.deleteUser")
	return
}
