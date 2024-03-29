// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_userWorkHistory client HTTP transport
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

// Client lists the hy_userWorkHistory service endpoint HTTP clients.
type Client struct {
	// GetUserWorkHistory Doer is the HTTP client used to make requests to the
	// getUserWorkHistory endpoint.
	GetUserWorkHistoryDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the hy_userWorkHistory service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetUserWorkHistoryDoer: doer,
		CORSDoer:               doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// GetUserWorkHistory returns an endpoint that makes HTTP requests to the
// hy_userWorkHistory service getUserWorkHistory server.
func (c *Client) GetUserWorkHistory() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserWorkHistoryRequest(c.encoder)
		decodeResponse = DecodeGetUserWorkHistoryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserWorkHistoryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserWorkHistoryDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hy_userWorkHistory", "getUserWorkHistory", err)
		}
		return decodeResponse(resp)
	}
}
