// Code generated by goa v3.10.2, DO NOT EDIT.
//
// HTTP request path constructors for the hy_company service.
//
// Command:
// $ goa gen resume/design

package client

import (
	"fmt"
)

// CompanyListHyCompanyPath returns the URL path to the hy_company service companyList HTTP endpoint.
func CompanyListHyCompanyPath() string {
	return "/api/company"
}

// GetCompanyHyCompanyPath returns the URL path to the hy_company service getCompany HTTP endpoint.
func GetCompanyHyCompanyPath(companyID int) string {
	return fmt.Sprintf("/api/company/%v", companyID)
}

// CreateCompanyHyCompanyPath returns the URL path to the hy_company service createCompany HTTP endpoint.
func CreateCompanyHyCompanyPath() string {
	return "/api/company"
}

// UpdateCompanyHyCompanyPath returns the URL path to the hy_company service updateCompany HTTP endpoint.
func UpdateCompanyHyCompanyPath(companyID int) string {
	return fmt.Sprintf("/api/company/%v", companyID)
}

// DeleteCompanyHyCompanyPath returns the URL path to the hy_company service deleteCompany HTTP endpoint.
func DeleteCompanyHyCompanyPath(companyID int) string {
	return fmt.Sprintf("/api/company/%v", companyID)
}
