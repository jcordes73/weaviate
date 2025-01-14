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
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5/esapi"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	schemaUC "github.com/semi-technologies/weaviate/usecases/schema"
	"github.com/sirupsen/logrus"
)

type internalKey string

func (k internalKey) String() string {
	return string(k)
}

const (
	keyVector    internalKey = "_embedding_vector"
	keyID        internalKey = "_uuid"
	keyKind      internalKey = "_kind"
	keyClassName internalKey = "_class_name"
	keyCreated   internalKey = "_created"
	keyUpdated   internalKey = "_updated"
	keyCache     internalKey = "_cache"
	keyCacheHot  internalKey = "_hot"

	// meta in references
	keyMeta                              internalKey = "meta"
	keyMetaClassification                internalKey = "classification"
	keyMetaClassificationWinningDistance internalKey = "winningDistance"
	keyMetaClassificationLosingDistance  internalKey = "losingDistance"

	// object meta
	keyObjectMeta internalKey = "_meta"
)

// Repo stores and retrieves vector info in elasticsearch
type Repo struct {
	client                    *elasticsearch.Client
	logger                    logrus.FieldLogger
	schemaGetter              schemaUC.SchemaGetter
	denormalizationDepthLimit int
	requestCounter            counter
	cacheIndexer              *cacheIndexer
	schemaRefFinder           schemaRefFinder
}

type schemaRefFinder interface {
	Find(className schema.ClassName) []filters.Path
}

type noopSchemaRefFinder struct{}

func (s *noopSchemaRefFinder) Find(className schema.ClassName) []filters.Path {
	return nil
}

type counter interface {
	Inc()
}

type noopCounter struct{}

func (c *noopCounter) Inc() {}

// NewRepo from existing es client
func NewRepo(client *elasticsearch.Client, logger logrus.FieldLogger,
	schemaGetter schemaUC.SchemaGetter, denormalizationLimit int) *Repo {
	return &Repo{
		client:                    client,
		logger:                    logger,
		schemaGetter:              schemaGetter,
		denormalizationDepthLimit: denormalizationLimit,
		requestCounter:            &noopCounter{},
		cacheIndexer:              nil,
		schemaRefFinder:           &noopSchemaRefFinder{},
	}
}

func (r *Repo) SetSchemaGetter(sg schemaUC.SchemaGetter) {
	r.schemaGetter = sg
}

func (r *Repo) SetSchemaRefFinder(srf schemaRefFinder) {
	r.schemaRefFinder = srf
}

func (r *Repo) WaitForStartup(maxWaitTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maxWaitTime)
	defer cancel()

	r.logger.
		WithField("action", "esvector_startup").
		WithField("maxWaitTime", maxWaitTime).
		Infof("waiting for es vector to start up (maximum %s)", maxWaitTime)

	var lastErr error

	for {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("esvector didn't start up in time: %v, last error: %v", err, lastErr)
		}

		_, err := r.client.Info()
		if err != nil {
			lastErr = err
			r.logger.WithError(err).WithField("action", "esvector_startup_cycle").
				Debug("esvector not ready yet, trying again in 1s")
		} else {
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}

// PutThing idempotently adds a Thing with its vector representation
func (r *Repo) PutThing(ctx context.Context,
	concept *models.Thing, vector []float32) error {
	err := r.putObject(ctx, kind.Thing, concept.ID.String(),
		concept.Class, concept.Schema, concept.Meta, vector, concept.CreationTimeUnix,
		concept.LastUpdateTimeUnix)
	if err != nil {
		return fmt.Errorf("put thing: %v", err)
	}

	return nil
}

// PutAction idempotently adds a Action with its vector representation
func (r *Repo) PutAction(ctx context.Context,
	concept *models.Action, vector []float32) error {
	err := r.putObject(ctx, kind.Action, concept.ID.String(),
		concept.Class, concept.Schema, concept.Meta, vector, concept.CreationTimeUnix,
		concept.LastUpdateTimeUnix)
	if err != nil {
		return fmt.Errorf("put action: %v", err)
	}

	return nil
}

func (r *Repo) objectBucket(k kind.Kind, id, className string, props models.PropertySchema,
	meta *models.ObjectMeta, vector []float32, createTime, updateTime int64) map[string]interface{} {

	bucket := map[string]interface{}{
		keyKind.String():       k.Name(),
		keyID.String():         id,
		keyClassName.String():  className,
		keyVector.String():     vectorToBase64(vector),
		keyCreated.String():    createTime,
		keyUpdated.String():    updateTime,
		keyObjectMeta.String(): meta,
	}

	ex := r.addPropsToBucket(bucket, props)
	return ex
}

