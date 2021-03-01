// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PlaceOrderHandlerFunc turns a function with the right signature into a place order handler
type PlaceOrderHandlerFunc func(PlaceOrderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PlaceOrderHandlerFunc) Handle(params PlaceOrderParams) middleware.Responder {
	return fn(params)
}

// PlaceOrderHandler interface for that can handle valid place order params
type PlaceOrderHandler interface {
	Handle(PlaceOrderParams) middleware.Responder
}

// NewPlaceOrder creates a new http.Handler for the place order operation
func NewPlaceOrder(ctx *middleware.Context, handler PlaceOrderHandler) *PlaceOrder {
	return &PlaceOrder{Context: ctx, Handler: handler}
}

/* PlaceOrder swagger:route POST /store/order store placeOrder

Place an order for a pet

*/
type PlaceOrder struct {
	Context *middleware.Context
	Handler PlaceOrderHandler
}

func (o *PlaceOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPlaceOrderParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
