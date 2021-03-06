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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ArchiveComponentHubReader is a Reader for the ArchiveComponentHub structure.
type ArchiveComponentHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveComponentHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveComponentHubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewArchiveComponentHubNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewArchiveComponentHubForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveComponentHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewArchiveComponentHubDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveComponentHubOK creates a ArchiveComponentHubOK with default headers values
func NewArchiveComponentHubOK() *ArchiveComponentHubOK {
	return &ArchiveComponentHubOK{}
}

/*ArchiveComponentHubOK handles this case with default header values.

A successful response.
*/
type ArchiveComponentHubOK struct {
}

func (o *ArchiveComponentHubOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/archive][%d] archiveComponentHubOK ", 200)
}

func (o *ArchiveComponentHubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewArchiveComponentHubNoContent creates a ArchiveComponentHubNoContent with default headers values
func NewArchiveComponentHubNoContent() *ArchiveComponentHubNoContent {
	return &ArchiveComponentHubNoContent{}
}

/*ArchiveComponentHubNoContent handles this case with default header values.

No content.
*/
type ArchiveComponentHubNoContent struct {
	Payload interface{}
}

func (o *ArchiveComponentHubNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/archive][%d] archiveComponentHubNoContent  %+v", 204, o.Payload)
}

func (o *ArchiveComponentHubNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveComponentHubNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveComponentHubForbidden creates a ArchiveComponentHubForbidden with default headers values
func NewArchiveComponentHubForbidden() *ArchiveComponentHubForbidden {
	return &ArchiveComponentHubForbidden{}
}

/*ArchiveComponentHubForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ArchiveComponentHubForbidden struct {
	Payload interface{}
}

func (o *ArchiveComponentHubForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/archive][%d] archiveComponentHubForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveComponentHubForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveComponentHubForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveComponentHubNotFound creates a ArchiveComponentHubNotFound with default headers values
func NewArchiveComponentHubNotFound() *ArchiveComponentHubNotFound {
	return &ArchiveComponentHubNotFound{}
}

/*ArchiveComponentHubNotFound handles this case with default header values.

Resource does not exist.
*/
type ArchiveComponentHubNotFound struct {
	Payload interface{}
}

func (o *ArchiveComponentHubNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/archive][%d] archiveComponentHubNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveComponentHubNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveComponentHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveComponentHubDefault creates a ArchiveComponentHubDefault with default headers values
func NewArchiveComponentHubDefault(code int) *ArchiveComponentHubDefault {
	return &ArchiveComponentHubDefault{
		_statusCode: code,
	}
}

/*ArchiveComponentHubDefault handles this case with default header values.

An unexpected error response.
*/
type ArchiveComponentHubDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the archive component hub default response
func (o *ArchiveComponentHubDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveComponentHubDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/archive][%d] ArchiveComponentHub default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveComponentHubDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ArchiveComponentHubDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
