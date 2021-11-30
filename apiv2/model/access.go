// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Access access
//
// swagger:model Access
type Access struct {

	// The action of the access
	Action string `json:"action,omitempty"`

	// The effect of the access
	Effect string `json:"effect,omitempty"`

	// The resource of the access
	Resource string `json:"resource,omitempty"`
}

// Validate validates this access
func (m *Access) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Access) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Access) UnmarshalBinary(b []byte) error {
	var res Access
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
