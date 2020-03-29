// Code generated by go-swagger; DO NOT EDIT.

package stock_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/stocksmith/stockadvisor-web-app/back-end/pkg/server/models"
)

// GetStocksQueryTopOKCode is the HTTP code returned for type GetStocksQueryTopOK
const GetStocksQueryTopOKCode int = 200

/*GetStocksQueryTopOK Successful Stock Query

swagger:response getStocksQueryTopOK
*/
type GetStocksQueryTopOK struct {

	/*
	  In: Body
	*/
	Payload []*models.NewsResponse `json:"body,omitempty"`
}

// NewGetStocksQueryTopOK creates GetStocksQueryTopOK with default headers values
func NewGetStocksQueryTopOK() *GetStocksQueryTopOK {

	return &GetStocksQueryTopOK{}
}

// WithPayload adds the payload to the get stocks query top o k response
func (o *GetStocksQueryTopOK) WithPayload(payload []*models.NewsResponse) *GetStocksQueryTopOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stocks query top o k response
func (o *GetStocksQueryTopOK) SetPayload(payload []*models.NewsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStocksQueryTopOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.NewsResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
