// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionRuleMetadata the tag retention rule metadata
//
// swagger:model RetentionRuleMetadata
type RetentionRuleMetadata struct {

	// rule action
	Action string `json:"action,omitempty"`

	// rule display text
	DisplayText string `json:"display_text,omitempty"`

	// rule params
	Params []*RetentionRuleParamMetadata `json:"params"`

	// rule id
	RuleTemplate string `json:"rule_template,omitempty"`
}

// Validate validates this retention rule metadata
func (m *RetentionRuleMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetentionRuleMetadata) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for i := 0; i < len(m.Params); i++ {
		if swag.IsZero(m.Params[i]) { // not required
			continue
		}

		if m.Params[i] != nil {
			if err := m.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetentionRuleMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionRuleMetadata) UnmarshalBinary(b []byte) error {
	var res RetentionRuleMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
