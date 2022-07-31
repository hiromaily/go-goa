// Code generated by goa v3.7.13, DO NOT EDIT.
//
// hy_usertech views
//
// Command:
// $ goa gen resume/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// UsertechCollection is the viewed result type that is projected based on a
// view.
type UsertechCollection struct {
	// Type to project
	Projected UsertechCollectionView
	// View to render
	View string
}

// UsertechCollectionView is a type that runs validations on a projected type.
type UsertechCollectionView []*UsertechView

// UsertechView is a type that runs validations on a projected type.
type UsertechView struct {
	// ID
	ID *int
	// Tech name
	TechName *string
}

var (
	// UsertechCollectionMap is a map indexing the attribute names of
	// UsertechCollection by view name.
	UsertechCollectionMap = map[string][]string{
		"default": {
			"id",
			"tech_name",
		},
		"tech": {
			"tech_name",
		},
	}
	// UsertechMap is a map indexing the attribute names of Usertech by view name.
	UsertechMap = map[string][]string{
		"default": {
			"id",
			"tech_name",
		},
		"tech": {
			"tech_name",
		},
	}
)

// ValidateUsertechCollection runs the validations defined on the viewed result
// type UsertechCollection.
func ValidateUsertechCollection(result UsertechCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUsertechCollectionView(result.Projected)
	case "tech":
		err = ValidateUsertechCollectionViewTech(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tech"})
	}
	return
}

// ValidateUsertechCollectionView runs the validations defined on
// UsertechCollectionView using the "default" view.
func ValidateUsertechCollectionView(result UsertechCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateUsertechView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUsertechCollectionViewTech runs the validations defined on
// UsertechCollectionView using the "tech" view.
func ValidateUsertechCollectionViewTech(result UsertechCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateUsertechViewTech(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUsertechView runs the validations defined on UsertechView using the
// "default" view.
func ValidateUsertechView(result *UsertechView) (err error) {
	if result.TechName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tech_name", "result"))
	}
	if result.ID != nil {
		if *result.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.id", *result.ID, 1, true))
		}
	}
	if result.TechName != nil {
		if utf8.RuneCountInString(*result.TechName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.tech_name", *result.TechName, utf8.RuneCountInString(*result.TechName), 1, true))
		}
	}
	if result.TechName != nil {
		if utf8.RuneCountInString(*result.TechName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.tech_name", *result.TechName, utf8.RuneCountInString(*result.TechName), 40, false))
		}
	}
	return
}

// ValidateUsertechViewTech runs the validations defined on UsertechView using
// the "tech" view.
func ValidateUsertechViewTech(result *UsertechView) (err error) {
	if result.TechName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tech_name", "result"))
	}
	if result.TechName != nil {
		if utf8.RuneCountInString(*result.TechName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.tech_name", *result.TechName, utf8.RuneCountInString(*result.TechName), 1, true))
		}
	}
	if result.TechName != nil {
		if utf8.RuneCountInString(*result.TechName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.tech_name", *result.TechName, utf8.RuneCountInString(*result.TechName), 40, false))
		}
	}
	return
}
