// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/hiromaily/go-goa/design
// --out=$(GOPATH)/src/github.com/hiromaily/go-goa
// --version=v1.1.0-dirty
//
// API "api": health Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// HealthHealthPath computes a request path to the health action of health.
func HealthHealthPath() string {

	return fmt.Sprintf("/api/_ah/health")
}

// Perform health check.
func (c *Client) HealthHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewHealthHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHealthHealthRequest create the request corresponding to the health action endpoint of the health resource.
func (c *Client) NewHealthHealthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
