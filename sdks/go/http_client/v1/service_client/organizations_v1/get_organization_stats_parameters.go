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

package organizations_v1

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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationStatsParams creates a new GetOrganizationStatsParams object
// with the default values initialized.
func NewGetOrganizationStatsParams() *GetOrganizationStatsParams {
	var ()
	return &GetOrganizationStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationStatsParamsWithTimeout creates a new GetOrganizationStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationStatsParamsWithTimeout(timeout time.Duration) *GetOrganizationStatsParams {
	var ()
	return &GetOrganizationStatsParams{

		timeout: timeout,
	}
}

// NewGetOrganizationStatsParamsWithContext creates a new GetOrganizationStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationStatsParamsWithContext(ctx context.Context) *GetOrganizationStatsParams {
	var ()
	return &GetOrganizationStatsParams{

		Context: ctx,
	}
}

// NewGetOrganizationStatsParamsWithHTTPClient creates a new GetOrganizationStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationStatsParamsWithHTTPClient(client *http.Client) *GetOrganizationStatsParams {
	var ()
	return &GetOrganizationStatsParams{
		HTTPClient: client,
	}
}

/*GetOrganizationStatsParams contains all the parameters to send to the API endpoint
for the get organization stats operation typically these are written to a http.Request
*/
type GetOrganizationStatsParams struct {

	/*Aggregate
	  Stats aggregate.

	*/
	Aggregate *string
	/*Groupby
	  Stats group.

	*/
	Groupby *string
	/*Kind
	  Stats Kind.

	*/
	Kind *string
	/*Limit
	  Limit size.

	*/
	Limit *int32
	/*Offset
	  Pagination offset.

	*/
	Offset *int32
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Query
	  Query filter the search.

	*/
	Query *string
	/*Sort
	  Sort to order the search.

	*/
	Sort *string
	/*Trunc
	  Stats trunc.

	*/
	Trunc *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organization stats params
func (o *GetOrganizationStatsParams) WithTimeout(timeout time.Duration) *GetOrganizationStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization stats params
func (o *GetOrganizationStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization stats params
func (o *GetOrganizationStatsParams) WithContext(ctx context.Context) *GetOrganizationStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization stats params
func (o *GetOrganizationStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization stats params
func (o *GetOrganizationStatsParams) WithHTTPClient(client *http.Client) *GetOrganizationStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization stats params
func (o *GetOrganizationStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregate adds the aggregate to the get organization stats params
func (o *GetOrganizationStatsParams) WithAggregate(aggregate *string) *GetOrganizationStatsParams {
	o.SetAggregate(aggregate)
	return o
}

// SetAggregate adds the aggregate to the get organization stats params
func (o *GetOrganizationStatsParams) SetAggregate(aggregate *string) {
	o.Aggregate = aggregate
}

// WithGroupby adds the groupby to the get organization stats params
func (o *GetOrganizationStatsParams) WithGroupby(groupby *string) *GetOrganizationStatsParams {
	o.SetGroupby(groupby)
	return o
}

// SetGroupby adds the groupby to the get organization stats params
func (o *GetOrganizationStatsParams) SetGroupby(groupby *string) {
	o.Groupby = groupby
}

// WithKind adds the kind to the get organization stats params
func (o *GetOrganizationStatsParams) WithKind(kind *string) *GetOrganizationStatsParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get organization stats params
func (o *GetOrganizationStatsParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithLimit adds the limit to the get organization stats params
func (o *GetOrganizationStatsParams) WithLimit(limit *int32) *GetOrganizationStatsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get organization stats params
func (o *GetOrganizationStatsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get organization stats params
func (o *GetOrganizationStatsParams) WithOffset(offset *int32) *GetOrganizationStatsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get organization stats params
func (o *GetOrganizationStatsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the get organization stats params
func (o *GetOrganizationStatsParams) WithOwner(owner string) *GetOrganizationStatsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get organization stats params
func (o *GetOrganizationStatsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the get organization stats params
func (o *GetOrganizationStatsParams) WithQuery(query *string) *GetOrganizationStatsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get organization stats params
func (o *GetOrganizationStatsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get organization stats params
func (o *GetOrganizationStatsParams) WithSort(sort *string) *GetOrganizationStatsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get organization stats params
func (o *GetOrganizationStatsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTrunc adds the trunc to the get organization stats params
func (o *GetOrganizationStatsParams) WithTrunc(trunc *string) *GetOrganizationStatsParams {
	o.SetTrunc(trunc)
	return o
}

// SetTrunc adds the trunc to the get organization stats params
func (o *GetOrganizationStatsParams) SetTrunc(trunc *string) {
	o.Trunc = trunc
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Aggregate != nil {

		// query param aggregate
		var qrAggregate string
		if o.Aggregate != nil {
			qrAggregate = *o.Aggregate
		}
		qAggregate := qrAggregate
		if qAggregate != "" {
			if err := r.SetQueryParam("aggregate", qAggregate); err != nil {
				return err
			}
		}

	}

	if o.Groupby != nil {

		// query param groupby
		var qrGroupby string
		if o.Groupby != nil {
			qrGroupby = *o.Groupby
		}
		qGroupby := qrGroupby
		if qGroupby != "" {
			if err := r.SetQueryParam("groupby", qGroupby); err != nil {
				return err
			}
		}

	}

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if o.Trunc != nil {

		// query param trunc
		var qrTrunc string
		if o.Trunc != nil {
			qrTrunc = *o.Trunc
		}
		qTrunc := qrTrunc
		if qTrunc != "" {
			if err := r.SetQueryParam("trunc", qTrunc); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}