// Code generated by goa v3.10.1, DO NOT EDIT.
//
// HTTP request path constructors for the hy_user service.
//
// Command:
// $ goa gen resume/design

package server

import (
	"fmt"
)

// UserListHyUserPath returns the URL path to the hy_user service userList HTTP endpoint.
func UserListHyUserPath() string {
	return "/user"
}

// GetUserHyUserPath returns the URL path to the hy_user service getUser HTTP endpoint.
func GetUserHyUserPath(userID int) string {
	return fmt.Sprintf("/user/%v", userID)
}

// CreateUserHyUserPath returns the URL path to the hy_user service createUser HTTP endpoint.
func CreateUserHyUserPath() string {
	return "/user"
}

// UpdateUserHyUserPath returns the URL path to the hy_user service updateUser HTTP endpoint.
func UpdateUserHyUserPath(userID int) string {
	return fmt.Sprintf("/user/%v", userID)
}

// DeleteUserHyUserPath returns the URL path to the hy_user service deleteUser HTTP endpoint.
func DeleteUserHyUserPath(userID int) string {
	return fmt.Sprintf("/user/%v", userID)
}
