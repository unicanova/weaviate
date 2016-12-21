package models


// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AdaptersActivateResponse adapters activate response
// swagger:model AdaptersActivateResponse
type AdaptersActivateResponse struct {

	// An ID that represents the link between the adapter and the user. The activationUrl will give this ID to the adapter for use in the accept API.
	ActivationID string `json:"activationId,omitempty"`

	// A URL to the adapter where the user should go to complete the activation. The URL contains the activationId and the user's email address.
	ActivationURL string `json:"activationUrl,omitempty"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#adaptersActivateResponse".
	Kind *string `json:"kind,omitempty"`
}

// Validate validates this adapters activate response
func (m *AdaptersActivateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}