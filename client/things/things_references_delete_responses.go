//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ThingsReferencesDeleteReader is a Reader for the ThingsReferencesDelete structure.
type ThingsReferencesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsReferencesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewThingsReferencesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewThingsReferencesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewThingsReferencesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewThingsReferencesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewThingsReferencesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsReferencesDeleteNoContent creates a ThingsReferencesDeleteNoContent with default headers values
func NewThingsReferencesDeleteNoContent() *ThingsReferencesDeleteNoContent {
	return &ThingsReferencesDeleteNoContent{}
}

/*ThingsReferencesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type ThingsReferencesDeleteNoContent struct {
}

func (o *ThingsReferencesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] thingsReferencesDeleteNoContent ", 204)
}

func (o *ThingsReferencesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesDeleteUnauthorized creates a ThingsReferencesDeleteUnauthorized with default headers values
func NewThingsReferencesDeleteUnauthorized() *ThingsReferencesDeleteUnauthorized {
	return &ThingsReferencesDeleteUnauthorized{}
}

/*ThingsReferencesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsReferencesDeleteUnauthorized struct {
}

func (o *ThingsReferencesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] thingsReferencesDeleteUnauthorized ", 401)
}

func (o *ThingsReferencesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesDeleteForbidden creates a ThingsReferencesDeleteForbidden with default headers values
func NewThingsReferencesDeleteForbidden() *ThingsReferencesDeleteForbidden {
	return &ThingsReferencesDeleteForbidden{}
}

/*ThingsReferencesDeleteForbidden handles this case with default header values.

Forbidden
*/
type ThingsReferencesDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] thingsReferencesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ThingsReferencesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesDeleteNotFound creates a ThingsReferencesDeleteNotFound with default headers values
func NewThingsReferencesDeleteNotFound() *ThingsReferencesDeleteNotFound {
	return &ThingsReferencesDeleteNotFound{}
}

/*ThingsReferencesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ThingsReferencesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] thingsReferencesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ThingsReferencesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesDeleteInternalServerError creates a ThingsReferencesDeleteInternalServerError with default headers values
func NewThingsReferencesDeleteInternalServerError() *ThingsReferencesDeleteInternalServerError {
	return &ThingsReferencesDeleteInternalServerError{}
}

/*ThingsReferencesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsReferencesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] thingsReferencesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsReferencesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
