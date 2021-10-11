// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// CreateGroupV1OKCode is the HTTP code returned for type CreateGroupV1OK
const CreateGroupV1OKCode int = 200

/*CreateGroupV1OK Success

swagger:response createGroupV1OK
*/
type CreateGroupV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.GroupResponse `json:"body,omitempty"`
}

// NewCreateGroupV1OK creates CreateGroupV1OK with default headers values
func NewCreateGroupV1OK() *CreateGroupV1OK {

	return &CreateGroupV1OK{}
}

// WithPayload adds the payload to the create group v1 o k response
func (o *CreateGroupV1OK) WithPayload(payload *models.GroupResponse) *CreateGroupV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 o k response
func (o *CreateGroupV1OK) SetPayload(payload *models.GroupResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1BadRequestCode is the HTTP code returned for type CreateGroupV1BadRequest
const CreateGroupV1BadRequestCode int = 400

/*CreateGroupV1BadRequest Bad Request

swagger:response createGroupV1BadRequest
*/
type CreateGroupV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1BadRequest creates CreateGroupV1BadRequest with default headers values
func NewCreateGroupV1BadRequest() *CreateGroupV1BadRequest {

	return &CreateGroupV1BadRequest{}
}

// WithPayload adds the payload to the create group v1 bad request response
func (o *CreateGroupV1BadRequest) WithPayload(payload *models.ErrResponse) *CreateGroupV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 bad request response
func (o *CreateGroupV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1UnauthorizedCode is the HTTP code returned for type CreateGroupV1Unauthorized
const CreateGroupV1UnauthorizedCode int = 401

/*CreateGroupV1Unauthorized Unauthorized

swagger:response createGroupV1Unauthorized
*/
type CreateGroupV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1Unauthorized creates CreateGroupV1Unauthorized with default headers values
func NewCreateGroupV1Unauthorized() *CreateGroupV1Unauthorized {

	return &CreateGroupV1Unauthorized{}
}

// WithPayload adds the payload to the create group v1 unauthorized response
func (o *CreateGroupV1Unauthorized) WithPayload(payload *models.ErrResponse) *CreateGroupV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 unauthorized response
func (o *CreateGroupV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1ForbiddenCode is the HTTP code returned for type CreateGroupV1Forbidden
const CreateGroupV1ForbiddenCode int = 403

/*CreateGroupV1Forbidden Forbidden

swagger:response createGroupV1Forbidden
*/
type CreateGroupV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1Forbidden creates CreateGroupV1Forbidden with default headers values
func NewCreateGroupV1Forbidden() *CreateGroupV1Forbidden {

	return &CreateGroupV1Forbidden{}
}

// WithPayload adds the payload to the create group v1 forbidden response
func (o *CreateGroupV1Forbidden) WithPayload(payload *models.ErrResponse) *CreateGroupV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 forbidden response
func (o *CreateGroupV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1NotFoundCode is the HTTP code returned for type CreateGroupV1NotFound
const CreateGroupV1NotFoundCode int = 404

/*CreateGroupV1NotFound Not Found

swagger:response createGroupV1NotFound
*/
type CreateGroupV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1NotFound creates CreateGroupV1NotFound with default headers values
func NewCreateGroupV1NotFound() *CreateGroupV1NotFound {

	return &CreateGroupV1NotFound{}
}

// WithPayload adds the payload to the create group v1 not found response
func (o *CreateGroupV1NotFound) WithPayload(payload *models.ErrResponse) *CreateGroupV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 not found response
func (o *CreateGroupV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1ConflictCode is the HTTP code returned for type CreateGroupV1Conflict
const CreateGroupV1ConflictCode int = 409

/*CreateGroupV1Conflict Conflict

swagger:response createGroupV1Conflict
*/
type CreateGroupV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1Conflict creates CreateGroupV1Conflict with default headers values
func NewCreateGroupV1Conflict() *CreateGroupV1Conflict {

	return &CreateGroupV1Conflict{}
}

// WithPayload adds the payload to the create group v1 conflict response
func (o *CreateGroupV1Conflict) WithPayload(payload *models.ErrResponse) *CreateGroupV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 conflict response
func (o *CreateGroupV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupV1InternalServerErrorCode is the HTTP code returned for type CreateGroupV1InternalServerError
const CreateGroupV1InternalServerErrorCode int = 500

/*CreateGroupV1InternalServerError Internal Server Error

swagger:response createGroupV1InternalServerError
*/
type CreateGroupV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewCreateGroupV1InternalServerError creates CreateGroupV1InternalServerError with default headers values
func NewCreateGroupV1InternalServerError() *CreateGroupV1InternalServerError {

	return &CreateGroupV1InternalServerError{}
}

// WithPayload adds the payload to the create group v1 internal server error response
func (o *CreateGroupV1InternalServerError) WithPayload(payload *models.ErrResponse) *CreateGroupV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group v1 internal server error response
func (o *CreateGroupV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}