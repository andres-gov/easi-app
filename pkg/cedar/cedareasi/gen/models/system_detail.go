// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemDetail system detail
// swagger:model SystemDetail
type SystemDetail struct {

	// business function alignments
	BusinessFunctionAlignments []string `json:"business_function_alignments"`

	// business owner
	BusinessOwner string `json:"business_owner,omitempty"`

	// business owner org
	BusinessOwnerOrg string `json:"business_owner_org,omitempty"`

	// contractor ftes
	ContractorFtes int32 `json:"contractor_ftes,omitempty"`

	// csp hosted
	CspHosted bool `json:"csp_hosted,omitempty"`

	// csp names
	CspNames []string `json:"csp_names"`

	// csp service model type
	CspServiceModelType string `json:"csp_service_model_type,omitempty"`

	// data categories
	DataCategories []string `json:"data_categories"`

	// data center names
	DataCenterNames []string `json:"data_center_names"`

	// federal ftes
	FederalFtes int32 `json:"federal_ftes,omitempty"`

	// yyyy-MM-dd
	GprChangesInNextRelease bool `json:"gpr_changes_in_next_release,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is business application
	IsBusinessApplication bool `json:"is_business_application,omitempty"`

	// is cms owned
	IsCmsOwned bool `json:"is_cms_owned,omitempty"`

	// is parent system
	IsParentSystem bool `json:"is_parent_system,omitempty"`

	// isso
	Isso string `json:"isso,omitempty"`

	// major program alignments
	MajorProgramAlignments []string `json:"major_program_alignments"`

	// mission essential functions
	MissionEssentialFunctions []*MissionEssentialFunction `json:"mission_essential_functions"`

	// yyyy-MM-dd
	// Format: date
	NextMajorProjectDate strfmt.Date `json:"next_major_project_date,omitempty"`

	// yyyy-MM-dd
	// Format: date
	NextPlannedProdReleaseDate strfmt.Date `json:"next_planned_prod_release_date,omitempty"`

	// parent system name
	ParentSystemName string `json:"parent_system_name,omitempty"`

	// product owner
	ProductOwner string `json:"product_owner,omitempty"`

	// software products
	// Required: true
	SoftwareProducts []*SoftwareProduct `json:"software_products"`

	// supported user count
	SupportedUserCount int32 `json:"supported_user_count,omitempty"`

	// system acronym
	// Required: true
	SystemAcronym *string `json:"system_acronym"`

	// system classification
	SystemClassification string `json:"system_classification,omitempty"`

	// system maintainer org
	SystemMaintainerOrg string `json:"system_maintainer_org,omitempty"`

	// system name
	// Required: true
	SystemName *string `json:"system_name"`

	// system owner
	SystemOwner string `json:"system_owner,omitempty"`

	// system state
	SystemState string `json:"system_state,omitempty"`

	// system type
	SystemType string `json:"system_type,omitempty"`

	// tlc profile reviewer
	TlcProfileReviewer string `json:"tlc_profile_reviewer,omitempty"`
}

// Validate validates this system detail
func (m *SystemDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissionEssentialFunctions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextMajorProjectDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPlannedProdReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemAcronym(formats); err != nil {
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

func (m *SystemDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SystemDetail) validateMissionEssentialFunctions(formats strfmt.Registry) error {

	if swag.IsZero(m.MissionEssentialFunctions) { // not required
		return nil
	}

	for i := 0; i < len(m.MissionEssentialFunctions); i++ {
		if swag.IsZero(m.MissionEssentialFunctions[i]) { // not required
			continue
		}

		if m.MissionEssentialFunctions[i] != nil {
			if err := m.MissionEssentialFunctions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mission_essential_functions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemDetail) validateNextMajorProjectDate(formats strfmt.Registry) error {

	if swag.IsZero(m.NextMajorProjectDate) { // not required
		return nil
	}

	if err := validate.FormatOf("next_major_project_date", "body", "date", m.NextMajorProjectDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SystemDetail) validateNextPlannedProdReleaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.NextPlannedProdReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("next_planned_prod_release_date", "body", "date", m.NextPlannedProdReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SystemDetail) validateSoftwareProducts(formats strfmt.Registry) error {

	if err := validate.Required("software_products", "body", m.SoftwareProducts); err != nil {
		return err
	}

	for i := 0; i < len(m.SoftwareProducts); i++ {
		if swag.IsZero(m.SoftwareProducts[i]) { // not required
			continue
		}

		if m.SoftwareProducts[i] != nil {
			if err := m.SoftwareProducts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("software_products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SystemDetail) validateSystemAcronym(formats strfmt.Registry) error {

	if err := validate.Required("system_acronym", "body", m.SystemAcronym); err != nil {
		return err
	}

	return nil
}

func (m *SystemDetail) validateSystemName(formats strfmt.Registry) error {

	if err := validate.Required("system_name", "body", m.SystemName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDetail) UnmarshalBinary(b []byte) error {
	var res SystemDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}