func (r *Repo) putObject(ctx context.Context,
	k kind.Kind, id, className string, props models.PropertySchema,
	meta *models.ObjectMeta, vector []float32, createTime, updateTime int64) error {

	bucket := r.objectBucket(k, id, className, props, meta, vector, createTime, updateTime)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(bucket)
	if err != nil {
		return fmt.Errorf("index request: encode json: %v", err)
	}

	req := esapi.IndexRequest{
		Index:      classIndexFromClassName(k, className),
		DocumentID: id,
		Body:       &buf,
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("index request: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		r.logger.WithField("action", "vector_index_put_concept").
			WithError(err).
			WithField("request", req).
			WithField("res", res).
			WithField("body_before_marshal", bucket).
			WithField("body", buf.String()).
			Errorf("put concept failed")

		return fmt.Errorf("index request: %v", err)
	}

	go r.invalidateCache(className, id)

	return nil
}

func (r *Repo) DeleteThing(ctx context.Context, className string, id strfmt.UUID) error {
	req := esapi.DeleteRequest{
		Index:      classIndexFromClassName(kind.Thing, className),
		DocumentID: id.String(),
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("index request: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		return fmt.Errorf("delete thing: %v", err)
	}

	return nil
}

func (r *Repo) DeleteAction(ctx context.Context, className string, id strfmt.UUID) error {
	req := esapi.DeleteRequest{
		Index:      classIndexFromClassName(kind.Action, className),
		DocumentID: id.String(),
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("index request: %v", err)
	}

	if err := errorResToErr(res, r.logger); err != nil {
		return fmt.Errorf("delete action: %v", err)
	}

	return nil
}

func (r *Repo) addPropsToBucket(bucket map[string]interface{}, props models.PropertySchema) map[string]interface{} {
	if props == nil {
		return bucket
	}

	hasRefs := false

	propsMap := props.(map[string]interface{})
	for key, value := range propsMap {
		if gc, ok := value.(*models.GeoCoordinates); ok {
			value = map[string]interface{}{
				"lat": gc.Latitude,
				"lon": gc.Longitude,
			}
		}

		if _, ok := value.(models.MultipleRef); ok {
			hasRefs = true
		}

		bucket[key] = value
	}

	bucket[keyCache.String()] = map[string]interface{}{
		// if a prop has Refs, it requires caching, therefore the intial state of
		// the cache is cold. However, if there are no ref props set,no caching is
		// required making the cache state hot
		keyCacheHot.String(): !hasRefs,
	}
	return bucket
}

func vectorToBase64(array []float32) string {
	bytes := make([]byte, 0, 4*len(array))
	for _, a := range array {
		bits := math.Float32bits(a)
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, bits)
		bytes = append(bytes, b...)
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

func base64ToVector(base64Str string) ([]float32, error) {
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	length := len(decoded)
	array := make([]float32, 0, length/4)

	for i := 0; i < len(decoded); i += 4 {
		bits := binary.BigEndian.Uint32(decoded[i : i+4])
		f := math.Float32frombits(bits)
		array = append(array, f)
	}
	return array, nil
}

func errorResToErr(res *esapi.Response, logger logrus.FieldLogger) error {
	if !res.IsError() {
		return nil
	}

	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		return fmt.Errorf("request is error: status: %s", res.Status())
	}

	logger.WithField("error", e).Error("error response from es")

	shardInfo := extractShardInfoFromError(e["error"].(map[string]interface{}))
	return fmt.Errorf("request is error: status: %v, type: %v, reason: %v, shards: %v",
		res.Status(),
		e["error"].(map[string]interface{})["type"],
		e["error"].(map[string]interface{})["reason"],
		shardInfo,
	)
}

func extractShardInfoFromError(errorMap map[string]interface{}) string {
	failedShards, ok := errorMap["failed_shards"]
	if !ok {
		return ""
	}

	asSlice, ok := failedShards.([]interface{})
	if !ok {
		return ""
	}

	if len(asSlice) == 0 {
		return ""
	}

	var msgs strings.Builder

	for i, shard := range asSlice {
		asMap, ok := shard.(map[string]interface{})
		if !ok {
			continue
		}

		reason, ok := asMap["reason"]
		if !ok {
			continue
		}

		reasonMap, ok := reason.(map[string]interface{})
		if !ok {
			continue
		}

		cause, ok := reasonMap["caused_by"]
		if !ok {
			continue
		}

		causeMap, ok := cause.(map[string]interface{})
		if !ok {
			continue
		}

		innerReason, ok := causeMap["reason"]
		if !ok {
			continue
		}

		if i != 0 {
			msgs.WriteString(", ")
		}
		msgs.WriteString(innerReason.(string))
	}

	return msgs.String()
}
