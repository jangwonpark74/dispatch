///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// EmitEventOKCode is the HTTP code returned for type EmitEventOK
const EmitEventOKCode int = 200

/*EmitEventOK Event emitted

swagger:response emitEventOK
*/
type EmitEventOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Emission `json:"body,omitempty"`
}

// NewEmitEventOK creates EmitEventOK with default headers values
func NewEmitEventOK() *EmitEventOK {

	return &EmitEventOK{}
}

// WithPayload adds the payload to the emit event o k response
func (o *EmitEventOK) WithPayload(payload *v1.Emission) *EmitEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the emit event o k response
func (o *EmitEventOK) SetPayload(payload *v1.Emission) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EmitEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EmitEventBadRequestCode is the HTTP code returned for type EmitEventBadRequest
const EmitEventBadRequestCode int = 400

/*EmitEventBadRequest Invalid input

swagger:response emitEventBadRequest
*/
type EmitEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewEmitEventBadRequest creates EmitEventBadRequest with default headers values
func NewEmitEventBadRequest() *EmitEventBadRequest {

	return &EmitEventBadRequest{}
}

// WithPayload adds the payload to the emit event bad request response
func (o *EmitEventBadRequest) WithPayload(payload *v1.Error) *EmitEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the emit event bad request response
func (o *EmitEventBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EmitEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EmitEventUnauthorizedCode is the HTTP code returned for type EmitEventUnauthorized
const EmitEventUnauthorizedCode int = 401

/*EmitEventUnauthorized Unauthorized Request

swagger:response emitEventUnauthorized
*/
type EmitEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewEmitEventUnauthorized creates EmitEventUnauthorized with default headers values
func NewEmitEventUnauthorized() *EmitEventUnauthorized {

	return &EmitEventUnauthorized{}
}

// WithPayload adds the payload to the emit event unauthorized response
func (o *EmitEventUnauthorized) WithPayload(payload *v1.Error) *EmitEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the emit event unauthorized response
func (o *EmitEventUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EmitEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EmitEventForbiddenCode is the HTTP code returned for type EmitEventForbidden
const EmitEventForbiddenCode int = 403

/*EmitEventForbidden access to this resource is forbidden

swagger:response emitEventForbidden
*/
type EmitEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewEmitEventForbidden creates EmitEventForbidden with default headers values
func NewEmitEventForbidden() *EmitEventForbidden {

	return &EmitEventForbidden{}
}

// WithPayload adds the payload to the emit event forbidden response
func (o *EmitEventForbidden) WithPayload(payload *v1.Error) *EmitEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the emit event forbidden response
func (o *EmitEventForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EmitEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*EmitEventDefault Unknown error

swagger:response emitEventDefault
*/
type EmitEventDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewEmitEventDefault creates EmitEventDefault with default headers values
func NewEmitEventDefault(code int) *EmitEventDefault {
	if code <= 0 {
		code = 500
	}

	return &EmitEventDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the emit event default response
func (o *EmitEventDefault) WithStatusCode(code int) *EmitEventDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the emit event default response
func (o *EmitEventDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the emit event default response
func (o *EmitEventDefault) WithPayload(payload *v1.Error) *EmitEventDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the emit event default response
func (o *EmitEventDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EmitEventDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
