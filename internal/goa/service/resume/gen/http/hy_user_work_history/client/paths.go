// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the hy_userWorkHistory service.
//
// Command:
// $ goa gen resume/design

package client

import (
	"fmt"
)

// GetUserWorkHistoryHyUserWorkHistoryPath returns the URL path to the hy_userWorkHistory service getUserWorkHistory HTTP endpoint.
func GetUserWorkHistoryHyUserWorkHistoryPath(userID int) string {
	return fmt.Sprintf("/user/%v/workhistory", userID)
}
