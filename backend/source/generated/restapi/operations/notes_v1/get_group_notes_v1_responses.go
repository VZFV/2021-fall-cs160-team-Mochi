// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// GetGroupNotesV1OKCode is the HTTP code returned for type GetGroupNotesV1OK
const GetGroupNotesV1OKCode int = 200

/*GetGroupNotesV1OK Success

swagger:response getGroupNotesV1OK
*/
type GetGroupNotesV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.NotesGetResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1OK creates GetGroupNotesV1OK with default headers values
func NewGetGroupNotesV1OK() *GetGroupNotesV1OK {

	return &GetGroupNotesV1OK{}
}

// WithPayload adds the payload to the get group notes v1 o k response
func (o *GetGroupNotesV1OK) WithPayload(payload *models.NotesGetResponse) *GetGroupNotesV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 o k response
func (o *GetGroupNotesV1OK) SetPayload(payload *models.NotesGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1BadRequestCode is the HTTP code returned for type GetGroupNotesV1BadRequest
const GetGroupNotesV1BadRequestCode int = 400

/*GetGroupNotesV1BadRequest Bad Request

swagger:response getGroupNotesV1BadRequest
*/
type GetGroupNotesV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1BadRequest creates GetGroupNotesV1BadRequest with default headers values
func NewGetGroupNotesV1BadRequest() *GetGroupNotesV1BadRequest {

	return &GetGroupNotesV1BadRequest{}
}

// WithPayload adds the payload to the get group notes v1 bad request response
func (o *GetGroupNotesV1BadRequest) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 bad request response
func (o *GetGroupNotesV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1UnauthorizedCode is the HTTP code returned for type GetGroupNotesV1Unauthorized
const GetGroupNotesV1UnauthorizedCode int = 401

/*GetGroupNotesV1Unauthorized Unauthorized

swagger:response getGroupNotesV1Unauthorized
*/
type GetGroupNotesV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1Unauthorized creates GetGroupNotesV1Unauthorized with default headers values
func NewGetGroupNotesV1Unauthorized() *GetGroupNotesV1Unauthorized {

	return &GetGroupNotesV1Unauthorized{}
}

// WithPayload adds the payload to the get group notes v1 unauthorized response
func (o *GetGroupNotesV1Unauthorized) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 unauthorized response
func (o *GetGroupNotesV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1ForbiddenCode is the HTTP code returned for type GetGroupNotesV1Forbidden
const GetGroupNotesV1ForbiddenCode int = 403

/*GetGroupNotesV1Forbidden Forbidden

swagger:response getGroupNotesV1Forbidden
*/
type GetGroupNotesV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1Forbidden creates GetGroupNotesV1Forbidden with default headers values
func NewGetGroupNotesV1Forbidden() *GetGroupNotesV1Forbidden {

	return &GetGroupNotesV1Forbidden{}
}

// WithPayload adds the payload to the get group notes v1 forbidden response
func (o *GetGroupNotesV1Forbidden) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 forbidden response
func (o *GetGroupNotesV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1NotFoundCode is the HTTP code returned for type GetGroupNotesV1NotFound
const GetGroupNotesV1NotFoundCode int = 404

/*GetGroupNotesV1NotFound Not Found

swagger:response getGroupNotesV1NotFound
*/
type GetGroupNotesV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1NotFound creates GetGroupNotesV1NotFound with default headers values
func NewGetGroupNotesV1NotFound() *GetGroupNotesV1NotFound {

	return &GetGroupNotesV1NotFound{}
}

// WithPayload adds the payload to the get group notes v1 not found response
func (o *GetGroupNotesV1NotFound) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 not found response
func (o *GetGroupNotesV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1ConflictCode is the HTTP code returned for type GetGroupNotesV1Conflict
const GetGroupNotesV1ConflictCode int = 409

/*GetGroupNotesV1Conflict Conflict

swagger:response getGroupNotesV1Conflict
*/
type GetGroupNotesV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1Conflict creates GetGroupNotesV1Conflict with default headers values
func NewGetGroupNotesV1Conflict() *GetGroupNotesV1Conflict {

	return &GetGroupNotesV1Conflict{}
}

// WithPayload adds the payload to the get group notes v1 conflict response
func (o *GetGroupNotesV1Conflict) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 conflict response
func (o *GetGroupNotesV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGroupNotesV1InternalServerErrorCode is the HTTP code returned for type GetGroupNotesV1InternalServerError
const GetGroupNotesV1InternalServerErrorCode int = 500

/*GetGroupNotesV1InternalServerError Internal Server Error

swagger:response getGroupNotesV1InternalServerError
*/
type GetGroupNotesV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetGroupNotesV1InternalServerError creates GetGroupNotesV1InternalServerError with default headers values
func NewGetGroupNotesV1InternalServerError() *GetGroupNotesV1InternalServerError {

	return &GetGroupNotesV1InternalServerError{}
}

// WithPayload adds the payload to the get group notes v1 internal server error response
func (o *GetGroupNotesV1InternalServerError) WithPayload(payload *models.ErrResponse) *GetGroupNotesV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get group notes v1 internal server error response
func (o *GetGroupNotesV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGroupNotesV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
