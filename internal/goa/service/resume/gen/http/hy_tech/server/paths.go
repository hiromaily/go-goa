// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the hy_tech service.
//
// Command:
// $ goa gen resume/design

package server

import (
	"fmt"
)

// TechListHyTechPath returns the URL path to the hy_tech service techList HTTP endpoint.
func TechListHyTechPath() string {
	return "/tech"
}

// GetTechHyTechPath returns the URL path to the hy_tech service getTech HTTP endpoint.
func GetTechHyTechPath(techID int) string {
	return fmt.Sprintf("/tech/%v", techID)
}

// CreateTechHyTechPath returns the URL path to the hy_tech service createTech HTTP endpoint.
func CreateTechHyTechPath() string {
	return "/tech"
}

// UpdateTechHyTechPath returns the URL path to the hy_tech service updateTech HTTP endpoint.
func UpdateTechHyTechPath(techID int) string {
	return fmt.Sprintf("/tech/%v", techID)
}

// DeleteTechHyTechPath returns the URL path to the hy_tech service deleteTech HTTP endpoint.
func DeleteTechHyTechPath(techID int) string {
	return fmt.Sprintf("/tech/%v", techID)
}
