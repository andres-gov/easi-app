// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/google/uuid"
)

// A document that belongs to an accessibility request
type AccessibilityRequestDocument struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	UploadedAt time.Time `json:"uploadedAt"`
}

// An edge of an AccessibilityRequestConnection
type AccessibilityRequestEdge struct {
	Cursor string                       `json:"cursor"`
	Node   *models.AccessibilityRequest `json:"node"`
}

// A collection of AccessibilityRequests
type AccessibilityRequestsConnection struct {
	Edges      []*AccessibilityRequestEdge `json:"edges"`
	TotalCount int                         `json:"totalCount"`
}

// Parameters required to create an AccessibilityRequest
type CreateAccessibilityRequestInput struct {
	IntakeID uuid.UUID `json:"intakeID"`
	Name     string    `json:"name"`
}

// Result of CreateAccessibilityRequest
type CreateAccessibilityRequestPayload struct {
	AccessibilityRequest *models.AccessibilityRequest `json:"accessibilityRequest"`
	UserErrors           []*UserError                 `json:"userErrors"`
}

// A collection of Systems
type SystemConnection struct {
	Edges      []*SystemEdge `json:"edges"`
	TotalCount int           `json:"totalCount"`
}

// An edge of an SystemConnection
type SystemEdge struct {
	Cursor string         `json:"cursor"`
	Node   *models.System `json:"node"`
}

// UserError represents application-level errors that are the result of
// either user or application developer error.
type UserError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}
