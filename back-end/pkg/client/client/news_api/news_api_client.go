// Code generated by go-swagger; DO NOT EDIT.

package news_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new news api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for news api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PostNewsQuery(params *PostNewsQueryParams) (*PostNewsQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostNewsQuery retrieves a news query

  Retrieves a news query
*/
func (a *Client) PostNewsQuery(params *PostNewsQueryParams) (*PostNewsQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNewsQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNewsQuery",
		Method:             "POST",
		PathPattern:        "/news/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostNewsQueryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNewsQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostNewsQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
