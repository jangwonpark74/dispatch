///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

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

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewRunFunctionParams creates a new RunFunctionParams object
// with the default values initialized.
func NewRunFunctionParams() *RunFunctionParams {
	var ()
	return &RunFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunFunctionParamsWithTimeout creates a new RunFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunFunctionParamsWithTimeout(timeout time.Duration) *RunFunctionParams {
	var ()
	return &RunFunctionParams{

		timeout: timeout,
	}
}

// NewRunFunctionParamsWithContext creates a new RunFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunFunctionParamsWithContext(ctx context.Context) *RunFunctionParams {
	var ()
	return &RunFunctionParams{

		Context: ctx,
	}
}

// NewRunFunctionParamsWithHTTPClient creates a new RunFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunFunctionParamsWithHTTPClient(client *http.Client) *RunFunctionParams {
	var ()
	return &RunFunctionParams{
		HTTPClient: client,
	}
}

/*RunFunctionParams contains all the parameters to send to the API endpoint
for the run function operation typically these are written to a http.Request
*/
type RunFunctionParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*Body*/
	Body *v1.Run
	/*FunctionName
	  Name of function to run or retreive runs for

	*/
	FunctionName *string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the run function params
func (o *RunFunctionParams) WithTimeout(timeout time.Duration) *RunFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run function params
func (o *RunFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run function params
func (o *RunFunctionParams) WithContext(ctx context.Context) *RunFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run function params
func (o *RunFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run function params
func (o *RunFunctionParams) WithHTTPClient(client *http.Client) *RunFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run function params
func (o *RunFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the run function params
func (o *RunFunctionParams) WithXDispatchOrg(xDispatchOrg string) *RunFunctionParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the run function params
func (o *RunFunctionParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithBody adds the body to the run function params
func (o *RunFunctionParams) WithBody(body *v1.Run) *RunFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the run function params
func (o *RunFunctionParams) SetBody(body *v1.Run) {
	o.Body = body
}

// WithFunctionName adds the functionName to the run function params
func (o *RunFunctionParams) WithFunctionName(functionName *string) *RunFunctionParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the run function params
func (o *RunFunctionParams) SetFunctionName(functionName *string) {
	o.FunctionName = functionName
}

// WithTags adds the tags to the run function params
func (o *RunFunctionParams) WithTags(tags []string) *RunFunctionParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the run function params
func (o *RunFunctionParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *RunFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.FunctionName != nil {

		// query param functionName
		var qrFunctionName string
		if o.FunctionName != nil {
			qrFunctionName = *o.FunctionName
		}
		qFunctionName := qrFunctionName
		if qFunctionName != "" {
			if err := r.SetQueryParam("functionName", qFunctionName); err != nil {
				return err
			}
		}

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
