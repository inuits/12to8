// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServicesMyUserListParams creates a new ServicesMyUserListParams object
// with the default values initialized.
func NewServicesMyUserListParams() *ServicesMyUserListParams {

	return &ServicesMyUserListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServicesMyUserListParamsWithTimeout creates a new ServicesMyUserListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServicesMyUserListParamsWithTimeout(timeout time.Duration) *ServicesMyUserListParams {

	return &ServicesMyUserListParams{

		timeout: timeout,
	}
}

// NewServicesMyUserListParamsWithContext creates a new ServicesMyUserListParams object
// with the default values initialized, and the ability to set a context for a request
func NewServicesMyUserListParamsWithContext(ctx context.Context) *ServicesMyUserListParams {

	return &ServicesMyUserListParams{

		Context: ctx,
	}
}

// NewServicesMyUserListParamsWithHTTPClient creates a new ServicesMyUserListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServicesMyUserListParamsWithHTTPClient(client *http.Client) *ServicesMyUserListParams {

	return &ServicesMyUserListParams{
		HTTPClient: client,
	}
}

/*ServicesMyUserListParams contains all the parameters to send to the API endpoint
for the services my user list operation typically these are written to a http.Request
*/
type ServicesMyUserListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the services my user list params
func (o *ServicesMyUserListParams) WithTimeout(timeout time.Duration) *ServicesMyUserListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the services my user list params
func (o *ServicesMyUserListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the services my user list params
func (o *ServicesMyUserListParams) WithContext(ctx context.Context) *ServicesMyUserListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the services my user list params
func (o *ServicesMyUserListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the services my user list params
func (o *ServicesMyUserListParams) WithHTTPClient(client *http.Client) *ServicesMyUserListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the services my user list params
func (o *ServicesMyUserListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServicesMyUserListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
