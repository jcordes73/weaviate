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

// ThingsReferencesUpdateReader is a Reader for the ThingsReferencesUpdate structure.
type ThingsReferencesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThingsReferencesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThingsReferencesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewThingsReferencesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewThingsReferencesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewThingsReferencesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewThingsReferencesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThingsReferencesUpdateOK creates a ThingsReferencesUpdateOK with default headers values
func NewThingsReferencesUpdateOK() *ThingsReferencesUpdateOK {
	return &ThingsReferencesUpdateOK{}
}

/*ThingsReferencesUpdateOK handles this case with default header values.

Successfully replaced all the references (success is based on the behavior of the datastore).
*/
type ThingsReferencesUpdateOK struct {
}

func (o *ThingsReferencesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /things/{id}/references/{propertyName}][%d] thingsReferencesUpdateOK ", 200)
}

func (o *ThingsReferencesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesUpdateUnauthorized creates a ThingsReferencesUpdateUnauthorized with default headers values
func NewThingsReferencesUpdateUnauthorized() *ThingsReferencesUpdateUnauthorized {
	return &ThingsReferencesUpdateUnauthorized{}
}

/*ThingsReferencesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ThingsReferencesUpdateUnauthorized struct {
}

func (o *ThingsReferencesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /things/{id}/references/{propertyName}][%d] thingsReferencesUpdateUnauthorized ", 401)
}

func (o *ThingsReferencesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewThingsReferencesUpdateForbidden creates a ThingsReferencesUpdateForbidden with default headers values
func NewThingsReferencesUpdateForbidden() *ThingsReferencesUpdateForbidden {
	return &ThingsReferencesUpdateForbidden{}
}

/*ThingsReferencesUpdateForbidden handles this case with default header values.

Forbidden
*/
type ThingsReferencesUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /things/{id}/references/{propertyName}][%d] thingsReferencesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ThingsReferencesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesUpdateUnprocessableEntity creates a ThingsReferencesUpdateUnprocessableEntity with default headers values
func NewThingsReferencesUpdateUnprocessableEntity() *ThingsReferencesUpdateUnprocessableEntity {
	return &ThingsReferencesUpdateUnprocessableEntity{}
}

/*ThingsReferencesUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ThingsReferencesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /things/{id}/references/{propertyName}][%d] thingsReferencesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ThingsReferencesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThingsReferencesUpdateInternalServerError creates a ThingsReferencesUpdateInternalServerError with default headers values
func NewThingsReferencesUpdateInternalServerError() *ThingsReferencesUpdateInternalServerError {
	return &ThingsReferencesUpdateInternalServerError{}
}

/*ThingsReferencesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ThingsReferencesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ThingsReferencesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /things/{id}/references/{propertyName}][%d] thingsReferencesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ThingsReferencesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
