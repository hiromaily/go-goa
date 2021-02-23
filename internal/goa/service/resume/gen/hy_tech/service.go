// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_tech service
//
// Command:
// $ goa gen resume/design

package hytech

import (
	"context"
	hytechviews "resume/gen/hy_tech/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The company service returns company data
type Service interface {
	// List all techs
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "id": id is the view used for C U D
	TechList(context.Context, *TechListPayload) (res TechCollection, view string, err error)
	// get tech with given tech id
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "detailid": only company's detail id
	//	- "id": only company's id
	//	- "idname": only company's id and name
	GetTech(context.Context, *GetTechPayload) (res *Company, view string, err error)
	// Create new tech
	CreateTech(context.Context, *CreateTechPayload) (err error)
	// Change tech properties
	UpdateTech(context.Context, *UpdateTechPayload) (err error)
	// Delete tech
	DeleteTech(context.Context, *DeleteTechPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hy_tech"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"techList", "getTech", "createTech", "updateTech", "deleteTech"}

// TechListPayload is the payload type of the hy_tech service techList method.
type TechListPayload struct {
	// JWT token used to perform authorization
	Token *string
}

// TechCollection is the result type of the hy_tech service techList method.
type TechCollection []*Tech

// GetTechPayload is the payload type of the hy_tech service getTech method.
type GetTechPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Tech ID
	TechID *int
}

// Company is the result type of the hy_tech service getTech method.
type Company struct {
	// ID
	ID *int
	// ID
	CompanyID *int
	// Company name
	Name        string
	HqFlg       *string
	CountryName *string
	// Company Address
	Address string
	// Datetime
	CreatedAt *string
	// Datetime
	UpdatedAt *string
}

// CreateTechPayload is the payload type of the hy_tech service createTech
// method.
type CreateTechPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Tech name
	Name string
}

// UpdateTechPayload is the payload type of the hy_tech service updateTech
// method.
type UpdateTechPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Tech ID
	TechID *int
	// Tech name
	Name string
}

// DeleteTechPayload is the payload type of the hy_tech service deleteTech
// method.
type DeleteTechPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Tech ID
	TechID *int
}

