// Code generated by go-swagger; DO NOT EDIT.

package stock_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stock api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stock api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetStocksQueryStock(params *GetStocksQueryStockParams) (*GetStocksQueryStockOK, error)

	GetStocksQueryTop(params *GetStocksQueryTopParams) (*GetStocksQueryTopOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStocksQueryStock querys information about a specific stock

  Querys information about a specific stock
*/
func (a *Client) GetStocksQueryStock(params *GetStocksQueryStockParams) (*GetStocksQueryStockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStocksQueryStockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStocksQueryStock",
		Method:             "GET",
		PathPattern:        "/stocks/queryStock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStocksQueryStockReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStocksQueryStockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStocksQueryStock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStocksQueryTop querys information about the top n stocks

  Querys information about the top N stocks
*/
func (a *Client) GetStocksQueryTop(params *GetStocksQueryTopParams) (*GetStocksQueryTopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStocksQueryTopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStocksQueryTop",
		Method:             "GET",
		PathPattern:        "/stocks/queryTop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStocksQueryTopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStocksQueryTopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStocksQueryTop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
