// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.1.0-dirty
//
// API "api": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// An authorized response (default view)
//
// Identifier: application/vnd.authorized+json; view=default
type Authorized struct {
	// JWT token
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the Authorized media type instance.
func (mt *Authorized) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// DecodeAuthorized decodes the Authorized instance encoded in resp body.
func (c *Client) DecodeAuthorized(resp *http.Response) (*Authorized, error) {
	var decoded Authorized
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A company information (default view)
//
// Identifier: application/vnd.company+json; view=default
type Company struct {
	Address string `form:"address" json:"address" xml:"address"`
	Country string `form:"country" json:"country" xml:"country"`
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Company media type instance.
func (mt *Company) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Country == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "country"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	return
}

// A company information (link view)
//
// Identifier: application/vnd.company+json; view=link
type CompanyLink struct {
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the CompanyLink media type instance.
func (mt *CompanyLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// A company information (tiny view)
//
// Identifier: application/vnd.company+json; view=tiny
type CompanyTiny struct {
	// API href of company
	Href string `form:"href" json:"href" xml:"href"`
	// ID of company
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the CompanyTiny media type instance.
func (mt *CompanyTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeCompany decodes the Company instance encoded in resp body.
func (c *Client) DecodeCompany(resp *http.Response) (*Company, error) {
	var decoded Company
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCompanyLink decodes the CompanyLink instance encoded in resp body.
func (c *Client) DecodeCompanyLink(resp *http.Response) (*CompanyLink, error) {
	var decoded CompanyLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeCompanyTiny decodes the CompanyTiny instance encoded in resp body.
func (c *Client) DecodeCompanyTiny(resp *http.Response) (*CompanyTiny, error) {
	var decoded CompanyTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CompanyCollection is the media type for an array of Company (default view)
//
// Identifier: application/vnd.company+json; type=collection; view=default
type CompanyCollection []*Company

// Validate validates the CompanyCollection media type instance.
func (mt CompanyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// CompanyCollection is the media type for an array of Company (tiny view)
//
// Identifier: application/vnd.company+json; type=collection; view=tiny
type CompanyTinyCollection []*CompanyTiny

// Validate validates the CompanyTinyCollection media type instance.
func (mt CompanyTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCompanyCollection decodes the CompanyCollection instance encoded in resp body.
func (c *Client) DecodeCompanyCollection(resp *http.Response) (CompanyCollection, error) {
	var decoded CompanyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeCompanyTinyCollection decodes the CompanyTinyCollection instance encoded in resp body.
func (c *Client) DecodeCompanyTinyCollection(resp *http.Response) (CompanyTinyCollection, error) {
	var decoded CompanyTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A user information (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Email string `form:"email" json:"email" xml:"email"`
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// A user information (link view)
//
// Identifier: application/vnd.user+json; view=link
type UserLink struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// A user information (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserLink decodes the UserLink instance encoded in resp body.
func (c *Client) DecodeUserLink(resp *http.Response) (*UserLink, error) {
	var decoded UserLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserTiny decodes the UserTiny instance encoded in resp body.
func (c *Client) DecodeUserTiny(resp *http.Response) (*UserTiny, error) {
	var decoded UserTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (tiny view)
//
// Identifier: application/vnd.user+json; type=collection; view=tiny
type UserTinyCollection []*UserTiny

// Validate validates the UserTinyCollection media type instance.
func (mt UserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeUserTinyCollection decodes the UserTinyCollection instance encoded in resp body.
func (c *Client) DecodeUserTinyCollection(resp *http.Response) (UserTinyCollection, error) {
	var decoded UserTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A user who belongs to which companies (default view)
//
// Identifier: application/vnd.usercomany+json; view=default
type Usercomany struct {
	// ID of user id
	CompanyID int `form:"company_id" json:"company_id" xml:"company_id"`
	// API href of bottle
	Href string `form:"href" json:"href" xml:"href"`
	// ID of user company
	ID int `form:"id" json:"id" xml:"id"`
	// ID of user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Usercomany media type instance.
func (mt *Usercomany) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// DecodeUsercomany decodes the Usercomany instance encoded in resp body.
func (c *Client) DecodeUsercomany(resp *http.Response) (*Usercomany, error) {
	var decoded Usercomany
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}