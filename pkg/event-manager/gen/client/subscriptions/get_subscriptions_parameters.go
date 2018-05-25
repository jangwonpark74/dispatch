///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewGetSubscriptionsParams creates a new GetSubscriptionsParams object
// with the default values initialized.
func NewGetSubscriptionsParams() *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriptionsParamsWithTimeout creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriptionsParamsWithTimeout(timeout time.Duration) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetSubscriptionsParamsWithContext creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriptionsParamsWithContext(ctx context.Context) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetSubscriptionsParamsWithHTTPClient creates a new GetSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriptionsParamsWithHTTPClient(client *http.Client) *GetSubscriptionsParams {
	var ()
	return &GetSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetSubscriptionsParams contains all the parameters to send to the API endpoint
for the get subscriptions operation typically these are written to a http.Request
*/
type GetSubscriptionsParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subscriptions params
func (o *GetSubscriptionsParams) WithTimeout(timeout time.Duration) *GetSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscriptions params
func (o *GetSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscriptions params
func (o *GetSubscriptionsParams) WithContext(ctx context.Context) *GetSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscriptions params
func (o *GetSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscriptions params
func (o *GetSubscriptionsParams) WithHTTPClient(client *http.Client) *GetSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscriptions params
func (o *GetSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the get subscriptions params
func (o *GetSubscriptionsParams) WithXDispatchOrg(xDispatchOrg string) *GetSubscriptionsParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the get subscriptions params
func (o *GetSubscriptionsParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithTags adds the tags to the get subscriptions params
func (o *GetSubscriptionsParams) WithTags(tags []string) *GetSubscriptionsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get subscriptions params
func (o *GetSubscriptionsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
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
