///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateFunctionReader is a Reader for the UpdateFunction structure.
type UpdateFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateFunctionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateFunctionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFunctionOK creates a UpdateFunctionOK with default headers values
func NewUpdateFunctionOK() *UpdateFunctionOK {
	return &UpdateFunctionOK{}
}

/*UpdateFunctionOK handles this case with default header values.

Successful update
*/
type UpdateFunctionOK struct {
	Payload *v1.Function
}

func (o *UpdateFunctionOK) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunctionOK  %+v", 200, o.Payload)
}

func (o *UpdateFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionBadRequest creates a UpdateFunctionBadRequest with default headers values
func NewUpdateFunctionBadRequest() *UpdateFunctionBadRequest {
	return &UpdateFunctionBadRequest{}
}

/*UpdateFunctionBadRequest handles this case with default header values.

Invalid input
*/
type UpdateFunctionBadRequest struct {
	Payload *v1.Error
}

func (o *UpdateFunctionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionUnauthorized creates a UpdateFunctionUnauthorized with default headers values
func NewUpdateFunctionUnauthorized() *UpdateFunctionUnauthorized {
	return &UpdateFunctionUnauthorized{}
}

/*UpdateFunctionUnauthorized handles this case with default header values.

Unauthorized Request
*/
type UpdateFunctionUnauthorized struct {
	Payload *v1.Error
}

func (o *UpdateFunctionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunctionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateFunctionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionForbidden creates a UpdateFunctionForbidden with default headers values
func NewUpdateFunctionForbidden() *UpdateFunctionForbidden {
	return &UpdateFunctionForbidden{}
}

/*UpdateFunctionForbidden handles this case with default header values.

access to this resource is forbidden
*/
type UpdateFunctionForbidden struct {
	Payload *v1.Error
}

func (o *UpdateFunctionForbidden) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunctionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateFunctionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionNotFound creates a UpdateFunctionNotFound with default headers values
func NewUpdateFunctionNotFound() *UpdateFunctionNotFound {
	return &UpdateFunctionNotFound{}
}

/*UpdateFunctionNotFound handles this case with default header values.

Function not found
*/
type UpdateFunctionNotFound struct {
	Payload *v1.Error
}

func (o *UpdateFunctionNotFound) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunctionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionDefault creates a UpdateFunctionDefault with default headers values
func NewUpdateFunctionDefault(code int) *UpdateFunctionDefault {
	return &UpdateFunctionDefault{
		_statusCode: code,
	}
}

/*UpdateFunctionDefault handles this case with default header values.

Unknown error
*/
type UpdateFunctionDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the update function default response
func (o *UpdateFunctionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFunctionDefault) Error() string {
	return fmt.Sprintf("[PUT /function/{functionName}][%d] updateFunction default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
