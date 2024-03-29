// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_user endpoints
//
// Command:
// $ goa gen resume/design

package hyuser

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "hy_user" service endpoints.
type Endpoints struct {
	UserList   goa.Endpoint
	GetUser    goa.Endpoint
	CreateUser goa.Endpoint
	UpdateUser goa.Endpoint
	DeleteUser goa.Endpoint
}

// NewEndpoints wraps the methods of the "hy_user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		UserList:   NewUserListEndpoint(s, a.JWTAuth),
		GetUser:    NewGetUserEndpoint(s, a.JWTAuth),
		CreateUser: NewCreateUserEndpoint(s, a.JWTAuth),
		UpdateUser: NewUpdateUserEndpoint(s, a.JWTAuth),
		DeleteUser: NewDeleteUserEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "hy_user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UserList = m(e.UserList)
	e.GetUser = m(e.GetUser)
	e.CreateUser = m(e.CreateUser)
	e.UpdateUser = m(e.UpdateUser)
	e.DeleteUser = m(e.DeleteUser)
}

// NewUserListEndpoint returns an endpoint function that calls the method
// "userList" of service "hy_user".
func NewUserListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UserListPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.UserList(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserCollection(res, view)
		return vres, nil
	}
}

// NewGetUserEndpoint returns an endpoint function that calls the method
// "getUser" of service "hy_user".
func NewGetUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.GetUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUser(res, view)
		return vres, nil
	}
}

// NewCreateUserEndpoint returns an endpoint function that calls the method
// "createUser" of service "hy_user".
func NewCreateUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.CreateUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUser(res, view)
		return vres, nil
	}
}

// NewUpdateUserEndpoint returns an endpoint function that calls the method
// "updateUser" of service "hy_user".
func NewUpdateUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateUser(ctx, p)
	}
}

// NewDeleteUserEndpoint returns an endpoint function that calls the method
// "deleteUser" of service "hy_user".
func NewDeleteUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteUser(ctx, p)
	}
}
