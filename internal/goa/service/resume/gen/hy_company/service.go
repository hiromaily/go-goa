// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_company service
//
// Command:
// $ goa gen resume/design

package hycompany

import (
	"context"
	hycompanyviews "resume/gen/hy_company/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The company service returns company data
type Service interface {
	// List all companies
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "detailid": only company's detail id
	//	- "id": only company's id
	//	- "idname": only company's id and name
	CompanyList(context.Context, *CompanyListPayload) (res CompanyCollection, view string, err error)
	// Retrieve company with given company_id
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "detailid": only company's detail id
	//	- "id": only company's id
	//	- "idname": only company's id and name
	GetCompanyGroup(context.Context, *GetCompanyGroupPayload) (res CompanyCollection, view string, err error)
	// Create new company
	CreateCompany(context.Context, *CreateCompanyPayload) (err error)
	// Change company properties
	UpdateCompany(context.Context, *UpdateCompanyPayload) (err error)
	// Delete company
	DeleteCompany(context.Context, *DeleteCompanyPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hy_company"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"companyList", "getCompanyGroup", "createCompany", "updateCompany", "deleteCompany"}

// CompanyListPayload is the payload type of the hy_company service companyList
// method.
type CompanyListPayload struct {
	// JWT token used to perform authorization
	Token *string
}

// CompanyCollection is the result type of the hy_company service companyList
// method.
type CompanyCollection []*Company

// GetCompanyGroupPayload is the payload type of the hy_company service
// getCompanyGroup method.
type GetCompanyGroupPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company ID
	CompanyID *int
	// Head Quarters flag
	IsHq *string
}

// CreateCompanyPayload is the payload type of the hy_company service
// createCompany method.
type CreateCompanyPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company name
	Name string
	// Country ID
	CountryID int
	// Company Address
	Address string
}

// UpdateCompanyPayload is the payload type of the hy_company service
// updateCompany method.
type UpdateCompanyPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company ID
	CompanyID *int
	// Company name
	Name string
	// Country ID
	CountryID int
	// Company Address
	Address string
}

// DeleteCompanyPayload is the payload type of the hy_company service
// deleteCompany method.
type DeleteCompanyPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company ID
	CompanyID *int
}

// A company information
type Company struct {
	// ID
	ID *int
	// ID
	CompanyID *int
	// Company name
	Name        string
	IsHq        *string
	CountryName *string
	// Company Address
	Address string
	// Datetime
	CreatedAt *string
	// Datetime
	UpdatedAt *string
}

