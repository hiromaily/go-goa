// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "api": hy_usertech Resource Client
//
// Command:
// $ goagen
// --design=github.com/hiromaily/go-goa/goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa/goa
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetUserDislikeTechHyUsertechPath computes a request path to the GetUserDislikeTech action of hy_usertech.
func GetUserDislikeTechHyUsertechPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/api/user/%s/disliketech", param0)
}

// Retrieve user's dislike techs.
func (c *Client) GetUserDislikeTechHyUsertech(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserDislikeTechHyUsertechRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserDislikeTechHyUsertechRequest create the request corresponding to the GetUserDislikeTech action endpoint of the hy_usertech resource.
func (c *Client) NewGetUserDislikeTechHyUsertechRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetUserLikeTechHyUsertechPath computes a request path to the GetUserLikeTech action of hy_usertech.
func GetUserLikeTechHyUsertechPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/api/user/%s/liketech", param0)
}

// Retrieve user's favorite techs.
func (c *Client) GetUserLikeTechHyUsertech(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserLikeTechHyUsertechRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserLikeTechHyUsertechRequest create the request corresponding to the GetUserLikeTech action endpoint of the hy_usertech resource.
func (c *Client) NewGetUserLikeTechHyUsertechRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
