// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_company client
//
// Command:
// $ goa gen resume/design

package hycompany

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hy_company" service client.
type Client struct {
	CompanyListEndpoint     goa.Endpoint
	GetCompanyGroupEndpoint goa.Endpoint
	CreateCompanyEndpoint   goa.Endpoint
	UpdateCompanyEndpoint   goa.Endpoint
	DeleteCompanyEndpoint   goa.Endpoint
}

// NewClient initializes a "hy_company" service client given the endpoints.
func NewClient(companyList, getCompanyGroup, createCompany, updateCompany, deleteCompany goa.Endpoint) *Client {
	return &Client{
		CompanyListEndpoint:     companyList,
		GetCompanyGroupEndpoint: getCompanyGroup,
		CreateCompanyEndpoint:   createCompany,
		UpdateCompanyEndpoint:   updateCompany,
		DeleteCompanyEndpoint:   deleteCompany,
	}
}

// CompanyList calls the "companyList" endpoint of the "hy_company" service.
// CompanyList may return the following errors:
//	- "NoContent" (type *goa.ServiceError)
//	- "BadRequest" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) CompanyList(ctx context.Context, p *CompanyListPayload) (res CompanyCollection, err error) {
	var ires interface{}
	ires, err = c.CompanyListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(CompanyCollection), nil
}

// GetCompanyGroup calls the "getCompanyGroup" endpoint of the "hy_company"
// service.
// GetCompanyGroup may return the following errors:
//	- "NoContent" (type *goa.ServiceError)
//	- "BadRequest" (type *goa.ServiceError)
//	- "NotFound" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) GetCompanyGroup(ctx context.Context, p *GetCompanyGroupPayload) (res CompanyCollection, err error) {
	var ires interface{}
	ires, err = c.GetCompanyGroupEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(CompanyCollection), nil
}

// CreateCompany calls the "createCompany" endpoint of the "hy_company" service.
// CreateCompany may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) CreateCompany(ctx context.Context, p *CreateCompanyPayload) (err error) {
	_, err = c.CreateCompanyEndpoint(ctx, p)
	return
}

// UpdateCompany calls the "updateCompany" endpoint of the "hy_company" service.
// UpdateCompany may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- "NotFound" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) UpdateCompany(ctx context.Context, p *UpdateCompanyPayload) (err error) {
	_, err = c.UpdateCompanyEndpoint(ctx, p)
	return
}

// DeleteCompany calls the "deleteCompany" endpoint of the "hy_company" service.
// DeleteCompany may return the following errors:
//	- "BadRequest" (type *goa.ServiceError)
//	- "NotFound" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) DeleteCompany(ctx context.Context, p *DeleteCompanyPayload) (err error) {
	_, err = c.DeleteCompanyEndpoint(ctx, p)
	return
}