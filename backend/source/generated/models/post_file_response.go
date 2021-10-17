// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostFileResponse post file response
//
// swagger:model postFileResponse
type PostFileResponse struct {

	// path of file
	NoteReference string `json:"note_reference,omitempty"`
}

// Validate validates this post file response
func (m *PostFileResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post file response based on context it is used
func (m *PostFileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostFileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostFileResponse) UnmarshalBinary(b []byte) error {
	var res PostFileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
