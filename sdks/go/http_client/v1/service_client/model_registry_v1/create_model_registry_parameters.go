// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package model_registry_v1

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

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateModelRegistryParams creates a new CreateModelRegistryParams object
// with the default values initialized.
func NewCreateModelRegistryParams() *CreateModelRegistryParams {
	var ()
	return &CreateModelRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateModelRegistryParamsWithTimeout creates a new CreateModelRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateModelRegistryParamsWithTimeout(timeout time.Duration) *CreateModelRegistryParams {
	var ()
	return &CreateModelRegistryParams{

		timeout: timeout,
	}
}

// NewCreateModelRegistryParamsWithContext creates a new CreateModelRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateModelRegistryParamsWithContext(ctx context.Context) *CreateModelRegistryParams {
	var ()
	return &CreateModelRegistryParams{

		Context: ctx,
	}
}

// NewCreateModelRegistryParamsWithHTTPClient creates a new CreateModelRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateModelRegistryParamsWithHTTPClient(client *http.Client) *CreateModelRegistryParams {
	var ()
	return &CreateModelRegistryParams{
		HTTPClient: client,
	}
}

/*CreateModelRegistryParams contains all the parameters to send to the API endpoint
for the create model registry operation typically these are written to a http.Request
*/
type CreateModelRegistryParams struct {

	/*Body
	  Model body

	*/
	Body *service_model.V1ModelRegistry
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create model registry params
func (o *CreateModelRegistryParams) WithTimeout(timeout time.Duration) *CreateModelRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create model registry params
func (o *CreateModelRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create model registry params
func (o *CreateModelRegistryParams) WithContext(ctx context.Context) *CreateModelRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create model registry params
func (o *CreateModelRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create model registry params
func (o *CreateModelRegistryParams) WithHTTPClient(client *http.Client) *CreateModelRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create model registry params
func (o *CreateModelRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create model registry params
func (o *CreateModelRegistryParams) WithBody(body *service_model.V1ModelRegistry) *CreateModelRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create model registry params
func (o *CreateModelRegistryParams) SetBody(body *service_model.V1ModelRegistry) {
	o.Body = body
}

// WithOwner adds the owner to the create model registry params
func (o *CreateModelRegistryParams) WithOwner(owner string) *CreateModelRegistryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create model registry params
func (o *CreateModelRegistryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateModelRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
