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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// BookmarkRunReader is a Reader for the BookmarkRun structure.
type BookmarkRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkRunOK creates a BookmarkRunOK with default headers values
func NewBookmarkRunOK() *BookmarkRunOK {
	return &BookmarkRunOK{}
}

/*BookmarkRunOK handles this case with default header values.

A successful response.
*/
type BookmarkRunOK struct {
}

func (o *BookmarkRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunOK ", 200)
}

func (o *BookmarkRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkRunNoContent creates a BookmarkRunNoContent with default headers values
func NewBookmarkRunNoContent() *BookmarkRunNoContent {
	return &BookmarkRunNoContent{}
}

/*BookmarkRunNoContent handles this case with default header values.

No content.
*/
type BookmarkRunNoContent struct {
	Payload interface{}
}

func (o *BookmarkRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunForbidden creates a BookmarkRunForbidden with default headers values
func NewBookmarkRunForbidden() *BookmarkRunForbidden {
	return &BookmarkRunForbidden{}
}

/*BookmarkRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type BookmarkRunForbidden struct {
	Payload interface{}
}

func (o *BookmarkRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunNotFound creates a BookmarkRunNotFound with default headers values
func NewBookmarkRunNotFound() *BookmarkRunNotFound {
	return &BookmarkRunNotFound{}
}

/*BookmarkRunNotFound handles this case with default header values.

Resource does not exist.
*/
type BookmarkRunNotFound struct {
	Payload interface{}
}

func (o *BookmarkRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] bookmarkRunNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkRunDefault creates a BookmarkRunDefault with default headers values
func NewBookmarkRunDefault(code int) *BookmarkRunDefault {
	return &BookmarkRunDefault{
		_statusCode: code,
	}
}

/*BookmarkRunDefault handles this case with default header values.

An unexpected error response.
*/
type BookmarkRunDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the bookmark run default response
func (o *BookmarkRunDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkRunDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{entity}/runs/{uuid}/bookmark][%d] BookmarkRun default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkRunDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
