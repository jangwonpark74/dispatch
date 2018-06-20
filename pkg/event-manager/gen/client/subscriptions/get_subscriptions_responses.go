///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetSubscriptionsReader is a Reader for the GetSubscriptions structure.
type GetSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionsOK creates a GetSubscriptionsOK with default headers values
func NewGetSubscriptionsOK() *GetSubscriptionsOK {
	return &GetSubscriptionsOK{}
}

/*GetSubscriptionsOK handles this case with default header values.

Successful operation
*/
type GetSubscriptionsOK struct {
	Payload []*v1.Subscription
}

func (o *GetSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsBadRequest creates a GetSubscriptionsBadRequest with default headers values
func NewGetSubscriptionsBadRequest() *GetSubscriptionsBadRequest {
	return &GetSubscriptionsBadRequest{}
}

/*GetSubscriptionsBadRequest handles this case with default header values.

Bad Request
*/
type GetSubscriptionsBadRequest struct {
	Payload *v1.Error
}

func (o *GetSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsUnauthorized creates a GetSubscriptionsUnauthorized with default headers values
func NewGetSubscriptionsUnauthorized() *GetSubscriptionsUnauthorized {
	return &GetSubscriptionsUnauthorized{}
}

/*GetSubscriptionsUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetSubscriptionsUnauthorized struct {
	Payload *v1.Error
}

func (o *GetSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsForbidden creates a GetSubscriptionsForbidden with default headers values
func NewGetSubscriptionsForbidden() *GetSubscriptionsForbidden {
	return &GetSubscriptionsForbidden{}
}

/*GetSubscriptionsForbidden handles this case with default header values.

access to this resource is forbidden
*/
type GetSubscriptionsForbidden struct {
	Payload *v1.Error
}

func (o *GetSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getSubscriptionsForbidden  %+v", 403, o.Payload)
}

func (o *GetSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsDefault creates a GetSubscriptionsDefault with default headers values
func NewGetSubscriptionsDefault(code int) *GetSubscriptionsDefault {
	return &GetSubscriptionsDefault{
		_statusCode: code,
	}
}

/*GetSubscriptionsDefault handles this case with default header values.

Unknown error
*/
type GetSubscriptionsDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get subscriptions default response
func (o *GetSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getSubscriptions default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}