// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_company HTTP client encoders and decoders
//
// Command:
// $ goa gen resume/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	hycompany "resume/gen/hy_company"
	hycompanyviews "resume/gen/hy_company/views"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildCompanyListRequest instantiates a HTTP request object with method and
// path set to call the "hy_company" service "companyList" endpoint
func (c *Client) BuildCompanyListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CompanyListHyCompanyPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_company", "companyList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCompanyListRequest returns an encoder for requests sent to the
// hy_company companyList server.
func EncodeCompanyListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hycompany.CompanyListPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_company", "companyList", "*hycompany.CompanyListPayload", v)
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

// DecodeCompanyListResponse returns a decoder for responses returned by the
// hy_company companyList endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCompanyListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CompanyListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_company", "companyList", err)
			}
			p := NewCompanyListCompanyCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := hycompanyviews.CompanyCollection{Projected: p, View: view}
			if err = hycompanyviews.ValidateCompanyCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_company", "companyList", err)
			}
			res := hycompany.NewCompanyCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_company", "companyList", resp.StatusCode, string(body))
		}
	}
}

// BuildGetCompanyGroupRequest instantiates a HTTP request object with method
// and path set to call the "hy_company" service "getCompanyGroup" endpoint
func (c *Client) BuildGetCompanyGroupRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		companyID int
	)
	{
		p, ok := v.(*hycompany.GetCompanyGroupPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_company", "getCompanyGroup", "*hycompany.GetCompanyGroupPayload", v)
		}
		if p.CompanyID != nil {
			companyID = *p.CompanyID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCompanyGroupHyCompanyPath(companyID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_company", "getCompanyGroup", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetCompanyGroupRequest returns an encoder for requests sent to the
// hy_company getCompanyGroup server.
func EncodeGetCompanyGroupRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hycompany.GetCompanyGroupPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_company", "getCompanyGroup", "*hycompany.GetCompanyGroupPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewGetCompanyGroupRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("hy_company", "getCompanyGroup", err)
		}
		return nil
	}
}

// DecodeGetCompanyGroupResponse returns a decoder for responses returned by
// the hy_company getCompanyGroup endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetCompanyGroupResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetCompanyGroupResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_company", "getCompanyGroup", err)
			}
			p := NewGetCompanyGroupCompanyCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := hycompanyviews.CompanyCollection{Projected: p, View: view}
			if err = hycompanyviews.ValidateCompanyCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_company", "getCompanyGroup", err)
			}
			res := hycompany.NewCompanyCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_company", "getCompanyGroup", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateCompanyRequest instantiates a HTTP request object with method and
// path set to call the "hy_company" service "createCompany" endpoint
func (c *Client) BuildCreateCompanyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateCompanyHyCompanyPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_company", "createCompany", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateCompanyRequest returns an encoder for requests sent to the
// hy_company createCompany server.
func EncodeCreateCompanyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hycompany.CreateCompanyPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_company", "createCompany", "*hycompany.CreateCompanyPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateCompanyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("hy_company", "createCompany", err)
		}
		return nil
	}
}

// DecodeCreateCompanyResponse returns a decoder for responses returned by the
// hy_company createCompany endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCreateCompanyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_company", "createCompany", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateCompanyRequest instantiates a HTTP request object with method and
// path set to call the "hy_company" service "updateCompany" endpoint
func (c *Client) BuildUpdateCompanyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		companyID int
	)
	{
		p, ok := v.(*hycompany.UpdateCompanyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_company", "updateCompany", "*hycompany.UpdateCompanyPayload", v)
		}
		if p.CompanyID != nil {
			companyID = *p.CompanyID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateCompanyHyCompanyPath(companyID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_company", "updateCompany", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateCompanyRequest returns an encoder for requests sent to the
// hy_company updateCompany server.
func EncodeUpdateCompanyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hycompany.UpdateCompanyPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_company", "updateCompany", "*hycompany.UpdateCompanyPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateCompanyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("hy_company", "updateCompany", err)
		}
		return nil
	}
}

// DecodeUpdateCompanyResponse returns a decoder for responses returned by the
// hy_company updateCompany endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeUpdateCompanyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_company", "updateCompany", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteCompanyRequest instantiates a HTTP request object with method and
// path set to call the "hy_company" service "deleteCompany" endpoint
func (c *Client) BuildDeleteCompanyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		companyID int
	)
	{
		p, ok := v.(*hycompany.DeleteCompanyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_company", "deleteCompany", "*hycompany.DeleteCompanyPayload", v)
		}
		if p.CompanyID != nil {
			companyID = *p.CompanyID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteCompanyHyCompanyPath(companyID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_company", "deleteCompany", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteCompanyRequest returns an encoder for requests sent to the
// hy_company deleteCompany server.
func EncodeDeleteCompanyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hycompany.DeleteCompanyPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_company", "deleteCompany", "*hycompany.DeleteCompanyPayload", v)
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

// DecodeDeleteCompanyResponse returns a decoder for responses returned by the
// hy_company deleteCompany endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDeleteCompanyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_company", "deleteCompany", resp.StatusCode, string(body))
		}
	}
}

// unmarshalCompanyResponseToHycompanyviewsCompanyView builds a value of type
// *hycompanyviews.CompanyView from a value of type *CompanyResponse.
func unmarshalCompanyResponseToHycompanyviewsCompanyView(v *CompanyResponse) *hycompanyviews.CompanyView {
	res := &hycompanyviews.CompanyView{
		ID:          v.ID,
		CompanyID:   v.CompanyID,
		Name:        v.Name,
		IsHq:        v.IsHq,
		CountryName: v.CountryName,
		Address:     v.Address,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	return res
}
