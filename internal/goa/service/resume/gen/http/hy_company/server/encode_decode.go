// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_company HTTP server encoders and decoders
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"io"
	"net/http"
	hycompanyviews "resume/gen/hy_company/views"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCompanyListResponse returns an encoder for responses returned by the
// hy_company companyList endpoint.
func EncodeCompanyListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(hycompanyviews.CompanyCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCompanyResponseCollection(res.Projected)
		case "detailid":
			body = NewCompanyResponseDetailidCollection(res.Projected)
		case "id":
			body = NewCompanyResponseIDCollection(res.Projected)
		case "idname":
			body = NewCompanyResponseIdnameCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCompanyListRequest returns a decoder for requests sent to the
// hy_company companyList endpoint.
func DecodeCompanyListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewCompanyListPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetCompanyGroupResponse returns an encoder for responses returned by
// the hy_company getCompanyGroup endpoint.
func EncodeGetCompanyGroupResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(hycompanyviews.CompanyCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCompanyResponseCollection(res.Projected)
		case "detailid":
			body = NewCompanyResponseDetailidCollection(res.Projected)
		case "id":
			body = NewCompanyResponseIDCollection(res.Projected)
		case "idname":
			body = NewCompanyResponseIdnameCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetCompanyGroupRequest returns a decoder for requests sent to the
// hy_company getCompanyGroup endpoint.
func DecodeGetCompanyGroupRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body GetCompanyGroupRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateGetCompanyGroupRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			companyID int
			token     *string

			params = mux.Vars(r)
		)
		{
			companyIDRaw := params["company_id"]
			v, err2 := strconv.ParseInt(companyIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("companyID", companyIDRaw, "integer"))
			}
			companyID = int(v)
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetCompanyGroupPayload(&body, companyID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeCreateCompanyResponse returns an encoder for responses returned by the
// hy_company createCompany endpoint.
func EncodeCreateCompanyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCreateCompanyRequest returns a decoder for requests sent to the
// hy_company createCompany endpoint.
func DecodeCreateCompanyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateCompanyRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateCompanyRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewCreateCompanyPayload(&body, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateCompanyResponse returns an encoder for responses returned by the
// hy_company updateCompany endpoint.
func EncodeUpdateCompanyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateCompanyRequest returns a decoder for requests sent to the
// hy_company updateCompany endpoint.
func DecodeUpdateCompanyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateCompanyRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateCompanyRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			companyID int
			token     *string

			params = mux.Vars(r)
		)
		{
			companyIDRaw := params["company_id"]
			v, err2 := strconv.ParseInt(companyIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("companyID", companyIDRaw, "integer"))
			}
			companyID = int(v)
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateCompanyPayload(&body, companyID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteCompanyResponse returns an encoder for responses returned by the
// hy_company deleteCompany endpoint.
func EncodeDeleteCompanyResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteCompanyRequest returns a decoder for requests sent to the
// hy_company deleteCompany endpoint.
func DecodeDeleteCompanyRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			companyID int
			token     *string
			err       error

			params = mux.Vars(r)
		)
		{
			companyIDRaw := params["company_id"]
			v, err2 := strconv.ParseInt(companyIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("companyID", companyIDRaw, "integer"))
			}
			companyID = int(v)
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteCompanyPayload(companyID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// marshalHycompanyviewsCompanyViewToCompanyResponse builds a value of type
// *CompanyResponse from a value of type *hycompanyviews.CompanyView.
func marshalHycompanyviewsCompanyViewToCompanyResponse(v *hycompanyviews.CompanyView) *CompanyResponse {
	res := &CompanyResponse{
		ID:          v.ID,
		CompanyID:   v.CompanyID,
		Name:        *v.Name,
		IsHq:        v.IsHq,
		CountryName: v.CountryName,
		Address:     *v.Address,
	}

	return res
}

// marshalHycompanyviewsCompanyViewToCompanyResponseDetailid builds a value of
// type *CompanyResponseDetailid from a value of type
// *hycompanyviews.CompanyView.
func marshalHycompanyviewsCompanyViewToCompanyResponseDetailid(v *hycompanyviews.CompanyView) *CompanyResponseDetailid {
	res := &CompanyResponseDetailid{
		ID: v.ID,
	}

	return res
}

// marshalHycompanyviewsCompanyViewToCompanyResponseID builds a value of type
// *CompanyResponseID from a value of type *hycompanyviews.CompanyView.
func marshalHycompanyviewsCompanyViewToCompanyResponseID(v *hycompanyviews.CompanyView) *CompanyResponseID {
	res := &CompanyResponseID{
		CompanyID: v.CompanyID,
	}

	return res
}

// marshalHycompanyviewsCompanyViewToCompanyResponseIdname builds a value of
// type *CompanyResponseIdname from a value of type *hycompanyviews.CompanyView.
func marshalHycompanyviewsCompanyViewToCompanyResponseIdname(v *hycompanyviews.CompanyView) *CompanyResponseIdname {
	res := &CompanyResponseIdname{
		CompanyID: v.CompanyID,
		Name:      *v.Name,
	}

	return res
}
