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

package esvector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v5/esapi"
)

// PutIndex idempotently creates an index
func (r *Repo) PutIndex(ctx context.Context, index string) error {

	ok, err := r.indexExists(ctx, index)
	if err != nil {
		return fmt.Errorf("create index: %v", err)
	}

	if ok {
		return nil
	}

	body := map[string]interface{}{
		"settings": map[string]interface{}{
			// "index.mapping.single_type": true,
		},
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return fmt.Errorf("create index: %v", err)
	}

	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("create index: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		return fmt.Errorf("create index: %v", err)
	}

	return nil
}

// DeleteIndex idempotently creates an index
func (r *Repo) DeleteIndex(ctx context.Context, index string) error {
	req := esapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("delete index: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		return fmt.Errorf("delete index: %v", err)
	}

	return nil
}

func (r *Repo) indexExists(ctx context.Context, index string) (bool, error) {
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}
	req.Do(ctx, r.client)

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return false, fmt.Errorf("check index exists: %v", err)
	}

	s := res.Status()
	switch s {
	case "200 OK":
		return true, nil
	case "404 Not Found":
		return false, nil
	default:
		return false, fmt.Errorf("check index exists: status: %s", s)
	}
}

// SetMappings for overall concept of weaviate index
func (r *Repo) SetMappings(ctx context.Context, index string,
	props map[string]interface{}) error {
	var buf bytes.Buffer

	// extend existing props with vectorProp
	props[keyVector.String()] = map[string]interface{}{
		"type":       "binary",
		"doc_values": true,
	}

	props[keyCreated.String()] = map[string]interface{}{
		"type": "date",
	}

	props[keyUpdated.String()] = map[string]interface{}{
		"type": "date",
	}

	props[keyID.String()] = map[string]interface{}{
		"type": "keyword",
	}

	props[keyClassName.String()] = map[string]interface{}{
		"type": "keyword",
	}

	props[keyKind.String()] = map[string]interface{}{
		"type": "keyword",
	}

	body := map[string]interface{}{
		"properties": props,
	}

	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return err
	}

	req := esapi.IndicesPutMappingRequest{
		Index: []string{index},
		Body:  &buf,
	}
	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("set mappings: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		return fmt.Errorf("set mappings: %v", err)
	}

	return nil
}
