///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

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

// NewGetFunctionsParams creates a new GetFunctionsParams object
// with the default values initialized.
func NewGetFunctionsParams() *GetFunctionsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetFunctionsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFunctionsParamsWithTimeout creates a new GetFunctionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFunctionsParamsWithTimeout(timeout time.Duration) *GetFunctionsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetFunctionsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewGetFunctionsParamsWithContext creates a new GetFunctionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFunctionsParamsWithContext(ctx context.Context) *GetFunctionsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetFunctionsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewGetFunctionsParamsWithHTTPClient creates a new GetFunctionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFunctionsParamsWithHTTPClient(client *http.Client) *GetFunctionsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetFunctionsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*GetFunctionsParams contains all the parameters to send to the API endpoint
for the get functions operation typically these are written to a http.Request
*/
type GetFunctionsParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*State
	  Function state

	*/
	State *string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get functions params
func (o *GetFunctionsParams) WithTimeout(timeout time.Duration) *GetFunctionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get functions params
func (o *GetFunctionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get functions params
func (o *GetFunctionsParams) WithContext(ctx context.Context) *GetFunctionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get functions params
func (o *GetFunctionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get functions params
func (o *GetFunctionsParams) WithHTTPClient(client *http.Client) *GetFunctionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get functions params
func (o *GetFunctionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the get functions params
func (o *GetFunctionsParams) WithXDispatchOrg(xDispatchOrg *string) *GetFunctionsParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the get functions params
func (o *GetFunctionsParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the get functions params
func (o *GetFunctionsParams) WithXDispatchProject(xDispatchProject *string) *GetFunctionsParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the get functions params
func (o *GetFunctionsParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithState adds the state to the get functions params
func (o *GetFunctionsParams) WithState(state *string) *GetFunctionsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get functions params
func (o *GetFunctionsParams) SetState(state *string) {
	o.State = state
}

// WithTags adds the tags to the get functions params
func (o *GetFunctionsParams) WithTags(tags []string) *GetFunctionsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get functions params
func (o *GetFunctionsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetFunctionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDispatchOrg != nil {

		// header param X-Dispatch-Org
		if err := r.SetHeaderParam("X-Dispatch-Org", *o.XDispatchOrg); err != nil {
			return err
		}

	}

	if o.XDispatchProject != nil {

		// header param X-Dispatch-Project
		if err := r.SetHeaderParam("X-Dispatch-Project", *o.XDispatchProject); err != nil {
			return err
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
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