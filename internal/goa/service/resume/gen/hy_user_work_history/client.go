// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_userWorkHistory client
//
// Command:
// $ goa gen resume/design

package hyuserworkhistory

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hy_userWorkHistory" service client.
type Client struct {
	GetUserWorkHistoryEndpoint goa.Endpoint
}

// NewClient initializes a "hy_userWorkHistory" service client given the
// endpoints.
func NewClient(getUserWorkHistory goa.Endpoint) *Client {
	return &Client{
		GetUserWorkHistoryEndpoint: getUserWorkHistory,
	}
}

// GetUserWorkHistory calls the "getUserWorkHistory" endpoint of the
// "hy_userWorkHistory" service.
// GetUserWorkHistory may return the following errors:
//   - "NotFound" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetUserWorkHistory(ctx context.Context, p *GetUserWorkHistoryPayload) (res UserworkhistoryCollection, err error) {
	var ires interface{}
	ires, err = c.GetUserWorkHistoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(UserworkhistoryCollection), nil
}
