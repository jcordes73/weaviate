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
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new schema API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schema API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SchemaActionsCreate creates a new action class in the ontology
*/
func (a *Client) SchemaActionsCreate(params *SchemaActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.create",
		Method:             "POST",
		PathPattern:        "/schema/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsCreateOK), nil

}

/*
SchemaActionsDelete removes an action class and all data in the instances from the ontology
*/
func (a *Client) SchemaActionsDelete(params *SchemaActionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/actions/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsDeleteOK), nil

}

/*
SchemaActionsPropertiesAdd adds a property to an action class
*/
func (a *Client) SchemaActionsPropertiesAdd(params *SchemaActionsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsPropertiesAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsPropertiesAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.properties.add",
		Method:             "POST",
		PathPattern:        "/schema/actions/{className}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsPropertiesAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsPropertiesAddOK), nil

}

/*
SchemaActionsPropertiesDelete removes a property from an action class
*/
func (a *Client) SchemaActionsPropertiesDelete(params *SchemaActionsPropertiesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsPropertiesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsPropertiesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.properties.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/actions/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsPropertiesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsPropertiesDeleteOK), nil

}

/*
SchemaActionsPropertiesUpdate renames or replace the keywords of the property
*/
func (a *Client) SchemaActionsPropertiesUpdate(params *SchemaActionsPropertiesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsPropertiesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsPropertiesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.properties.update",
		Method:             "PUT",
		PathPattern:        "/schema/actions/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsPropertiesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsPropertiesUpdateOK), nil

}

/*
SchemaActionsUpdate renames or replace the keywords of the action
*/
func (a *Client) SchemaActionsUpdate(params *SchemaActionsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaActionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaActionsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.actions.update",
		Method:             "PUT",
		PathPattern:        "/schema/actions/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaActionsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaActionsUpdateOK), nil

}

/*
SchemaDump dumps the current the database schema
*/
func (a *Client) SchemaDump(params *SchemaDumpParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaDumpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.dump",
		Method:             "GET",
		PathPattern:        "/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaDumpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaDumpOK), nil

}

/*
SchemaThingsCreate creates a new thing class in the ontology
*/
func (a *Client) SchemaThingsCreate(params *SchemaThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.create",
		Method:             "POST",
		PathPattern:        "/schema/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsCreateOK), nil

}

/*
SchemaThingsDelete removes a thing class and all data in the instances from the ontology
*/
func (a *Client) SchemaThingsDelete(params *SchemaThingsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/things/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsDeleteOK), nil

}

/*
SchemaThingsPropertiesAdd adds a property to a thing class
*/
func (a *Client) SchemaThingsPropertiesAdd(params *SchemaThingsPropertiesAddParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsPropertiesAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsPropertiesAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.properties.add",
		Method:             "POST",
		PathPattern:        "/schema/things/{className}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsPropertiesAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsPropertiesAddOK), nil

}

/*
SchemaThingsPropertiesDelete removes a property from a thing class
*/
func (a *Client) SchemaThingsPropertiesDelete(params *SchemaThingsPropertiesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsPropertiesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsPropertiesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.properties.delete",
		Method:             "DELETE",
		PathPattern:        "/schema/things/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsPropertiesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsPropertiesDeleteOK), nil

}

/*
SchemaThingsPropertiesUpdate renames or replace the keywords of the property
*/
func (a *Client) SchemaThingsPropertiesUpdate(params *SchemaThingsPropertiesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsPropertiesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsPropertiesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.properties.update",
		Method:             "PUT",
		PathPattern:        "/schema/things/{className}/properties/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsPropertiesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsPropertiesUpdateOK), nil

}

/*
SchemaThingsUpdate renames or replace the keywords of the thing
*/
func (a *Client) SchemaThingsUpdate(params *SchemaThingsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SchemaThingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemaThingsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schema.things.update",
		Method:             "PUT",
		PathPattern:        "/schema/things/{className}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SchemaThingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemaThingsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
