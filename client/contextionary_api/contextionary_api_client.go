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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contextionary api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contextionary api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
C11yConcepts checks if a concept is part of the contextionary

Checks if a concept is part of the contextionary. Concepts should be concatenated as described here: https://github.com/semi-technologies/weaviate/blob/master/docs/en/use/ontology-schema.md#camelcase
*/
func (a *Client) C11yConcepts(params *C11yConceptsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yConceptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yConceptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.concepts",
		Method:             "GET",
		PathPattern:        "/c11y/concepts/{concept}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yConceptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*C11yConceptsOK), nil

}

/*
C11yCorpusGet checks if a word or word string is part of the contextionary

Analyzes a sentence based on the contextionary
*/
func (a *Client) C11yCorpusGet(params *C11yCorpusGetParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yCorpusGetParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.corpus.get",
		Method:             "POST",
		PathPattern:        "/c11y/corpus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yCorpusGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
C11yExtensions extends the contextionary with custom concepts

Extend the contextionary with your own custom concepts
*/
func (a *Client) C11yExtensions(params *C11yExtensionsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yExtensionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.extensions",
		Method:             "POST",
		PathPattern:        "/c11y/extensions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yExtensionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*C11yExtensionsOK), nil

}

/*
C11yWords checks if a word or word string is part of the contextionary

Checks if a word or wordString is part of the contextionary. Words should be concatenated as described here: https://github.com/semi-technologies/weaviate/blob/master/docs/en/use/ontology-schema.md#camelcase
*/
func (a *Client) C11yWords(params *C11yWordsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yWordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yWordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.words",
		Method:             "GET",
		PathPattern:        "/c11y/words/{words}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yWordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*C11yWordsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
