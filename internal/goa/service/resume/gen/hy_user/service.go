// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_user service
//
// Command:
// $ goa gen resume/design

package hyuser

import (
	"context"
	hyuserviews "resume/gen/hy_user/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The user service returns user data
type Service interface {
	// List all users
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "id"
	UserList(context.Context, *UserListPayload) (res UserCollection, view string, err error)
	// Get user by given user id
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "id"
	GetUser(context.Context, *GetUserPayload) (res *User, view string, err error)
	// Create new user
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "id"
	CreateUser(context.Context, *CreateUserPayload) (res *User, view string, err error)
	// Update user data
	UpdateUser(context.Context, *UpdateUserPayload) (err error)
	// Delete user
	DeleteUser(context.Context, *DeleteUserPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hy_user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"userList", "getUser", "createUser", "updateUser", "deleteUser"}

// CreateUserPayload is the payload type of the hy_user service createUser
// method.
type CreateUserPayload struct {
	// JWT token used to perform authorization
	Token *string
	// User name
	UserName string
	// E-mail of user
	Email string
	// Password
	Password string
}

// DeleteUserPayload is the payload type of the hy_user service deleteUser
// method.
type DeleteUserPayload struct {
	// JWT token used to perform authorization
	Token *string
	// User ID
	UserID int
}

// GetUserPayload is the payload type of the hy_user service getUser method.
type GetUserPayload struct {
	// JWT token used to perform authorization
	Token *string
	// User ID
	UserID int
}

// UpdateUserPayload is the payload type of the hy_user service updateUser
// method.
type UpdateUserPayload struct {
	// JWT token used to perform authorization
	Token *string
	// User ID
	UserID int
	// User name
	UserName *string
	// E-mail of user
	Email *string
	// Password
	Password *string
}

// User is the result type of the hy_user service getUser method.
type User struct {
	// ID
	ID *int
	// User name
	UserName *string
	// E-mail of user
	Email *string
	// Password
	Password *string
	// Datetime
	CreatedAt *string
	// Datetime
	UpdatedAt *string
}

// UserCollection is the result type of the hy_user service userList method.
type UserCollection []*User

// UserListPayload is the payload type of the hy_user service userList method.
type UserListPayload struct {
	// JWT token used to perform authorization
	Token *string
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "NotFound", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "BadRequest", false, false, false)
}

// NewUserCollection initializes result type UserCollection from viewed result
// type UserCollection.
func NewUserCollection(vres hyuserviews.UserCollection) UserCollection {
	var res UserCollection
	switch vres.View {
	case "default", "":
		res = newUserCollection(vres.Projected)
	case "id":
		res = newUserCollectionID(vres.Projected)
	}
	return res
}

// NewViewedUserCollection initializes viewed result type UserCollection from
// result type UserCollection using the given view.
func NewViewedUserCollection(res UserCollection, view string) hyuserviews.UserCollection {
	var vres hyuserviews.UserCollection
	switch view {
	case "default", "":
		p := newUserCollectionView(res)
		vres = hyuserviews.UserCollection{Projected: p, View: "default"}
	case "id":
		p := newUserCollectionViewID(res)
		vres = hyuserviews.UserCollection{Projected: p, View: "id"}
	}
	return vres
}

// NewUser initializes result type User from viewed result type User.
func NewUser(vres *hyuserviews.User) *User {
	var res *User
	switch vres.View {
	case "default", "":
		res = newUser(vres.Projected)
	case "id":
		res = newUserID(vres.Projected)
	}
	return res
}

// NewViewedUser initializes viewed result type User from result type User
// using the given view.
func NewViewedUser(res *User, view string) *hyuserviews.User {
	var vres *hyuserviews.User
	switch view {
	case "default", "":
		p := newUserView(res)
		vres = &hyuserviews.User{Projected: p, View: "default"}
	case "id":
		p := newUserViewID(res)
		vres = &hyuserviews.User{Projected: p, View: "id"}
	}
	return vres
}

// newUserCollection converts projected type UserCollection to service type
// UserCollection.
func newUserCollection(vres hyuserviews.UserCollectionView) UserCollection {
	res := make(UserCollection, len(vres))
	for i, n := range vres {
		res[i] = newUser(n)
	}
	return res
}

// newUserCollectionID converts projected type UserCollection to service type
// UserCollection.
func newUserCollectionID(vres hyuserviews.UserCollectionView) UserCollection {
	res := make(UserCollection, len(vres))
	for i, n := range vres {
		res[i] = newUserID(n)
	}
	return res
}

// newUserCollectionView projects result type UserCollection to projected type
// UserCollectionView using the "default" view.
func newUserCollectionView(res UserCollection) hyuserviews.UserCollectionView {
	vres := make(hyuserviews.UserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newUserView(n)
	}
	return vres
}

// newUserCollectionViewID projects result type UserCollection to projected
// type UserCollectionView using the "id" view.
func newUserCollectionViewID(res UserCollection) hyuserviews.UserCollectionView {
	vres := make(hyuserviews.UserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newUserViewID(n)
	}
	return vres
}

// newUser converts projected type User to service type User.
func newUser(vres *hyuserviews.UserView) *User {
	res := &User{
		ID:       vres.ID,
		UserName: vres.UserName,
		Email:    vres.Email,
	}
	return res
}

// newUserID converts projected type User to service type User.
func newUserID(vres *hyuserviews.UserView) *User {
	res := &User{
		ID: vres.ID,
	}
	return res
}

// newUserView projects result type User to projected type UserView using the
// "default" view.
func newUserView(res *User) *hyuserviews.UserView {
	vres := &hyuserviews.UserView{
		ID:       res.ID,
		UserName: res.UserName,
		Email:    res.Email,
	}
	return vres
}

// newUserViewID projects result type User to projected type UserView using the
// "id" view.
func newUserViewID(res *User) *hyuserviews.UserView {
	vres := &hyuserviews.UserView{
		ID: res.ID,
	}
	return vres
}
