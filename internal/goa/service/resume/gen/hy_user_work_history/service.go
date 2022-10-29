// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_userWorkHistory service
//
// Command:
// $ goa gen resume/design

package hyuserworkhistory

import (
	"context"
	hyuserworkhistoryviews "resume/gen/hy_user_work_history/views"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The user work history service returns user working history data
type Service interface {
	// get user's work history
	GetUserWorkHistory(context.Context, *GetUserWorkHistoryPayload) (res UserworkhistoryCollection, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hy_userWorkHistory"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"getUserWorkHistory"}

// GetUserWorkHistoryPayload is the payload type of the hy_userWorkHistory
// service getUserWorkHistory method.
type GetUserWorkHistoryPayload struct {
	// JWT token used to perform authorization
	Token *string
	// User ID
	UserID int
}

// A user information
type Userworkhistory struct {
	// Job Title
	Title *string
	// Company name
	CompanyName *string
	// Country code
	CountryName *string
	// worked period
	Term *string
	// job description
	Description interface{}
	// used techs
	Techs interface{}
}

// UserworkhistoryCollection is the result type of the hy_userWorkHistory
// service getUserWorkHistory method.
type UserworkhistoryCollection []*Userworkhistory

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "NotFound", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "BadRequest", false, false, false)
}

// NewUserworkhistoryCollection initializes result type
// UserworkhistoryCollection from viewed result type UserworkhistoryCollection.
func NewUserworkhistoryCollection(vres hyuserworkhistoryviews.UserworkhistoryCollection) UserworkhistoryCollection {
	return newUserworkhistoryCollection(vres.Projected)
}

// NewViewedUserworkhistoryCollection initializes viewed result type
// UserworkhistoryCollection from result type UserworkhistoryCollection using
// the given view.
func NewViewedUserworkhistoryCollection(res UserworkhistoryCollection, view string) hyuserworkhistoryviews.UserworkhistoryCollection {
	p := newUserworkhistoryCollectionView(res)
	return hyuserworkhistoryviews.UserworkhistoryCollection{Projected: p, View: "default"}
}

// newUserworkhistoryCollection converts projected type
// UserworkhistoryCollection to service type UserworkhistoryCollection.
func newUserworkhistoryCollection(vres hyuserworkhistoryviews.UserworkhistoryCollectionView) UserworkhistoryCollection {
	res := make(UserworkhistoryCollection, len(vres))
	for i, n := range vres {
		res[i] = newUserworkhistory(n)
	}
	return res
}

// newUserworkhistoryCollectionView projects result type
// UserworkhistoryCollection to projected type UserworkhistoryCollectionView
// using the "default" view.
func newUserworkhistoryCollectionView(res UserworkhistoryCollection) hyuserworkhistoryviews.UserworkhistoryCollectionView {
	vres := make(hyuserworkhistoryviews.UserworkhistoryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newUserworkhistoryView(n)
	}
	return vres
}

// newUserworkhistory converts projected type Userworkhistory to service type
// Userworkhistory.
func newUserworkhistory(vres *hyuserworkhistoryviews.UserworkhistoryView) *Userworkhistory {
	res := &Userworkhistory{
		Title:       vres.Title,
		CompanyName: vres.CompanyName,
		CountryName: vres.CountryName,
		Term:        vres.Term,
		Description: vres.Description,
		Techs:       vres.Techs,
	}
	return res
}

// newUserworkhistoryView projects result type Userworkhistory to projected
// type UserworkhistoryView using the "default" view.
func newUserworkhistoryView(res *Userworkhistory) *hyuserworkhistoryviews.UserworkhistoryView {
	vres := &hyuserworkhistoryviews.UserworkhistoryView{
		Title:       res.Title,
		CompanyName: res.CompanyName,
		CountryName: res.CountryName,
		Term:        res.Term,
		Description: res.Description,
		Techs:       res.Techs,
	}
	return vres
}
