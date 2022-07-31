// Code generated by goa v3.7.13, DO NOT EDIT.
//
// hy_userWorkHistory HTTP server types
//
// Command:
// $ goa gen resume/design

package server

import (
	hyuserworkhistory "resume/gen/hy_user_work_history"
	hyuserworkhistoryviews "resume/gen/hy_user_work_history/views"
)

// UserworkhistoryResponseCollection is the type of the "hy_userWorkHistory"
// service "getUserWorkHistory" endpoint HTTP response body.
type UserworkhistoryResponseCollection []*UserworkhistoryResponse

// UserworkhistoryResponse is used to define fields on response body types.
type UserworkhistoryResponse struct {
	// Job Title
	Title string `form:"title" json:"title" xml:"title"`
	// Company name
	Company string `form:"company" json:"company" xml:"company"`
	// Country code
	Country string `form:"country" json:"country" xml:"country"`
	// worked period
	Term *string `form:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// job description
	Description interface{} `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// used techs
	Techs interface{} `form:"techs,omitempty" json:"techs,omitempty" xml:"techs,omitempty"`
}

// NewUserworkhistoryResponseCollection builds the HTTP response body from the
// result of the "getUserWorkHistory" endpoint of the "hy_userWorkHistory"
// service.
func NewUserworkhistoryResponseCollection(res hyuserworkhistoryviews.UserworkhistoryCollectionView) UserworkhistoryResponseCollection {
	body := make([]*UserworkhistoryResponse, len(res))
	for i, val := range res {
		body[i] = marshalHyuserworkhistoryviewsUserworkhistoryViewToUserworkhistoryResponse(val)
	}
	return body
}

// NewGetUserWorkHistoryPayload builds a hy_userWorkHistory service
// getUserWorkHistory endpoint payload.
func NewGetUserWorkHistoryPayload(userID int, token *string) *hyuserworkhistory.GetUserWorkHistoryPayload {
	v := &hyuserworkhistory.GetUserWorkHistoryPayload{}
	v.UserID = &userID
	v.Token = token

	return v
}
