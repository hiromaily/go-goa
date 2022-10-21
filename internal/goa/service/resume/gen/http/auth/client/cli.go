// Code generated by goa v3.10.1, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen resume/design

package client

import (
	"encoding/json"
	"fmt"
	auth "resume/gen/auth"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildLoginPayload builds the payload for the auth login endpoint from CLI
// flags.
func BuildLoginPayload(authLoginBody string) (*auth.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(authLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"hy@gmail.com\",\n      \"password\": \"xxxxxxxx\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if utf8.RuneCountInString(body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 8, true))
		}
		if utf8.RuneCountInString(body.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 20, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &auth.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}
