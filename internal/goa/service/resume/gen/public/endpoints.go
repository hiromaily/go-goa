// Code generated by goa v3.2.6, DO NOT EDIT.
//
// public endpoints
//
// Command:
// $ goa gen resume/design

package public

import (
	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "public" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "public" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "public" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}