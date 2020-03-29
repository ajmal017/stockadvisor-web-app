// Code generated by go-swagger; DO NOT EDIT.

package news_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostNewsQueryHandlerFunc turns a function with the right signature into a post news query handler
type PostNewsQueryHandlerFunc func(PostNewsQueryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostNewsQueryHandlerFunc) Handle(params PostNewsQueryParams) middleware.Responder {
	return fn(params)
}

// PostNewsQueryHandler interface for that can handle valid post news query params
type PostNewsQueryHandler interface {
	Handle(PostNewsQueryParams) middleware.Responder
}

// NewPostNewsQuery creates a new http.Handler for the post news query operation
func NewPostNewsQuery(ctx *middleware.Context, handler PostNewsQueryHandler) *PostNewsQuery {
	return &PostNewsQuery{Context: ctx, Handler: handler}
}

/*PostNewsQuery swagger:route POST /news/query News API postNewsQuery

Retrieves a news query

Retrieves a news query

*/
type PostNewsQuery struct {
	Context *middleware.Context
	Handler PostNewsQueryHandler
}

func (o *PostNewsQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostNewsQueryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
