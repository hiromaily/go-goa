// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_userWorkHistory endpoints
//
// Command:
// $ goa gen resume/design

package hyuserworkhistory

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "hy_userWorkHistory" service endpoints.
type Endpoints struct {
	GetUserWorkHistory goa.Endpoint
}

// NewEndpoints wraps the methods of the "hy_userWorkHistory" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetUserWorkHistory: NewGetUserWorkHistoryEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "hy_userWorkHistory" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetUserWorkHistory = m(e.GetUserWorkHistory)
}

// NewGetUserWorkHistoryEndpoint returns an endpoint function that calls the
// method "getUserWorkHistory" of service "hy_userWorkHistory".
func NewGetUserWorkHistoryEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserWorkHistoryPayload)
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
		res, err := s.GetUserWorkHistory(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserworkhistoryCollection(res, "default")
		return vres, nil
	}
}
