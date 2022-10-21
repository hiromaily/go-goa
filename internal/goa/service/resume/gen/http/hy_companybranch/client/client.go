// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_companybranch client HTTP transport
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

// Client lists the hy_companybranch service endpoint HTTP clients.
type Client struct {
	// GetCompanyBranch Doer is the HTTP client used to make requests to the
	// getCompanyBranch endpoint.
	GetCompanyBranchDoer goahttp.Doer

	// CreateCompanyBranch Doer is the HTTP client used to make requests to the
	// createCompanyBranch endpoint.
	CreateCompanyBranchDoer goahttp.Doer

	// UpdateCompanyBranch Doer is the HTTP client used to make requests to the
	// updateCompanyBranch endpoint.
	UpdateCompanyBranchDoer goahttp.Doer

	// DeleteCompanyBranch Doer is the HTTP client used to make requests to the
	// deleteCompanyBranch endpoint.
	DeleteCompanyBranchDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the hy_companybranch service
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
		GetCompanyBranchDoer:    doer,
		CreateCompanyBranchDoer: doer,
		UpdateCompanyBranchDoer: doer,
		DeleteCompanyBranchDoer: doer,
		CORSDoer:                doer,
		RestoreResponseBody:     restoreBody,
		scheme:                  scheme,
		host:                    host,
		decoder:                 dec,
		encoder:                 enc,
	}
}

// GetCompanyBranch returns an endpoint that makes HTTP requests to the
// hy_companybranch service getCompanyBranch server.
func (c *Client) GetCompanyBranch() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetCompanyBranchRequest(c.encoder)
		decodeResponse = DecodeGetCompanyBranchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCompanyBranchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCompanyBranchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hy_companybranch", "getCompanyBranch", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCompanyBranch returns an endpoint that makes HTTP requests to the
// hy_companybranch service createCompanyBranch server.
func (c *Client) CreateCompanyBranch() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCompanyBranchRequest(c.encoder)
		decodeResponse = DecodeCreateCompanyBranchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCompanyBranchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCompanyBranchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hy_companybranch", "createCompanyBranch", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateCompanyBranch returns an endpoint that makes HTTP requests to the
// hy_companybranch service updateCompanyBranch server.
func (c *Client) UpdateCompanyBranch() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateCompanyBranchRequest(c.encoder)
		decodeResponse = DecodeUpdateCompanyBranchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateCompanyBranchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateCompanyBranchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hy_companybranch", "updateCompanyBranch", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCompanyBranch returns an endpoint that makes HTTP requests to the
// hy_companybranch service deleteCompanyBranch server.
func (c *Client) DeleteCompanyBranch() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteCompanyBranchRequest(c.encoder)
		decodeResponse = DecodeDeleteCompanyBranchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCompanyBranchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCompanyBranchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hy_companybranch", "deleteCompanyBranch", err)
		}
		return decodeResponse(resp)
	}
}
