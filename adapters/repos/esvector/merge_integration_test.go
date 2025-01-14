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

// +build integrationTest

package esvector

import (
	"context"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v5"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/kinds"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This test suite does not care about caching and other side offects of adding
// a ref. This is only a mechanism of mergin as outlined in
// https://github.com/semi-technologies/weaviate/issues/1021 while making sure
// that we don't reintroduce the issues from
// https://github.com/semi-technologies/weaviate/issues/1016

func Test_MergingObjects(t *testing.T) {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9201"},
	})
	require.Nil(t, err)
	schema := schema.Schema{
		Things: &models.Schema{
			Classes: []*models.Class{
				&models.Class{
					Class: "MergeTestTarget",
					Properties: []*models.Property{
						&models.Property{
							Name:     "name",
							DataType: []string{"string"},
						},
					},
				},
				&models.Class{
					Class: "MergeTestSource",
					Properties: []*models.Property{ // tries to have "one of each property type"
						&models.Property{
							Name:     "string",
							DataType: []string{"string"},
						},
						&models.Property{
							Name:     "text",
							DataType: []string{"text"},
						},
						&models.Property{
							Name:     "number",
							DataType: []string{"number"},
						},
						&models.Property{
							Name:     "int",
							DataType: []string{"int"},
						},
						&models.Property{
							Name:     "date",
							DataType: []string{"date"},
						},
						&models.Property{
							Name:     "geo",
							DataType: []string{"geoCoordinates"},
						},
						&models.Property{
							Name:     "toTarget",
							DataType: []string{"MergeTestTarget"},
						},
					},
				},
			},
		},
	}
	schemaGetter := &fakeSchemaGetter{schema: schema}
	logger := logrus.New()
	repo := NewRepo(client, logger, schemaGetter, 2)
	waitForEsToBeReady(t, repo)
	migrator := NewMigrator(repo)

	t.Run("add required classes", func(t *testing.T) {
		for _, class := range schema.Things.Classes {
			t.Run(fmt.Sprintf("add %s", class.Class), func(t *testing.T) {
				err := migrator.AddClass(context.Background(), kind.Thing, class)
				require.Nil(t, err)
			})
		}
	})

	target1 := strfmt.UUID("897be7cc-1ae1-4b40-89d9-d3ea98037638")
	target2 := strfmt.UUID("5cc94aba-93e4-408a-ab19-3d803216a04e")
	target3 := strfmt.UUID("81982705-8b1e-4228-b84c-911818d7ee85")
	target4 := strfmt.UUID("7f69c263-17f4-4529-a54d-891a7c008ca4")
	sourceID := strfmt.UUID("8738ddd5-a0ed-408d-a5d6-6f818fd56be6")

	t.Run("add objects", func(t *testing.T) {
		err := repo.PutThing(context.Background(), &models.Thing{
			ID:    sourceID,
			Class: "MergeTestSource",
			Schema: map[string]interface{}{
				"string": "only the string prop set",
			},
		}, []float32{0.5})
		require.Nil(t, err)

		targets := []strfmt.UUID{target1, target2, target3, target4}

		for i, target := range targets {
			err = repo.PutThing(context.Background(), &models.Thing{
				ID:    target,
				Class: "MergeTestTarget",
				Schema: map[string]interface{}{
					"name": fmt.Sprintf("target item %d", i),
				},
			}, []float32{0.5})
			require.Nil(t, err)
		}
	})

	refreshAll(t, client)

	t.Run("merge other previously unset properties into it", func(t *testing.T) {
		// source, err := crossref.ParseSource(fmt.Sprintf(
		// 	"weaviate://localhost/things/AddingBatchReferencesTestSource/%s/toTarget", sourceID))
		// require.Nil(t, err)
		// targets := []strfmt.UUID{target1, target2}

		md := kinds.MergeDocument{
			Class: "MergeTestSource",
			ID:    sourceID,
			Kind:  kind.Thing,
			PrimitiveSchema: map[string]interface{}{
				"number": 7.0,
				"int":    9,
				"geo": &models.GeoCoordinates{
					Latitude:  30.2,
					Longitude: 60.2,
				},
				"text": "some text",
			},
		}

		err := repo.Merge(context.Background(), md)
		assert.Nil(t, err)
	})

	refreshAll(t, client)

	t.Run("check that the object was successfully merged", func(t *testing.T) {
		source, err := repo.ThingByID(context.Background(), sourceID, nil, false)
		require.Nil(t, err)

		schema := source.Thing().Schema.(map[string]interface{})
		expectedSchema := map[string]interface{}{
			// from original
			"string": "only the string prop set",

			// from merge
			"number": 7.0,
			"int":    float64(9),
			"geo": &models.GeoCoordinates{
				Latitude:  30.2,
				Longitude: 60.2,
			},
			"text": "some text",
		}

		assert.Equal(t, expectedSchema, schema)
	})

	t.Run("add a reference and replace one prop", func(t *testing.T) {
		source, err := crossref.ParseSource(fmt.Sprintf(
			"weaviate://localhost/things/MergeTestSource/%s/toTarget", sourceID))
		require.Nil(t, err)
		targets := []strfmt.UUID{target1}
		refs := make(kinds.BatchReferences, len(targets), len(targets))
		for i, target := range targets {
			to, err := crossref.Parse(fmt.Sprintf("weaviate://localhost/things/%s", target))
			require.Nil(t, err)
			refs[i] = kinds.BatchReference{
				Err:  nil,
				From: source,
				To:   to,
			}
		}
		md := kinds.MergeDocument{
			Class: "MergeTestSource",
			ID:    sourceID,
			Kind:  kind.Thing,
			PrimitiveSchema: map[string]interface{}{
				"string": "let's update the string prop",
			},
			References: refs,
		}
		err = repo.Merge(context.Background(), md)
		assert.Nil(t, err)
	})

	refreshAll(t, client)

	t.Run("check that the object was successfully merged", func(t *testing.T) {
		source, err := repo.ThingByID(context.Background(), sourceID, nil, false)
		require.Nil(t, err)

		ref, err := crossref.Parse(fmt.Sprintf("weaviate://localhost/things/%s", target1))
		require.Nil(t, err)

		schema := source.Thing().Schema.(map[string]interface{})
		expectedSchema := map[string]interface{}{
			"string": "let's update the string prop",
			"number": 7.0,
			"int":    float64(9),
			"geo": &models.GeoCoordinates{
				Latitude:  30.2,
				Longitude: 60.2,
			},
			"text": "some text",
			"toTarget": models.MultipleRef{
				ref.SingleRef(),
			},
		}

		assert.Equal(t, expectedSchema, schema)
	})

	t.Run("add more references in rapid succession", func(t *testing.T) {
		// this test case prevents a regression on gh-1016
		source, err := crossref.ParseSource(fmt.Sprintf(
			"weaviate://localhost/things/MergeTestSource/%s/toTarget", sourceID))
		require.Nil(t, err)
		targets := []strfmt.UUID{target2, target3, target4}
		refs := make(kinds.BatchReferences, len(targets), len(targets))
		for i, target := range targets {
			to, err := crossref.Parse(fmt.Sprintf("weaviate://localhost/things/%s", target))
			require.Nil(t, err)
			refs[i] = kinds.BatchReference{
				Err:  nil,
				From: source,
				To:   to,
			}
		}
		md := kinds.MergeDocument{
			Class:      "MergeTestSource",
			ID:         sourceID,
			Kind:       kind.Thing,
			References: refs,
		}
		err = repo.Merge(context.Background(), md)
		assert.Nil(t, err)
	})

	refreshAll(t, client)

	t.Run("check all references are now present", func(t *testing.T) {
		source, err := repo.ThingByID(context.Background(), sourceID, nil, false)
		require.Nil(t, err)

		refs := source.Thing().Schema.(map[string]interface{})["toTarget"]
		refsSlice, ok := refs.(models.MultipleRef)
		require.True(t, ok, fmt.Sprintf("toTarget must be models.MultipleRef, but got %#v", refs))

		foundBeacons := []string{}
		for _, ref := range refsSlice {
			foundBeacons = append(foundBeacons, ref.Beacon.String())
		}
		expectedBeacons := []string{
			fmt.Sprintf("weaviate://localhost/things/%s", target1),
			fmt.Sprintf("weaviate://localhost/things/%s", target2),
			fmt.Sprintf("weaviate://localhost/things/%s", target3),
			fmt.Sprintf("weaviate://localhost/things/%s", target4),
		}

		assert.ElementsMatch(t, foundBeacons, expectedBeacons)
	})
}
