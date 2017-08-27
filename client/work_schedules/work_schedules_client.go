// Code generated by go-swagger; DO NOT EDIT.

package work_schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new work schedules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for work schedules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WorkSchedulesList APIs endpoint that allows workschedules to be viewed or edited

API endpoint that allows workschedules to be viewed or edited.
*/
func (a *Client) WorkSchedulesList(params *WorkSchedulesListParams, authInfo runtime.ClientAuthInfoWriter) (*WorkSchedulesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkSchedulesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "work_schedules_list",
		Method:             "GET",
		PathPattern:        "/api/v1/work_schedules/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkSchedulesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WorkSchedulesListOK), nil

}

/*
WorkSchedulesRead APIs endpoint that allows workschedules to be viewed or edited

API endpoint that allows workschedules to be viewed or edited.
*/
func (a *Client) WorkSchedulesRead(params *WorkSchedulesReadParams, authInfo runtime.ClientAuthInfoWriter) (*WorkSchedulesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWorkSchedulesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "work_schedules_read",
		Method:             "GET",
		PathPattern:        "/api/v1/work_schedules/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WorkSchedulesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WorkSchedulesReadOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
