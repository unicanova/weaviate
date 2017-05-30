package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Event event
// swagger:model Event
type Event struct {

	// command
	Command *EventCommand `json:"command,omitempty"`

	// Command id.
	CommandID strfmt.UUID `json:"commandId,omitempty"`

	// Time the event was generated in milliseconds since epoch UTC.
	TimeMs int64 `json:"timeMs,omitempty"`

	// User that caused the event (if applicable).
	Userkey string `json:"userkey,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	if m.Command != nil {

		if err := m.Command.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("command")
			}
			return err
		}
	}

	return nil
}

// EventCommand event command
// swagger:model EventCommand
type EventCommand struct {

	// command parameters
	CommandParameters *CommandParameters `json:"commandParameters,omitempty"`
}

// Validate validates this event command
func (m *EventCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}