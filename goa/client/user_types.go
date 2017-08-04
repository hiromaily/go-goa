// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application User Types
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// commonResponse user type.
type commonResponse struct {
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the commonResponse type instance.
func (ut *commonResponse) Validate() (err error) {
	if ut.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.CreatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_at`, *ut.CreatedAt, goa.FormatDateTime, err2))
		}
	}
	if ut.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.UpdatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.updated_at`, *ut.UpdatedAt, goa.FormatDateTime, err2))
		}
	}
	return
}

// Publicize creates CommonResponse from commonResponse
func (ut *commonResponse) Publicize() *CommonResponse {
	var pub CommonResponse
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.UpdatedAt != nil {
		pub.UpdatedAt = ut.UpdatedAt
	}
	return &pub
}

// CommonResponse user type.
type CommonResponse struct {
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// Validate validates the CommonResponse type instance.
func (ut *CommonResponse) Validate() (err error) {
	if ut.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.CreatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_at`, *ut.CreatedAt, goa.FormatDateTime, err2))
		}
	}
	if ut.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *ut.UpdatedAt); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.updated_at`, *ut.UpdatedAt, goa.FormatDateTime, err2))
		}
	}
	return
}

// companyPayload user type.
type companyPayload struct {
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	// Company name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the companyPayload type instance.
func (ut *companyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 2, true))
		}
	}
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 80, false))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 1, true))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 999, false))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 40, false))
		}
	}
	return
}

// Publicize creates CompanyPayload from companyPayload
func (ut *companyPayload) Publicize() *CompanyPayload {
	var pub CompanyPayload
	if ut.Address != nil {
		pub.Address = ut.Address
	}
	if ut.CountryID != nil {
		pub.CountryID = ut.CountryID
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// CompanyPayload user type.
type CompanyPayload struct {
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	// Company name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the CompanyPayload type instance.
func (ut *CompanyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 2, true))
		}
	}
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 80, false))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 1, true))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 999, false))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 40, false))
		}
	}
	return
}

// companyTinyPayload user type.
type companyTinyPayload struct {
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
}

// Validate validates the companyTinyPayload type instance.
func (ut *companyTinyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 2, true))
		}
	}
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 80, false))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 1, true))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 999, false))
		}
	}
	return
}

// Publicize creates CompanyTinyPayload from companyTinyPayload
func (ut *companyTinyPayload) Publicize() *CompanyTinyPayload {
	var pub CompanyTinyPayload
	if ut.Address != nil {
		pub.Address = ut.Address
	}
	if ut.CountryID != nil {
		pub.CountryID = ut.CountryID
	}
	return &pub
}

// CompanyTinyPayload user type.
type CompanyTinyPayload struct {
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
}

// Validate validates the CompanyTinyPayload type instance.
func (ut *CompanyTinyPayload) Validate() (err error) {
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 2, true))
		}
	}
	if ut.Address != nil {
		if utf8.RuneCountInString(*ut.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.address`, *ut.Address, utf8.RuneCountInString(*ut.Address), 80, false))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 1, true))
		}
	}
	if ut.CountryID != nil {
		if *ut.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.country_id`, *ut.CountryID, 999, false))
		}
	}
	return
}

// loginPayload user type.
type loginPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, false))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, false))
		}
	}
	return
}

// userPayload user type.
type userPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, false))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 2, true))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 20, false))
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
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.UserName != nil {
		pub.UserName = ut.UserName
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// User name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 20, false))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 2, true))
		}
	}
	if ut.UserName != nil {
		if utf8.RuneCountInString(*ut.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.user_name`, *ut.UserName, utf8.RuneCountInString(*ut.UserName), 20, false))
		}
	}
	return
}
