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

// ThingsReferencesCreateReader is a Reader for the ThingsReferencesCreate structure.
type ThingsReferencesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsReferencesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThingsReferencesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewThingsReferencesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewThingsReferencesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewThingsReferencesCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewThingsReferencesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsReferencesCreateOK creates a ThingsReferencesCreateOK with default headers values
func NewThingsReferencesCreateOK() *ThingsReferencesCreateOK {
	return &ThingsReferencesCreateOK{}
}

/*ThingsReferencesCreateOK handles this case with default header values.

Successfully added the reference.
*/
type ThingsReferencesCreateOK struct {
}

func (o *ThingsReferencesCreateOK) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] thingsReferencesCreateOK ", 200)
}

func (o *ThingsReferencesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesCreateUnauthorized creates a ThingsReferencesCreateUnauthorized with default headers values
func NewThingsReferencesCreateUnauthorized() *ThingsReferencesCreateUnauthorized {
	return &ThingsReferencesCreateUnauthorized{}
}

/*ThingsReferencesCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsReferencesCreateUnauthorized struct {
}

func (o *ThingsReferencesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] thingsReferencesCreateUnauthorized ", 401)
}

func (o *ThingsReferencesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesCreateForbidden creates a ThingsReferencesCreateForbidden with default headers values
func NewThingsReferencesCreateForbidden() *ThingsReferencesCreateForbidden {
	return &ThingsReferencesCreateForbidden{}
}

/*ThingsReferencesCreateForbidden handles this case with default header values.

Forbidden
*/
type ThingsReferencesCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] thingsReferencesCreateForbidden  %+v", 403, o.Payload)
}

func (o *ThingsReferencesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesCreateUnprocessableEntity creates a ThingsReferencesCreateUnprocessableEntity with default headers values
func NewThingsReferencesCreateUnprocessableEntity() *ThingsReferencesCreateUnprocessableEntity {
	return &ThingsReferencesCreateUnprocessableEntity{}
}

/*ThingsReferencesCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ThingsReferencesCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] thingsReferencesCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ThingsReferencesCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesCreateInternalServerError creates a ThingsReferencesCreateInternalServerError with default headers values
func NewThingsReferencesCreateInternalServerError() *ThingsReferencesCreateInternalServerError {
	return &ThingsReferencesCreateInternalServerError{}
}

/*ThingsReferencesCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsReferencesCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] thingsReferencesCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsReferencesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
