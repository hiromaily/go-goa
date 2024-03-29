// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_user views
//
// Command:
// $ goa gen resume/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UserCollection is the viewed result type that is projected based on a view.
type UserCollection struct {
	// Type to project
	Projected UserCollectionView
	// View to render
	View string
}

// User is the viewed result type that is projected based on a view.
type User struct {
	// Type to project
	Projected *UserView
	// View to render
	View string
}

// UserCollectionView is a type that runs validations on a projected type.
type UserCollectionView []*UserView

// UserView is a type that runs validations on a projected type.
type UserView struct {
	// Key ID
	UserID *int
	// User name
	UserName *string
	// E-mail of user
	Email *string
	// Password
	Password *string
	// Datetime
	CreatedAt *string
	// Datetime
	UpdatedAt *string
}

var (
	// UserCollectionMap is a map indexing the attribute names of UserCollection by
	// view name.
	UserCollectionMap = map[string][]string{
		"default": {
			"user_id",
			"user_name",
			"email",
		},
		"id": {
			"user_id",
		},
	}
	// UserMap is a map indexing the attribute names of User by view name.
	UserMap = map[string][]string{
		"default": {
			"user_id",
			"user_name",
			"email",
		},
		"id": {
			"user_id",
		},
	}
)

// ValidateUserCollection runs the validations defined on the viewed result
// type UserCollection.
func ValidateUserCollection(result UserCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserCollectionView(result.Projected)
	case "id":
		err = ValidateUserCollectionViewID(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "id"})
	}
	return
}

// ValidateUser runs the validations defined on the viewed result type User.
func ValidateUser(result *User) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserView(result.Projected)
	case "id":
		err = ValidateUserViewID(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "id"})
	}
	return
}

// ValidateUserCollectionView runs the validations defined on
// UserCollectionView using the "default" view.
func ValidateUserCollectionView(result UserCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateUserView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUserCollectionViewID runs the validations defined on
// UserCollectionView using the "id" view.
func ValidateUserCollectionViewID(result UserCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateUserViewID(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUserView runs the validations defined on UserView using the
// "default" view.
func ValidateUserView(result *UserView) (err error) {
	if result.UserID != nil {
		if *result.UserID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.user_id", *result.UserID, 1, true))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 2, true))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 20, false))
		}
	}
	if result.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email", *result.Email, goa.FormatEmail))
	}
	return
}

// ValidateUserViewID runs the validations defined on UserView using the "id"
// view.
func ValidateUserViewID(result *UserView) (err error) {
	if result.UserID != nil {
		if *result.UserID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.user_id", *result.UserID, 1, true))
		}
	}
	return
}
