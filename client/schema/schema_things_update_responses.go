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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaThingsUpdateReader is a Reader for the SchemaThingsUpdate structure.
type SchemaThingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaThingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemaThingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSchemaThingsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSchemaThingsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewSchemaThingsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSchemaThingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaThingsUpdateOK creates a SchemaThingsUpdateOK with default headers values
func NewSchemaThingsUpdateOK() *SchemaThingsUpdateOK {
	return &SchemaThingsUpdateOK{}
}

/*SchemaThingsUpdateOK handles this case with default header values.

Changes applied.
*/
type SchemaThingsUpdateOK struct {
}

func (o *SchemaThingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] schemaThingsUpdateOK ", 200)
}

func (o *SchemaThingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaThingsUpdateUnauthorized creates a SchemaThingsUpdateUnauthorized with default headers values
func NewSchemaThingsUpdateUnauthorized() *SchemaThingsUpdateUnauthorized {
	return &SchemaThingsUpdateUnauthorized{}
}

/*SchemaThingsUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaThingsUpdateUnauthorized struct {
}

func (o *SchemaThingsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] schemaThingsUpdateUnauthorized ", 401)
}

func (o *SchemaThingsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaThingsUpdateForbidden creates a SchemaThingsUpdateForbidden with default headers values
func NewSchemaThingsUpdateForbidden() *SchemaThingsUpdateForbidden {
	return &SchemaThingsUpdateForbidden{}
}

/*SchemaThingsUpdateForbidden handles this case with default header values.

Forbidden
*/
type SchemaThingsUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] schemaThingsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SchemaThingsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsUpdateUnprocessableEntity creates a SchemaThingsUpdateUnprocessableEntity with default headers values
func NewSchemaThingsUpdateUnprocessableEntity() *SchemaThingsUpdateUnprocessableEntity {
	return &SchemaThingsUpdateUnprocessableEntity{}
}

/*SchemaThingsUpdateUnprocessableEntity handles this case with default header values.

Invalid update.
*/
type SchemaThingsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] schemaThingsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaThingsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsUpdateInternalServerError creates a SchemaThingsUpdateInternalServerError with default headers values
func NewSchemaThingsUpdateInternalServerError() *SchemaThingsUpdateInternalServerError {
	return &SchemaThingsUpdateInternalServerError{}
}

/*SchemaThingsUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaThingsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/things/{className}][%d] schemaThingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaThingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
