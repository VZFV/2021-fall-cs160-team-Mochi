// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNoteMembersHandlerFunc turns a function with the right signature into a get note members handler
type GetNoteMembersHandlerFunc func(GetNoteMembersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNoteMembersHandlerFunc) Handle(params GetNoteMembersParams) middleware.Responder {
	return fn(params)
}

// GetNoteMembersHandler interface for that can handle valid get note members params
type GetNoteMembersHandler interface {
	Handle(GetNoteMembersParams) middleware.Responder
}

// NewGetNoteMembers creates a new http.Handler for the get note members operation
func NewGetNoteMembers(ctx *middleware.Context, handler GetNoteMembersHandler) *GetNoteMembers {
	return &GetNoteMembers{Context: ctx, Handler: handler}
}

/* GetNoteMembers swagger:route GET /v1/notes/{id}/members notesV1 getNoteMembers

get members of notes

get members of notes

*/
type GetNoteMembers struct {
	Context *middleware.Context
	Handler GetNoteMembersHandler
}

func (o *GetNoteMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetNoteMembersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}