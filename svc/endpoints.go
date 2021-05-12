// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 72999ebd2f
// Version Date: Wed Mar 17 08:36:51 UTC 2021

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	pb "todo"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	httpServerOptions    map[string][]httptransport.ServerOption
	httpRequestDecoders  map[string]httptransport.DecodeRequestFunc
	httpResponseEncoders map[string]httptransport.EncodeResponseFunc
	httpHandlerFuncs     map[string]func(http.ResponseWriter, *http.Request)

	CreateTodoEndpoint endpoint.Endpoint
	GetAllEndpoint     endpoint.Endpoint
	GetTodoEndpoint    endpoint.Endpoint
	DeleteTodoEndpoint endpoint.Endpoint
	UpdateTodoEndpoint endpoint.Endpoint
}

func NewEndpoints() Endpoints {
	return Endpoints{
		httpServerOptions:    make(map[string][]httptransport.ServerOption),
		httpRequestDecoders:  make(map[string]httptransport.DecodeRequestFunc),
		httpResponseEncoders: make(map[string]httptransport.EncodeResponseFunc),
		httpHandlerFuncs:     make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// Endpoints

func (e Endpoints) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	response, err := e.CreateTodoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CreateTodoResponse), nil
}

func (e Endpoints) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	response, err := e.GetAllEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.GetAllResponse), nil
}

func (e Endpoints) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	response, err := e.GetTodoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.GetTodoResponse), nil
}

func (e Endpoints) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	response, err := e.DeleteTodoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.DeleteTodoResponse), nil
}

func (e Endpoints) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	response, err := e.UpdateTodoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UpdateTodoResponse), nil
}

// Make Endpoints

func MakeCreateTodoEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateTodoRequest)
		v, err := s.CreateTodo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetAllEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetAllRequest)
		v, err := s.GetAll(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetTodoEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetTodoRequest)
		v, err := s.GetTodo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDeleteTodoEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DeleteTodoRequest)
		v, err := s.DeleteTodo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateTodoEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UpdateTodoRequest)
		v, err := s.UpdateTodo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"CreateTodo": {},
		"GetAll":     {},
		"GetTodo":    {},
		"DeleteTodo": {},
		"UpdateTodo": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "CreateTodo" {
			e.CreateTodoEndpoint = middleware(e.CreateTodoEndpoint)
		}
		if inc == "GetAll" {
			e.GetAllEndpoint = middleware(e.GetAllEndpoint)
		}
		if inc == "GetTodo" {
			e.GetTodoEndpoint = middleware(e.GetTodoEndpoint)
		}
		if inc == "DeleteTodo" {
			e.DeleteTodoEndpoint = middleware(e.DeleteTodoEndpoint)
		}
		if inc == "UpdateTodo" {
			e.UpdateTodoEndpoint = middleware(e.UpdateTodoEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"CreateTodo": {},
		"GetAll":     {},
		"GetTodo":    {},
		"DeleteTodo": {},
		"UpdateTodo": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "CreateTodo" {
			e.CreateTodoEndpoint = middleware("CreateTodo", e.CreateTodoEndpoint)
		}
		if inc == "GetAll" {
			e.GetAllEndpoint = middleware("GetAll", e.GetAllEndpoint)
		}
		if inc == "GetTodo" {
			e.GetTodoEndpoint = middleware("GetTodo", e.GetTodoEndpoint)
		}
		if inc == "DeleteTodo" {
			e.DeleteTodoEndpoint = middleware("DeleteTodo", e.DeleteTodoEndpoint)
		}
		if inc == "UpdateTodo" {
			e.UpdateTodoEndpoint = middleware("UpdateTodo", e.UpdateTodoEndpoint)
		}
	}
}

// WrapAllWithHttpOptionExcept wraps each Endpoint entry of filed HttpServerOptions of struct Endpoints with a
// httptransport.ServerOption.
// Use this for applying a set of server options to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllWithHttpOptionExcept(serverOption, "Status", "Ping")
func (e *Endpoints) WrapAllWithHttpOptionExcept(serverOption httptransport.ServerOption, excluded ...string) {
	included := map[string]struct{}{
		"CreateTodo": {},
		"GetAll":     {},
		"GetTodo":    {},
		"DeleteTodo": {},
		"UpdateTodo": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		var options []httptransport.ServerOption
		if o, ok := e.httpServerOptions[inc]; ok {
			options = append(o, serverOption)
		} else {
			options = make([]httptransport.ServerOption, 1)
			options[0] = serverOption
		}
		e.httpServerOptions[inc] = options
	}
}

// WrapWithHttpOption wraps one Endpoint entry of filed HttpServerOptions of struct Endpoints with a
// httptransport.ServerOption.
// WrapWithHttpOption(serverOption, "Status")
func (e *Endpoints) WrapWithHttpOption(endpoint string, serverOption httptransport.ServerOption) {
	var options []httptransport.ServerOption
	if o, ok := e.httpServerOptions[endpoint]; ok {
		options = append(o, serverOption)
	} else {
		options = []httptransport.ServerOption{
			serverOption,
		}
	}
	e.httpServerOptions[endpoint] = options
}

// GetHttpServerOptions returns all httptransport.ServerOption associated with the given endpoint.
func (e Endpoints) GetHttpServerOptions(endpoint string) []httptransport.ServerOption {
	if options, ok := e.httpServerOptions[endpoint]; ok {
		return options
	}
	return make([]httptransport.ServerOption, 0)
}

// SetHttpRequestDecoder assigns a httptransport.DecodeRequestFunc to an endpoint.
func (e Endpoints) SetHttpRequestDecoder(endpoint string, decoder httptransport.DecodeRequestFunc) {
	e.httpRequestDecoders[endpoint] = decoder
}

// GetHttpRequestDecoder returns the httptransport.DecodeRequestFunc associated with the given endpoint.
func (e Endpoints) GetHttpRequestDecoder(endpoint string, fallback httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc {
	if decoder, ok := e.httpRequestDecoders[endpoint]; ok {
		return decoder
	}
	return fallback
}

// SetHttpResponseEncoder assigns a httptransport.EncodeResponseFunc to an endpoint.
func (e Endpoints) SetHttpResponseEncoder(endpoint string, encoder httptransport.EncodeResponseFunc) {
	e.httpResponseEncoders[endpoint] = encoder
}

// GetHttpResponseEncoder returns the httptransport.EncodeResponseFunc associated with the given endpoint.
func (e Endpoints) GetHttpResponseEncoder(endpoint string, fallback httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc {
	if encoder, ok := e.httpResponseEncoders[endpoint]; ok {
		return encoder
	}
	return fallback
}

// SetHttpHandlerFunc assigns a custom http HandlerFunc to an endpoint instead of using the default one.
func (e Endpoints) SetHttpHandlerFunc(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	e.httpHandlerFuncs[endpoint] = handler
}

// GetHttpHandlerFunc returns the http HandlerFunc for the given endpoint.
func (e Endpoints) GetHttpHandlerFunc(endpoint string) func(http.ResponseWriter, *http.Request) {
	if handler, ok := e.httpHandlerFuncs[endpoint]; ok {
		return handler
	}
	return nil
}

// HasHttpHandlerFunc checks if a custom http HandlerFunc is associated with the given endpoint.
func (e Endpoints) HasHttpHandlerFunc(endpoint string) bool {
	_, ok := e.httpHandlerFuncs[endpoint]
	return ok
}
