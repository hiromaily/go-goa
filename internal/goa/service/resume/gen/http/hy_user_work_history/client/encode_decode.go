// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_userWorkHistory HTTP client encoders and decoders
//
// Command:
// $ goa gen resume/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	hyuserworkhistory "resume/gen/hy_user_work_history"
	hyuserworkhistoryviews "resume/gen/hy_user_work_history/views"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetUserWorkHistoryRequest instantiates a HTTP request object with
// method and path set to call the "hy_userWorkHistory" service
// "getUserWorkHistory" endpoint
func (c *Client) BuildGetUserWorkHistoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*hyuserworkhistory.GetUserWorkHistoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_userWorkHistory", "getUserWorkHistory", "*hyuserworkhistory.GetUserWorkHistoryPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserWorkHistoryHyUserWorkHistoryPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_userWorkHistory", "getUserWorkHistory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserWorkHistoryRequest returns an encoder for requests sent to the
// hy_userWorkHistory getUserWorkHistory server.
func EncodeGetUserWorkHistoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuserworkhistory.GetUserWorkHistoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_userWorkHistory", "getUserWorkHistory", "*hyuserworkhistory.GetUserWorkHistoryPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeGetUserWorkHistoryResponse returns a decoder for responses returned by
// the hy_userWorkHistory getUserWorkHistory endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeGetUserWorkHistoryResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetUserWorkHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserWorkHistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_userWorkHistory", "getUserWorkHistory", err)
			}
			p := NewGetUserWorkHistoryUserworkhistoryCollectionOK(body)
			view := "default"
			vres := hyuserworkhistoryviews.UserworkhistoryCollection{Projected: p, View: view}
			if err = hyuserworkhistoryviews.ValidateUserworkhistoryCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_userWorkHistory", "getUserWorkHistory", err)
			}
			res := hyuserworkhistory.NewUserworkhistoryCollection(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetUserWorkHistoryNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_userWorkHistory", "getUserWorkHistory", err)
			}
			err = ValidateGetUserWorkHistoryNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_userWorkHistory", "getUserWorkHistory", err)
			}
			return nil, NewGetUserWorkHistoryNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_userWorkHistory", "getUserWorkHistory", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserworkhistoryResponseToHyuserworkhistoryviewsUserworkhistoryView
// builds a value of type *hyuserworkhistoryviews.UserworkhistoryView from a
// value of type *UserworkhistoryResponse.
func unmarshalUserworkhistoryResponseToHyuserworkhistoryviewsUserworkhistoryView(v *UserworkhistoryResponse) *hyuserworkhistoryviews.UserworkhistoryView {
	res := &hyuserworkhistoryviews.UserworkhistoryView{
		Title:       v.Title,
		CompanyName: v.CompanyName,
		CountryName: v.CountryName,
		Term:        v.Term,
		Description: v.Description,
		Techs:       v.Techs,
	}

	return res
}
