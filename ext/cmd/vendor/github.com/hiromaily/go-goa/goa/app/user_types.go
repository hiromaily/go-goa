// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.1.0-dirty
//
// API "api": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// bottlePayload user type.
type bottlePayload struct {
	// Date of creation
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// user id
	CreatedBy *int `form:"created_by,omitempty" json:"created_by,omitempty" xml:"created_by,omitempty"`
	// Date of update
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// user id
	UpdatedBy *int `form:"updated_by,omitempty" json:"updated_by,omitempty" xml:"updated_by,omitempty"`
}

// Publicize creates BottlePayload from bottlePayload
func (ut *bottlePayload) Publicize() *BottlePayload {
	var pub BottlePayload
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.CreatedBy != nil {
		pub.CreatedBy = ut.CreatedBy
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	if ut.UpdatedBy != nil {
		pub.UpdatedBy = ut.UpdatedBy
	}
	return &pub
}

// BottlePayload user type.
type BottlePayload struct {
	// Date of creation
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// user id
	CreatedBy *int `form:"created_by,omitempty" json:"created_by,omitempty" xml:"created_by,omitempty"`
	// Date of update
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// user id
	UpdatedBy *int `form:"updated_by,omitempty" json:"updated_by,omitempty" xml:"updated_by,omitempty"`
}

// companyPayload user type.
type companyPayload struct {
	// Address of company
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country of HQ
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Name of company
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Publicize creates CompanyPayload from companyPayload
func (ut *companyPayload) Publicize() *CompanyPayload {
	var pub CompanyPayload
	if ut.Address != nil {
		pub.Address = ut.Address
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// CompanyPayload user type.
type CompanyPayload struct {
	// Address of company
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country of HQ
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Name of company
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// userPayload user type.
type userPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}
