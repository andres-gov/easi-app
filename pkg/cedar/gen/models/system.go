// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// System system
// swagger:model System
type System struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// is business application
	IsBusinessApplication bool `json:"is_business_application,omitempty"`

	// is cms owned
	IsCmsOwned bool `json:"is_cms_owned,omitempty"`

	// is parent system
	IsParentSystem bool `json:"is_parent_system,omitempty"`

	// parent system name
	ParentSystemName string `json:"parent_system_name,omitempty"`

	// system acronym
	SystemAcronym string `json:"system_acronym,omitempty"`

	// system classification
	SystemClassification string `json:"system_classification,omitempty"`

	// system name
	// Required: true
	SystemName *string `json:"system_name"`

	// system state
	SystemState string `json:"system_state,omitempty"`

	// system type
	SystemType string `json:"system_type,omitempty"`
}

// Validate validates this system
func (m *System) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *System) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *System) validateSystemName(formats strfmt.Registry) error {

	if err := validate.Required("system_name", "body", m.SystemName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *System) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *System) UnmarshalBinary(b []byte) error {
	var res System
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