// MakeNoContent builds a goa.ServiceError from an error.
func MakeNoContent(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NoContent",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NotFound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewCompanyCollection initializes result type CompanyCollection from viewed
// result type CompanyCollection.
func NewCompanyCollection(vres hycompanyviews.CompanyCollection) CompanyCollection {
	var res CompanyCollection
	switch vres.View {
	case "default", "":
		res = newCompanyCollection(vres.Projected)
	case "detailid":
		res = newCompanyCollectionDetailid(vres.Projected)
	case "id":
		res = newCompanyCollectionID(vres.Projected)
	case "idname":
		res = newCompanyCollectionIdname(vres.Projected)
	}
	return res
}

// NewViewedCompanyCollection initializes viewed result type CompanyCollection
// from result type CompanyCollection using the given view.
func NewViewedCompanyCollection(res CompanyCollection, view string) hycompanyviews.CompanyCollection {
	var vres hycompanyviews.CompanyCollection
	switch view {
	case "default", "":
		p := newCompanyCollectionView(res)
		vres = hycompanyviews.CompanyCollection{Projected: p, View: "default"}
	case "detailid":
		p := newCompanyCollectionViewDetailid(res)
		vres = hycompanyviews.CompanyCollection{Projected: p, View: "detailid"}
	case "id":
		p := newCompanyCollectionViewID(res)
		vres = hycompanyviews.CompanyCollection{Projected: p, View: "id"}
	case "idname":
		p := newCompanyCollectionViewIdname(res)
		vres = hycompanyviews.CompanyCollection{Projected: p, View: "idname"}
	}
	return vres
}

// newCompanyCollection converts projected type CompanyCollection to service
// type CompanyCollection.
func newCompanyCollection(vres hycompanyviews.CompanyCollectionView) CompanyCollection {
	res := make(CompanyCollection, len(vres))
	for i, n := range vres {
		res[i] = newCompany(n)
	}
	return res
}

// newCompanyCollectionDetailid converts projected type CompanyCollection to
// service type CompanyCollection.
func newCompanyCollectionDetailid(vres hycompanyviews.CompanyCollectionView) CompanyCollection {
	res := make(CompanyCollection, len(vres))
	for i, n := range vres {
		res[i] = newCompanyDetailid(n)
	}
	return res
}

// newCompanyCollectionID converts projected type CompanyCollection to service
// type CompanyCollection.
func newCompanyCollectionID(vres hycompanyviews.CompanyCollectionView) CompanyCollection {
	res := make(CompanyCollection, len(vres))
	for i, n := range vres {
		res[i] = newCompanyID(n)
	}
	return res
}

// newCompanyCollectionIdname converts projected type CompanyCollection to
// service type CompanyCollection.
func newCompanyCollectionIdname(vres hycompanyviews.CompanyCollectionView) CompanyCollection {
	res := make(CompanyCollection, len(vres))
	for i, n := range vres {
		res[i] = newCompanyIdname(n)
	}
	return res
}

// newCompanyCollectionView projects result type CompanyCollection to projected
// type CompanyCollectionView using the "default" view.
func newCompanyCollectionView(res CompanyCollection) hycompanyviews.CompanyCollectionView {
	vres := make(hycompanyviews.CompanyCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCompanyView(n)
	}
	return vres
}

// newCompanyCollectionViewDetailid projects result type CompanyCollection to
// projected type CompanyCollectionView using the "detailid" view.
func newCompanyCollectionViewDetailid(res CompanyCollection) hycompanyviews.CompanyCollectionView {
	vres := make(hycompanyviews.CompanyCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCompanyViewDetailid(n)
	}
	return vres
}

// newCompanyCollectionViewID projects result type CompanyCollection to
// projected type CompanyCollectionView using the "id" view.
func newCompanyCollectionViewID(res CompanyCollection) hycompanyviews.CompanyCollectionView {
	vres := make(hycompanyviews.CompanyCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCompanyViewID(n)
	}
	return vres
}

// newCompanyCollectionViewIdname projects result type CompanyCollection to
// projected type CompanyCollectionView using the "idname" view.
func newCompanyCollectionViewIdname(res CompanyCollection) hycompanyviews.CompanyCollectionView {
	vres := make(hycompanyviews.CompanyCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCompanyViewIdname(n)
	}
	return vres
}

// newCompany converts projected type Company to service type Company.
func newCompany(vres *hycompanyviews.CompanyView) *Company {
	res := &Company{
		ID:          vres.ID,
		CompanyID:   vres.CompanyID,
		IsHq:        vres.IsHq,
		CountryName: vres.CountryName,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Address != nil {
		res.Address = *vres.Address
	}
	return res
}

// newCompanyDetailid converts projected type Company to service type Company.
func newCompanyDetailid(vres *hycompanyviews.CompanyView) *Company {
	res := &Company{
		ID: vres.ID,
	}
	return res
}

// newCompanyID converts projected type Company to service type Company.
func newCompanyID(vres *hycompanyviews.CompanyView) *Company {
	res := &Company{
		CompanyID: vres.CompanyID,
	}
	return res
}

// newCompanyIdname converts projected type Company to service type Company.
func newCompanyIdname(vres *hycompanyviews.CompanyView) *Company {
	res := &Company{
		CompanyID: vres.CompanyID,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newCompanyView projects result type Company to projected type CompanyView
// using the "default" view.
func newCompanyView(res *Company) *hycompanyviews.CompanyView {
	vres := &hycompanyviews.CompanyView{
		ID:          res.ID,
		CompanyID:   res.CompanyID,
		Name:        &res.Name,
		IsHq:        res.IsHq,
		CountryName: res.CountryName,
		Address:     &res.Address,
	}
	return vres
}

// newCompanyViewDetailid projects result type Company to projected type
// CompanyView using the "detailid" view.
func newCompanyViewDetailid(res *Company) *hycompanyviews.CompanyView {
	vres := &hycompanyviews.CompanyView{
		ID: res.ID,
	}
	return vres
}

// newCompanyViewID projects result type Company to projected type CompanyView
// using the "id" view.
func newCompanyViewID(res *Company) *hycompanyviews.CompanyView {
	vres := &hycompanyviews.CompanyView{
		CompanyID: res.CompanyID,
	}
	return vres
}

// newCompanyViewIdname projects result type Company to projected type
// CompanyView using the "idname" view.
func newCompanyViewIdname(res *Company) *hycompanyviews.CompanyView {
	vres := &hycompanyviews.CompanyView{
		CompanyID: res.CompanyID,
		Name:      &res.Name,
	}
	return vres
}
