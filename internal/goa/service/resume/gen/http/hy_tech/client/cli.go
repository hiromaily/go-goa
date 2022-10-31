// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_tech HTTP client CLI support package
//
// Command:
// $ goa gen resume/design

package client

import (
	"encoding/json"
	"fmt"
	hytech "resume/gen/hy_tech"
	"strconv"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// BuildTechListPayload builds the payload for the hy_tech techList endpoint
// from CLI flags.
func BuildTechListPayload(hyTechTechListToken string) (*hytech.TechListPayload, error) {
	var token *string
	{
		if hyTechTechListToken != "" {
			token = &hyTechTechListToken
		}
	}
	v := &hytech.TechListPayload{}
	v.Token = token

	return v, nil
}

// BuildGetTechPayload builds the payload for the hy_tech getTech endpoint from
// CLI flags.
func BuildGetTechPayload(hyTechGetTechTechID string, hyTechGetTechToken string) (*hytech.GetTechPayload, error) {
	var err error
	var techID int
	{
		var v int64
		v, err = strconv.ParseInt(hyTechGetTechTechID, 10, strconv.IntSize)
		techID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for techID, must be INT")
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyTechGetTechToken != "" {
			token = &hyTechGetTechToken
		}
	}
	v := &hytech.GetTechPayload{}
	v.TechID = techID
	v.Token = token

	return v, nil
}

// BuildCreateTechPayload builds the payload for the hy_tech createTech
// endpoint from CLI flags.
func BuildCreateTechPayload(hyTechCreateTechBody string, hyTechCreateTechToken string) (*hytech.CreateTechPayload, error) {
	var err error
	var body CreateTechRequestBody
	{
		err = json.Unmarshal([]byte(hyTechCreateTechBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"tech_name\": \"Golang\"\n   }'")
		}
		if utf8.RuneCountInString(body.TechName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.tech_name", body.TechName, utf8.RuneCountInString(body.TechName), 1, true))
		}
		if utf8.RuneCountInString(body.TechName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.tech_name", body.TechName, utf8.RuneCountInString(body.TechName), 40, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyTechCreateTechToken != "" {
			token = &hyTechCreateTechToken
		}
	}
	v := &hytech.CreateTechPayload{
		TechName: body.TechName,
	}
	v.Token = token

	return v, nil
}

// BuildUpdateTechPayload builds the payload for the hy_tech updateTech
// endpoint from CLI flags.
func BuildUpdateTechPayload(hyTechUpdateTechBody string, hyTechUpdateTechTechID string, hyTechUpdateTechToken string) (*hytech.UpdateTechPayload, error) {
	var err error
	var body UpdateTechRequestBody
	{
		err = json.Unmarshal([]byte(hyTechUpdateTechBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"tech_name\": \"Golang\"\n   }'")
		}
		if utf8.RuneCountInString(body.TechName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.tech_name", body.TechName, utf8.RuneCountInString(body.TechName), 1, true))
		}
		if utf8.RuneCountInString(body.TechName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.tech_name", body.TechName, utf8.RuneCountInString(body.TechName), 40, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var techID int
	{
		var v int64
		v, err = strconv.ParseInt(hyTechUpdateTechTechID, 10, strconv.IntSize)
		techID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for techID, must be INT")
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyTechUpdateTechToken != "" {
			token = &hyTechUpdateTechToken
		}
	}
	v := &hytech.UpdateTechPayload{
		TechName: body.TechName,
	}
	v.TechID = techID
	v.Token = token

	return v, nil
}

// BuildDeleteTechPayload builds the payload for the hy_tech deleteTech
// endpoint from CLI flags.
func BuildDeleteTechPayload(hyTechDeleteTechTechID string, hyTechDeleteTechToken string) (*hytech.DeleteTechPayload, error) {
	var err error
	var techID int
	{
		var v int64
		v, err = strconv.ParseInt(hyTechDeleteTechTechID, 10, strconv.IntSize)
		techID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for techID, must be INT")
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if hyTechDeleteTechToken != "" {
			token = &hyTechDeleteTechToken
		}
	}
	v := &hytech.DeleteTechPayload{}
	v.TechID = techID
	v.Token = token

	return v, nil
}
