// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/TonyZZhang/goharbor-client/v5/apiv2/model"
)

// NewUpdateConfigurationsParams creates a new UpdateConfigurationsParams object
// with the default values initialized.
func NewUpdateConfigurationsParams() *UpdateConfigurationsParams {
	var ()
	return &UpdateConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConfigurationsParamsWithTimeout creates a new UpdateConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateConfigurationsParamsWithTimeout(timeout time.Duration) *UpdateConfigurationsParams {
	var ()
	return &UpdateConfigurationsParams{

		timeout: timeout,
	}
}

// NewUpdateConfigurationsParamsWithContext creates a new UpdateConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateConfigurationsParamsWithContext(ctx context.Context) *UpdateConfigurationsParams {
	var ()
	return &UpdateConfigurationsParams{

		Context: ctx,
	}
}

// NewUpdateConfigurationsParamsWithHTTPClient creates a new UpdateConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateConfigurationsParamsWithHTTPClient(client *http.Client) *UpdateConfigurationsParams {
	var ()
	return &UpdateConfigurationsParams{
		HTTPClient: client,
	}
}

/*UpdateConfigurationsParams contains all the parameters to send to the API endpoint
for the update configurations operation typically these are written to a http.Request
*/
type UpdateConfigurationsParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Configurations
	  The configuration map can contain a subset of the attributes of the schema, which are to be updated.

	*/
	Configurations *model.Configurations

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update configurations params
func (o *UpdateConfigurationsParams) WithTimeout(timeout time.Duration) *UpdateConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update configurations params
func (o *UpdateConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update configurations params
func (o *UpdateConfigurationsParams) WithContext(ctx context.Context) *UpdateConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update configurations params
func (o *UpdateConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update configurations params
func (o *UpdateConfigurationsParams) WithHTTPClient(client *http.Client) *UpdateConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update configurations params
func (o *UpdateConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update configurations params
func (o *UpdateConfigurationsParams) WithXRequestID(xRequestID *string) *UpdateConfigurationsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update configurations params
func (o *UpdateConfigurationsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithConfigurations adds the configurations to the update configurations params
func (o *UpdateConfigurationsParams) WithConfigurations(configurations *model.Configurations) *UpdateConfigurationsParams {
	o.SetConfigurations(configurations)
	return o
}

// SetConfigurations adds the configurations to the update configurations params
func (o *UpdateConfigurationsParams) SetConfigurations(configurations *model.Configurations) {
	o.Configurations = configurations
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Configurations != nil {
		if err := r.SetBodyParam(o.Configurations); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
