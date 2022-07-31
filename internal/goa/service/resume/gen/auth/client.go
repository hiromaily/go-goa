// Code generated by goa v3.7.13, DO NOT EDIT.
//
// auth client
//
// Command:
// $ goa gen resume/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "auth" service client.
type Client struct {
	LoginEndpoint goa.Endpoint
}

// NewClient initializes a "auth" service client given the endpoints.
func NewClient(login goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint: login,
	}
}

// Login calls the "login" endpoint of the "auth" service.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *Authorized, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Authorized), nil
}
