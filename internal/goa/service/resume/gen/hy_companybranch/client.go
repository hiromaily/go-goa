// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_companybranch client
//
// Command:
// $ goa gen resume/design

package hycompanybranch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hy_companybranch" service client.
type Client struct {
	GetCompanyBranchEndpoint    goa.Endpoint
	CreateCompanyBranchEndpoint goa.Endpoint
	UpdateCompanyBranchEndpoint goa.Endpoint
	DeleteCompanyBranchEndpoint goa.Endpoint
}

// NewClient initializes a "hy_companybranch" service client given the
// endpoints.
func NewClient(getCompanyBranch, createCompanyBranch, updateCompanyBranch, deleteCompanyBranch goa.Endpoint) *Client {
	return &Client{
		GetCompanyBranchEndpoint:    getCompanyBranch,
		CreateCompanyBranchEndpoint: createCompanyBranch,
		UpdateCompanyBranchEndpoint: updateCompanyBranch,
		DeleteCompanyBranchEndpoint: deleteCompanyBranch,
	}
}

// GetCompanyBranch calls the "getCompanyBranch" endpoint of the
// "hy_companybranch" service.
// GetCompanyBranch may return the following errors:
//	- "NotFound" (type *goa.ServiceError)
//	- "BadRequest" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) GetCompanyBranch(ctx context.Context, p *GetCompanyBranchPayload) (err error) {
	_, err = c.GetCompanyBranchEndpoint(ctx, p)
	return
}

// CreateCompanyBranch calls the "createCompanyBranch" endpoint of the
// "hy_companybranch" service.
// CreateCompanyBranch may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) CreateCompanyBranch(ctx context.Context, p *CreateCompanyBranchPayload) (err error) {
	_, err = c.CreateCompanyBranchEndpoint(ctx, p)
	return
}

// UpdateCompanyBranch calls the "updateCompanyBranch" endpoint of the
// "hy_companybranch" service.
// UpdateCompanyBranch may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- "NotFound" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) UpdateCompanyBranch(ctx context.Context, p *UpdateCompanyBranchPayload) (err error) {
	_, err = c.UpdateCompanyBranchEndpoint(ctx, p)
	return
}

// DeleteCompanyBranch calls the "deleteCompanyBranch" endpoint of the
// "hy_companybranch" service.
// DeleteCompanyBranch may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- "NotFound" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) DeleteCompanyBranch(ctx context.Context, p *DeleteCompanyBranchPayload) (err error) {
	_, err = c.DeleteCompanyBranchEndpoint(ctx, p)
	return
}
