///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAPIParams creates a new DeleteAPIParams object
// with the default values initialized.
func NewDeleteAPIParams() *DeleteAPIParams {
	var ()
	return &DeleteAPIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIParamsWithTimeout creates a new DeleteAPIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIParamsWithTimeout(timeout time.Duration) *DeleteAPIParams {
	var ()
	return &DeleteAPIParams{

		timeout: timeout,
	}
}

// NewDeleteAPIParamsWithContext creates a new DeleteAPIParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIParamsWithContext(ctx context.Context) *DeleteAPIParams {
	var ()
	return &DeleteAPIParams{

		Context: ctx,
	}
}

// NewDeleteAPIParamsWithHTTPClient creates a new DeleteAPIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIParamsWithHTTPClient(client *http.Client) *DeleteAPIParams {
	var ()
	return &DeleteAPIParams{
		HTTPClient: client,
	}
}

/*DeleteAPIParams contains all the parameters to send to the API endpoint
for the delete API operation typically these are written to a http.Request
*/
type DeleteAPIParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*API
	  Name of API to work on

	*/
	API string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API params
func (o *DeleteAPIParams) WithTimeout(timeout time.Duration) *DeleteAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API params
func (o *DeleteAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API params
func (o *DeleteAPIParams) WithContext(ctx context.Context) *DeleteAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API params
func (o *DeleteAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API params
func (o *DeleteAPIParams) WithHTTPClient(client *http.Client) *DeleteAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API params
func (o *DeleteAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the delete API params
func (o *DeleteAPIParams) WithXDispatchOrg(xDispatchOrg string) *DeleteAPIParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the delete API params
func (o *DeleteAPIParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithAPI adds the api to the delete API params
func (o *DeleteAPIParams) WithAPI(api string) *DeleteAPIParams {
	o.SetAPI(api)
	return o
}

// SetAPI adds the api to the delete API params
func (o *DeleteAPIParams) SetAPI(api string) {
	o.API = api
}

// WithTags adds the tags to the delete API params
func (o *DeleteAPIParams) WithTags(tags []string) *DeleteAPIParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete API params
func (o *DeleteAPIParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	// path param api
	if err := r.SetPathParam("api", o.API); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
