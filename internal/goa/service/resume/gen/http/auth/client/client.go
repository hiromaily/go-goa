// Code generated by goa v3.10.1, DO NOT EDIT.
//
// auth client HTTP transport
//
// Command:
// $ goa gen resume/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the auth service endpoint HTTP clients.
type Client struct {
	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the auth service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		LoginDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Login returns an endpoint that makes HTTP requests to the auth service login
// server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("auth", "login", err)
		}
		return decodeResponse(resp)
	}
}
