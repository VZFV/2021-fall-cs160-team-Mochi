// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewFindNotesByUsernameParams creates a new FindNotesByUsernameParams object
//
// There are no default values defined in the spec.
func NewFindNotesByUsernameParams() FindNotesByUsernameParams {

	return FindNotesByUsernameParams{}
}

// FindNotesByUsernameParams contains all the bound params for the find notes by username operation
// typically these are obtained from a http.Request
//
// swagger:parameters findNotesByUsername
type FindNotesByUsernameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Bearer token based Authorization
	  Required: true
	  In: header
	*/
	Authorization string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindNotesByUsernameParams() beforehand.
func (o *FindNotesByUsernameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *FindNotesByUsernameParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Authorization", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Authorization", "header", raw); err != nil {
		return err
	}
	o.Authorization = raw

	return nil
}
