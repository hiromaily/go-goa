// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_companybranch endpoints
//
// Command:
// $ goa gen resume/design

package hycompanybranch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "hy_companybranch" service endpoints.
type Endpoints struct {
	GetCompanyBranch    goa.Endpoint
	CreateCompanyBranch goa.Endpoint
	UpdateCompanyBranch goa.Endpoint
	DeleteCompanyBranch goa.Endpoint
}

// NewEndpoints wraps the methods of the "hy_companybranch" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetCompanyBranch:    NewGetCompanyBranchEndpoint(s, a.JWTAuth),
		CreateCompanyBranch: NewCreateCompanyBranchEndpoint(s, a.JWTAuth),
		UpdateCompanyBranch: NewUpdateCompanyBranchEndpoint(s, a.JWTAuth),
		DeleteCompanyBranch: NewDeleteCompanyBranchEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "hy_companybranch" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetCompanyBranch = m(e.GetCompanyBranch)
	e.CreateCompanyBranch = m(e.CreateCompanyBranch)
	e.UpdateCompanyBranch = m(e.UpdateCompanyBranch)
	e.DeleteCompanyBranch = m(e.DeleteCompanyBranch)
}

// NewGetCompanyBranchEndpoint returns an endpoint function that calls the
// method "getCompanyBranch" of service "hy_companybranch".
func NewGetCompanyBranchEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetCompanyBranchPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.GetCompanyBranch(ctx, p)
	}
}

// NewCreateCompanyBranchEndpoint returns an endpoint function that calls the
// method "createCompanyBranch" of service "hy_companybranch".
func NewCreateCompanyBranchEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateCompanyBranchPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateCompanyBranch(ctx, p)
	}
}

// NewUpdateCompanyBranchEndpoint returns an endpoint function that calls the
// method "updateCompanyBranch" of service "hy_companybranch".
func NewUpdateCompanyBranchEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateCompanyBranchPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateCompanyBranch(ctx, p)
	}
}

// NewDeleteCompanyBranchEndpoint returns an endpoint function that calls the
// method "deleteCompanyBranch" of service "hy_companybranch".
func NewDeleteCompanyBranchEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteCompanyBranchPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access"},
			RequiredScopes: []string{"api:access"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteCompanyBranch(ctx, p)
	}
}
