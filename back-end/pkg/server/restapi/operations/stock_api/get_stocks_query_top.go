// Code generated by go-swagger; DO NOT EDIT.

package stock_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStocksQueryTopHandlerFunc turns a function with the right signature into a get stocks query top handler
type GetStocksQueryTopHandlerFunc func(GetStocksQueryTopParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStocksQueryTopHandlerFunc) Handle(params GetStocksQueryTopParams) middleware.Responder {
	return fn(params)
}

// GetStocksQueryTopHandler interface for that can handle valid get stocks query top params
type GetStocksQueryTopHandler interface {
	Handle(GetStocksQueryTopParams) middleware.Responder
}

// NewGetStocksQueryTop creates a new http.Handler for the get stocks query top operation
func NewGetStocksQueryTop(ctx *middleware.Context, handler GetStocksQueryTopHandler) *GetStocksQueryTop {
	return &GetStocksQueryTop{Context: ctx, Handler: handler}
}

/*GetStocksQueryTop swagger:route GET /stocks/queryTop Stock API getStocksQueryTop

Querys information about the top N stocks

Querys information about the top N stocks

*/
type GetStocksQueryTop struct {
	Context *middleware.Context
	Handler GetStocksQueryTopHandler
}

func (o *GetStocksQueryTop) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStocksQueryTopParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
