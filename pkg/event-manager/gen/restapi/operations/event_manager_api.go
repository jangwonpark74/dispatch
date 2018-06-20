///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmware/dispatch/pkg/event-manager/gen/restapi/operations/drivers"
	"github.com/vmware/dispatch/pkg/event-manager/gen/restapi/operations/events"
	"github.com/vmware/dispatch/pkg/event-manager/gen/restapi/operations/subscriptions"
)

// NewEventManagerAPI creates a new EventManager instance
func NewEventManagerAPI(spec *loads.Document) *EventManagerAPI {
	return &EventManagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		DriversAddDriverHandler: drivers.AddDriverHandlerFunc(func(params drivers.AddDriverParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversAddDriver has not yet been implemented")
		}),
		DriversAddDriverTypeHandler: drivers.AddDriverTypeHandlerFunc(func(params drivers.AddDriverTypeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversAddDriverType has not yet been implemented")
		}),
		SubscriptionsAddSubscriptionHandler: subscriptions.AddSubscriptionHandlerFunc(func(params subscriptions.AddSubscriptionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SubscriptionsAddSubscription has not yet been implemented")
		}),
		DriversDeleteDriverHandler: drivers.DeleteDriverHandlerFunc(func(params drivers.DeleteDriverParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversDeleteDriver has not yet been implemented")
		}),
		DriversDeleteDriverTypeHandler: drivers.DeleteDriverTypeHandlerFunc(func(params drivers.DeleteDriverTypeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversDeleteDriverType has not yet been implemented")
		}),
		SubscriptionsDeleteSubscriptionHandler: subscriptions.DeleteSubscriptionHandlerFunc(func(params subscriptions.DeleteSubscriptionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SubscriptionsDeleteSubscription has not yet been implemented")
		}),
		EventsEmitEventHandler: events.EmitEventHandlerFunc(func(params events.EmitEventParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsEmitEvent has not yet been implemented")
		}),
		DriversGetDriverHandler: drivers.GetDriverHandlerFunc(func(params drivers.GetDriverParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversGetDriver has not yet been implemented")
		}),
		DriversGetDriverTypeHandler: drivers.GetDriverTypeHandlerFunc(func(params drivers.GetDriverTypeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversGetDriverType has not yet been implemented")
		}),
		DriversGetDriverTypesHandler: drivers.GetDriverTypesHandlerFunc(func(params drivers.GetDriverTypesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversGetDriverTypes has not yet been implemented")
		}),
		DriversGetDriversHandler: drivers.GetDriversHandlerFunc(func(params drivers.GetDriversParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversGetDrivers has not yet been implemented")
		}),
		SubscriptionsGetSubscriptionHandler: subscriptions.GetSubscriptionHandlerFunc(func(params subscriptions.GetSubscriptionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SubscriptionsGetSubscription has not yet been implemented")
		}),
		SubscriptionsGetSubscriptionsHandler: subscriptions.GetSubscriptionsHandlerFunc(func(params subscriptions.GetSubscriptionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SubscriptionsGetSubscriptions has not yet been implemented")
		}),
		DriversUpdateDriverHandler: drivers.UpdateDriverHandlerFunc(func(params drivers.UpdateDriverParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversUpdateDriver has not yet been implemented")
		}),
		DriversUpdateDriverTypeHandler: drivers.UpdateDriverTypeHandlerFunc(func(params drivers.UpdateDriverTypeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DriversUpdateDriverType has not yet been implemented")
		}),
		SubscriptionsUpdateSubscriptionHandler: subscriptions.UpdateSubscriptionHandlerFunc(func(params subscriptions.UpdateSubscriptionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation SubscriptionsUpdateSubscription has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// Applies when the "Cookie" header is set
		CookieAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (cookie) Cookie from header param [Cookie] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*EventManagerAPI VMware Dispatch Event Manager
 */
type EventManagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// CookieAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Cookie provided in the header
	CookieAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// DriversAddDriverHandler sets the operation handler for the add driver operation
	DriversAddDriverHandler drivers.AddDriverHandler
	// DriversAddDriverTypeHandler sets the operation handler for the add driver type operation
	DriversAddDriverTypeHandler drivers.AddDriverTypeHandler
	// SubscriptionsAddSubscriptionHandler sets the operation handler for the add subscription operation
	SubscriptionsAddSubscriptionHandler subscriptions.AddSubscriptionHandler
	// DriversDeleteDriverHandler sets the operation handler for the delete driver operation
	DriversDeleteDriverHandler drivers.DeleteDriverHandler
	// DriversDeleteDriverTypeHandler sets the operation handler for the delete driver type operation
	DriversDeleteDriverTypeHandler drivers.DeleteDriverTypeHandler
	// SubscriptionsDeleteSubscriptionHandler sets the operation handler for the delete subscription operation
	SubscriptionsDeleteSubscriptionHandler subscriptions.DeleteSubscriptionHandler
	// EventsEmitEventHandler sets the operation handler for the emit event operation
	EventsEmitEventHandler events.EmitEventHandler
	// DriversGetDriverHandler sets the operation handler for the get driver operation
	DriversGetDriverHandler drivers.GetDriverHandler
	// DriversGetDriverTypeHandler sets the operation handler for the get driver type operation
	DriversGetDriverTypeHandler drivers.GetDriverTypeHandler
	// DriversGetDriverTypesHandler sets the operation handler for the get driver types operation
	DriversGetDriverTypesHandler drivers.GetDriverTypesHandler
	// DriversGetDriversHandler sets the operation handler for the get drivers operation
	DriversGetDriversHandler drivers.GetDriversHandler
	// SubscriptionsGetSubscriptionHandler sets the operation handler for the get subscription operation
	SubscriptionsGetSubscriptionHandler subscriptions.GetSubscriptionHandler
	// SubscriptionsGetSubscriptionsHandler sets the operation handler for the get subscriptions operation
	SubscriptionsGetSubscriptionsHandler subscriptions.GetSubscriptionsHandler
	// DriversUpdateDriverHandler sets the operation handler for the update driver operation
	DriversUpdateDriverHandler drivers.UpdateDriverHandler
	// DriversUpdateDriverTypeHandler sets the operation handler for the update driver type operation
	DriversUpdateDriverTypeHandler drivers.UpdateDriverTypeHandler
	// SubscriptionsUpdateSubscriptionHandler sets the operation handler for the update subscription operation
	SubscriptionsUpdateSubscriptionHandler subscriptions.UpdateSubscriptionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *EventManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *EventManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *EventManagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *EventManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *EventManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *EventManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *EventManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the EventManagerAPI
func (o *EventManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.CookieAuth == nil {
		unregistered = append(unregistered, "CookieAuth")
	}

	if o.DriversAddDriverHandler == nil {
		unregistered = append(unregistered, "drivers.AddDriverHandler")
	}

	if o.DriversAddDriverTypeHandler == nil {
		unregistered = append(unregistered, "drivers.AddDriverTypeHandler")
	}

	if o.SubscriptionsAddSubscriptionHandler == nil {
		unregistered = append(unregistered, "subscriptions.AddSubscriptionHandler")
	}

	if o.DriversDeleteDriverHandler == nil {
		unregistered = append(unregistered, "drivers.DeleteDriverHandler")
	}

	if o.DriversDeleteDriverTypeHandler == nil {
		unregistered = append(unregistered, "drivers.DeleteDriverTypeHandler")
	}

	if o.SubscriptionsDeleteSubscriptionHandler == nil {
		unregistered = append(unregistered, "subscriptions.DeleteSubscriptionHandler")
	}

	if o.EventsEmitEventHandler == nil {
		unregistered = append(unregistered, "events.EmitEventHandler")
	}

	if o.DriversGetDriverHandler == nil {
		unregistered = append(unregistered, "drivers.GetDriverHandler")
	}

	if o.DriversGetDriverTypeHandler == nil {
		unregistered = append(unregistered, "drivers.GetDriverTypeHandler")
	}

	if o.DriversGetDriverTypesHandler == nil {
		unregistered = append(unregistered, "drivers.GetDriverTypesHandler")
	}

	if o.DriversGetDriversHandler == nil {
		unregistered = append(unregistered, "drivers.GetDriversHandler")
	}

	if o.SubscriptionsGetSubscriptionHandler == nil {
		unregistered = append(unregistered, "subscriptions.GetSubscriptionHandler")
	}

	if o.SubscriptionsGetSubscriptionsHandler == nil {
		unregistered = append(unregistered, "subscriptions.GetSubscriptionsHandler")
	}

	if o.DriversUpdateDriverHandler == nil {
		unregistered = append(unregistered, "drivers.UpdateDriverHandler")
	}

	if o.DriversUpdateDriverTypeHandler == nil {
		unregistered = append(unregistered, "drivers.UpdateDriverTypeHandler")
	}

	if o.SubscriptionsUpdateSubscriptionHandler == nil {
		unregistered = append(unregistered, "subscriptions.UpdateSubscriptionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *EventManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *EventManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "bearer":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		case "cookie":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CookieAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *EventManagerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *EventManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *EventManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *EventManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the event manager API
func (o *EventManagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *EventManagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/drivers"] = drivers.NewAddDriver(o.context, o.DriversAddDriverHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/drivertypes"] = drivers.NewAddDriverType(o.context, o.DriversAddDriverTypeHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/subscriptions"] = subscriptions.NewAddSubscription(o.context, o.SubscriptionsAddSubscriptionHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/drivers/{driverName}"] = drivers.NewDeleteDriver(o.context, o.DriversDeleteDriverHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/drivertypes/{driverTypeName}"] = drivers.NewDeleteDriverType(o.context, o.DriversDeleteDriverTypeHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/subscriptions/{subscriptionName}"] = subscriptions.NewDeleteSubscription(o.context, o.SubscriptionsDeleteSubscriptionHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = events.NewEmitEvent(o.context, o.EventsEmitEventHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drivers/{driverName}"] = drivers.NewGetDriver(o.context, o.DriversGetDriverHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drivertypes/{driverTypeName}"] = drivers.NewGetDriverType(o.context, o.DriversGetDriverTypeHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drivertypes"] = drivers.NewGetDriverTypes(o.context, o.DriversGetDriverTypesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/drivers"] = drivers.NewGetDrivers(o.context, o.DriversGetDriversHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/subscriptions/{subscriptionName}"] = subscriptions.NewGetSubscription(o.context, o.SubscriptionsGetSubscriptionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/subscriptions"] = subscriptions.NewGetSubscriptions(o.context, o.SubscriptionsGetSubscriptionsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/drivers/{driverName}"] = drivers.NewUpdateDriver(o.context, o.DriversUpdateDriverHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/drivertypes/{driverTypeName}"] = drivers.NewUpdateDriverType(o.context, o.DriversUpdateDriverTypeHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/subscriptions/{subscriptionName}"] = subscriptions.NewUpdateSubscription(o.context, o.SubscriptionsUpdateSubscriptionHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *EventManagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *EventManagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *EventManagerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *EventManagerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}