// A tech information
type Tech struct {
	// ID
	ID *int
	// Tech name
	Name string
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

// NewTechCollection initializes result type TechCollection from viewed result
// type TechCollection.
func NewTechCollection(vres hytechviews.TechCollection) TechCollection {
	var res TechCollection
	switch vres.View {
	case "default", "":
		res = newTechCollection(vres.Projected)
	case "id":
		res = newTechCollectionID(vres.Projected)
	}
	return res
}

// NewViewedTechCollection initializes viewed result type TechCollection from
// result type TechCollection using the given view.
func NewViewedTechCollection(res TechCollection, view string) hytechviews.TechCollection {
	var vres hytechviews.TechCollection
	switch view {
	case "default", "":
		p := newTechCollectionView(res)
		vres = hytechviews.TechCollection{Projected: p, View: "default"}
	case "id":
		p := newTechCollectionViewID(res)
		vres = hytechviews.TechCollection{Projected: p, View: "id"}
	}
	return vres
}

// NewCompany initializes result type Company from viewed result type Company.
func NewCompany(vres *hytechviews.Company) *Company {
	var res *Company
	switch vres.View {
	case "default", "":
		res = newCompany(vres.Projected)
	case "detailid":
		res = newCompanyDetailid(vres.Projected)
	case "id":
		res = newCompanyID(vres.Projected)
	case "idname":
		res = newCompanyIdname(vres.Projected)
	}
	return res
}

// NewViewedCompany initializes viewed result type Company from result type
// Company using the given view.
func NewViewedCompany(res *Company, view string) *hytechviews.Company {
	var vres *hytechviews.Company
	switch view {
	case "default", "":
		p := newCompanyView(res)
		vres = &hytechviews.Company{Projected: p, View: "default"}
	case "detailid":
		p := newCompanyViewDetailid(res)
		vres = &hytechviews.Company{Projected: p, View: "detailid"}
	case "id":
		p := newCompanyViewID(res)
		vres = &hytechviews.Company{Projected: p, View: "id"}
	case "idname":
		p := newCompanyViewIdname(res)
		vres = &hytechviews.Company{Projected: p, View: "idname"}
	}
	return vres
}

// newTechCollection converts projected type TechCollection to service type
// TechCollection.
func newTechCollection(vres hytechviews.TechCollectionView) TechCollection {
	res := make(TechCollection, len(vres))
	for i, n := range vres {
		res[i] = newTech(n)
	}
	return res
}

// newTechCollectionID converts projected type TechCollection to service type
// TechCollection.
func newTechCollectionID(vres hytechviews.TechCollectionView) TechCollection {
	res := make(TechCollection, len(vres))
	for i, n := range vres {
		res[i] = newTechID(n)
	}
	return res
}

// newTechCollectionView projects result type TechCollection to projected type
// TechCollectionView using the "default" view.
func newTechCollectionView(res TechCollection) hytechviews.TechCollectionView {
	vres := make(hytechviews.TechCollectionView, len(res))
	for i, n := range res {
		vres[i] = newTechView(n)
	}
	return vres
}

// newTechCollectionViewID projects result type TechCollection to projected
// type TechCollectionView using the "id" view.
func newTechCollectionViewID(res TechCollection) hytechviews.TechCollectionView {
	vres := make(hytechviews.TechCollectionView, len(res))
	for i, n := range res {
		vres[i] = newTechViewID(n)
	}
	return vres
}

// newTech converts projected type Tech to service type Tech.
func newTech(vres *hytechviews.TechView) *Tech {
	res := &Tech{
		ID: vres.ID,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newTechID converts projected type Tech to service type Tech.
func newTechID(vres *hytechviews.TechView) *Tech {
	res := &Tech{
		ID: vres.ID,
	}
	return res
}

// newTechView projects result type Tech to projected type TechView using the
// "default" view.
func newTechView(res *Tech) *hytechviews.TechView {
	vres := &hytechviews.TechView{
		ID:   res.ID,
		Name: &res.Name,
	}
	return vres
}

// newTechViewID projects result type Tech to projected type TechView using the
// "id" view.
func newTechViewID(res *Tech) *hytechviews.TechView {
	vres := &hytechviews.TechView{
		ID: res.ID,
	}
	return vres
}

// newCompany converts projected type Company to service type Company.
func newCompany(vres *hytechviews.CompanyView) *Company {
	res := &Company{
		ID:          vres.ID,
		CompanyID:   vres.CompanyID,
		HqFlg:       vres.HqFlg,
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
func newCompanyDetailid(vres *hytechviews.CompanyView) *Company {
	res := &Company{
		ID: vres.ID,
	}
	return res
}

// newCompanyID converts projected type Company to service type Company.
func newCompanyID(vres *hytechviews.CompanyView) *Company {
	res := &Company{
		CompanyID: vres.CompanyID,
	}
	return res
}

// newCompanyIdname converts projected type Company to service type Company.
func newCompanyIdname(vres *hytechviews.CompanyView) *Company {
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
func newCompanyView(res *Company) *hytechviews.CompanyView {
	vres := &hytechviews.CompanyView{
		ID:          res.ID,
		CompanyID:   res.CompanyID,
		Name:        &res.Name,
		HqFlg:       res.HqFlg,
		CountryName: res.CountryName,
		Address:     &res.Address,
	}
	return vres
}

// newCompanyViewDetailid projects result type Company to projected type
// CompanyView using the "detailid" view.
func newCompanyViewDetailid(res *Company) *hytechviews.CompanyView {
	vres := &hytechviews.CompanyView{
		ID: res.ID,
	}
	return vres
}

// newCompanyViewID projects result type Company to projected type CompanyView
// using the "id" view.
func newCompanyViewID(res *Company) *hytechviews.CompanyView {
	vres := &hytechviews.CompanyView{
		CompanyID: res.CompanyID,
	}
	return vres
}

// newCompanyViewIdname projects result type Company to projected type
// CompanyView using the "idname" view.
func newCompanyViewIdname(res *Company) *hytechviews.CompanyView {
	vres := &hytechviews.CompanyView{
		CompanyID: res.CompanyID,
		Name:      &res.Name,
	}
	return vres
}
