///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFunctionsHandlerFunc turns a function with the right signature into a get functions handler
type GetFunctionsHandlerFunc func(GetFunctionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFunctionsHandlerFunc) Handle(params GetFunctionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFunctionsHandler interface for that can handle valid get functions params
type GetFunctionsHandler interface {
	Handle(GetFunctionsParams, interface{}) middleware.Responder
}

// NewGetFunctions creates a new http.Handler for the get functions operation
func NewGetFunctions(ctx *middleware.Context, handler GetFunctionsHandler) *GetFunctions {
	return &GetFunctions{Context: ctx, Handler: handler}
}

/*GetFunctions swagger:route GET / Store getFunctions

List all existing functions

*/
type GetFunctions struct {
	Context *middleware.Context
	Handler GetFunctionsHandler
}

func (o *GetFunctions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFunctionsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
