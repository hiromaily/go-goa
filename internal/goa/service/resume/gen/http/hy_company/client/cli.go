// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_company HTTP client CLI support package
//
// Command:
// $ goa gen resume/design

package client

import (
	"encoding/json"
	"fmt"
	hycompany "resume/gen/hy_company"
	"strconv"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildCompanyListPayload builds the payload for the hy_company companyList
// endpoint from CLI flags.
func BuildCompanyListPayload(hyCompanyCompanyListToken string) (*hycompany.CompanyListPayload, error) {
	var token *string
	{
		if hyCompanyCompanyListToken != "" {
			token = &hyCompanyCompanyListToken
		}
	}
	v := &hycompany.CompanyListPayload{}
	v.Token = token

	return v, nil
}

// BuildGetCompanyPayload builds the payload for the hy_company getCompany
// endpoint from CLI flags.
func BuildGetCompanyPayload(hyCompanyGetCompanyCompanyID string, hyCompanyGetCompanyToken string) (*hycompany.GetCompanyPayload, error) {
	var err error
	var companyID int
	{
		var v int64
		v, err = strconv.ParseInt(hyCompanyGetCompanyCompanyID, 10, strconv.IntSize)
		companyID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for companyID, must be INT")
		}
		if companyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("companyID", companyID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyCompanyGetCompanyToken != "" {
			token = &hyCompanyGetCompanyToken
		}
	}
	v := &hycompany.GetCompanyPayload{}
	v.CompanyID = companyID
	v.Token = token

	return v, nil
}

// BuildCreateCompanyPayload builds the payload for the hy_company
// createCompany endpoint from CLI flags.
func BuildCreateCompanyPayload(hyCompanyCreateCompanyBody string, hyCompanyCreateCompanyToken string) (*hycompany.CreateCompanyPayload, error) {
	var err error
	var body CreateCompanyRequestBody
	{
		err = json.Unmarshal([]byte(hyCompanyCreateCompanyBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": \"Shinagawa Tokyo\",\n      \"company_name\": \"Company\",\n      \"country_id\": 110\n   }'")
		}
		if body.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", body.CountryID, 1, true))
		}
		if body.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", body.CountryID, 999, false))
		}
		if utf8.RuneCountInString(body.CompanyName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", body.CompanyName, utf8.RuneCountInString(body.CompanyName), 2, true))
		}
		if utf8.RuneCountInString(body.CompanyName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", body.CompanyName, utf8.RuneCountInString(body.CompanyName), 40, false))
		}
		if utf8.RuneCountInString(body.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", body.Address, utf8.RuneCountInString(body.Address), 2, true))
		}
		if utf8.RuneCountInString(body.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", body.Address, utf8.RuneCountInString(body.Address), 80, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyCompanyCreateCompanyToken != "" {
			token = &hyCompanyCreateCompanyToken
		}
	}
	v := &hycompany.CreateCompanyPayload{
		CountryID:   body.CountryID,
		CompanyName: body.CompanyName,
		Address:     body.Address,
	}
	v.Token = token

	return v, nil
}

// BuildUpdateCompanyPayload builds the payload for the hy_company
// updateCompany endpoint from CLI flags.
func BuildUpdateCompanyPayload(hyCompanyUpdateCompanyBody string, hyCompanyUpdateCompanyCompanyID string, hyCompanyUpdateCompanyToken string) (*hycompany.UpdateCompanyPayload, error) {
	var err error
	var body UpdateCompanyRequestBody
	{
		err = json.Unmarshal([]byte(hyCompanyUpdateCompanyBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": \"Shinagawa Tokyo\",\n      \"company_name\": \"Company\",\n      \"country_id\": 110\n   }'")
		}
		if body.CountryID != nil {
			if *body.CountryID < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 1, true))
			}
		}
		if body.CountryID != nil {
			if *body.CountryID > 999 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 999, false))
			}
		}
		if body.CompanyName != nil {
			if utf8.RuneCountInString(*body.CompanyName) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", *body.CompanyName, utf8.RuneCountInString(*body.CompanyName), 2, true))
			}
		}
		if body.CompanyName != nil {
			if utf8.RuneCountInString(*body.CompanyName) > 40 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", *body.CompanyName, utf8.RuneCountInString(*body.CompanyName), 40, false))
			}
		}
		if body.Address != nil {
			if utf8.RuneCountInString(*body.Address) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 2, true))
			}
		}
		if body.Address != nil {
			if utf8.RuneCountInString(*body.Address) > 80 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 80, false))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var companyID int
	{
		var v int64
		v, err = strconv.ParseInt(hyCompanyUpdateCompanyCompanyID, 10, strconv.IntSize)
		companyID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for companyID, must be INT")
		}
		if companyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("companyID", companyID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyCompanyUpdateCompanyToken != "" {
			token = &hyCompanyUpdateCompanyToken
		}
	}
	v := &hycompany.UpdateCompanyPayload{
		CountryID:   body.CountryID,
		CompanyName: body.CompanyName,
		Address:     body.Address,
	}
	v.CompanyID = companyID
	v.Token = token

	return v, nil
}

// BuildDeleteCompanyPayload builds the payload for the hy_company
// deleteCompany endpoint from CLI flags.
func BuildDeleteCompanyPayload(hyCompanyDeleteCompanyCompanyID string, hyCompanyDeleteCompanyToken string) (*hycompany.DeleteCompanyPayload, error) {
	var err error
	var companyID int
	{
		var v int64
		v, err = strconv.ParseInt(hyCompanyDeleteCompanyCompanyID, 10, strconv.IntSize)
		companyID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for companyID, must be INT")
		}
		if companyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("companyID", companyID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyCompanyDeleteCompanyToken != "" {
			token = &hyCompanyDeleteCompanyToken
		}
	}
	v := &hycompany.DeleteCompanyPayload{}
	v.CompanyID = companyID
	v.Token = token

	return v, nil
}
