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

package component_hub_v1

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

// NewPatchComponentVersionParams creates a new PatchComponentVersionParams object
// with the default values initialized.
func NewPatchComponentVersionParams() *PatchComponentVersionParams {
	var ()
	return &PatchComponentVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchComponentVersionParamsWithTimeout creates a new PatchComponentVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchComponentVersionParamsWithTimeout(timeout time.Duration) *PatchComponentVersionParams {
	var ()
	return &PatchComponentVersionParams{

		timeout: timeout,
	}
}

// NewPatchComponentVersionParamsWithContext creates a new PatchComponentVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchComponentVersionParamsWithContext(ctx context.Context) *PatchComponentVersionParams {
	var ()
	return &PatchComponentVersionParams{

		Context: ctx,
	}
}

// NewPatchComponentVersionParamsWithHTTPClient creates a new PatchComponentVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchComponentVersionParamsWithHTTPClient(client *http.Client) *PatchComponentVersionParams {
	var ()
	return &PatchComponentVersionParams{
		HTTPClient: client,
	}
}

/*PatchComponentVersionParams contains all the parameters to send to the API endpoint
for the patch component version operation typically these are written to a http.Request
*/
type PatchComponentVersionParams struct {

	/*Body
	  Component version body

	*/
	Body *service_model.V1ComponentVersion
	/*Component
	  Component name

	*/
	Component string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*VersionName
	  Optional component name, should be a valid fully qualified value: name[:version]

	*/
	VersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch component version params
func (o *PatchComponentVersionParams) WithTimeout(timeout time.Duration) *PatchComponentVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch component version params
func (o *PatchComponentVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch component version params
func (o *PatchComponentVersionParams) WithContext(ctx context.Context) *PatchComponentVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch component version params
func (o *PatchComponentVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch component version params
func (o *PatchComponentVersionParams) WithHTTPClient(client *http.Client) *PatchComponentVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch component version params
func (o *PatchComponentVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch component version params
func (o *PatchComponentVersionParams) WithBody(body *service_model.V1ComponentVersion) *PatchComponentVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch component version params
func (o *PatchComponentVersionParams) SetBody(body *service_model.V1ComponentVersion) {
	o.Body = body
}

// WithComponent adds the component to the patch component version params
func (o *PatchComponentVersionParams) WithComponent(component string) *PatchComponentVersionParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the patch component version params
func (o *PatchComponentVersionParams) SetComponent(component string) {
	o.Component = component
}

// WithOwner adds the owner to the patch component version params
func (o *PatchComponentVersionParams) WithOwner(owner string) *PatchComponentVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch component version params
func (o *PatchComponentVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithVersionName adds the versionName to the patch component version params
func (o *PatchComponentVersionParams) WithVersionName(versionName string) *PatchComponentVersionParams {
	o.SetVersionName(versionName)
	return o
}

// SetVersionName adds the versionName to the patch component version params
func (o *PatchComponentVersionParams) SetVersionName(versionName string) {
	o.VersionName = versionName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchComponentVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component
	if err := r.SetPathParam("component", o.Component); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param version.name
	if err := r.SetPathParam("version.name", o.VersionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
