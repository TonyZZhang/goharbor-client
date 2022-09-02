// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExecHistory exec history
//
// swagger:model ExecHistory
type ExecHistory struct {

	// the creation time of purge job.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// if purge job was deleted.
	Deleted bool `json:"deleted,omitempty"`

	// the id of purge job.
	ID int64 `json:"id,omitempty"`

	// the job kind of purge job.
	JobKind string `json:"job_kind,omitempty"`

	// the job name of purge job.
	JobName string `json:"job_name,omitempty"`

	// the job parameters of purge job.
	JobParameters string `json:"job_parameters,omitempty"`

	// the status of purge job.
	JobStatus string `json:"job_status,omitempty"`

	// schedule
	Schedule *ScheduleObj `json:"schedule,omitempty"`

	// the update time of purge job.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this exec history
func (m *ExecHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecHistory) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExecHistory) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *ExecHistory) validateUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecHistory) UnmarshalBinary(b []byte) error {
	var res ExecHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
