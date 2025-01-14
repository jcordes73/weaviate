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
	uuid "github.com/satori/go.uuid"
	"github.com/semi-technologies/weaviate/entities/aggregation"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Aggregations(t *testing.T) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9201"},
	})
	require.Nil(t, err)

	logger := logrus.New()
	schemaGetter := &fakeSchemaGetter{
		schema: schema.Schema{
			Things: &models.Schema{
				Classes: []*models.Class{productClass},
			},
		},
	}
	repo := NewRepo(client, logger, schemaGetter, 3)
	waitForEsToBeReady(t, repo)
	migrator := NewMigrator(repo)

	t.Run("prepare test schema and data ",
		prepareCompanyTestSchemaAndData(repo, migrator))

	t.Run("numerical aggregations with grouping",
		testNumericalAggregationsWithGrouping(repo))

	t.Run("numerical aggregations without grouping (formerly Meta)",
		testNumericalAggregationsWithoutGrouping(repo))

	t.Run("clean up",
		cleanupCompanyTestSchemaAndData(repo, migrator))

}

func prepareCompanyTestSchemaAndData(repo *Repo,
	migrator *Migrator) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("creating the class", func(t *testing.T) {
			require.Nil(t,
				migrator.AddClass(context.Background(), kind.Thing, productClass))
			require.Nil(t,
				migrator.AddClass(context.Background(), kind.Thing, companyClass))
		})

		for i, schema := range companies {
			t.Run(fmt.Sprintf("importing company %d", i), func(t *testing.T) {
				fixture := models.Thing{
					Class:  companyClass.Class,
					ID:     strfmt.UUID(uuid.Must(uuid.NewV4()).String()),
					Schema: schema,
				}
				require.Nil(t,
					repo.PutThing(context.Background(), &fixture, []float32{0, 0, 0, 0}))
			})
		}

		refreshAll(t, repo.client)
	}
}

func cleanupCompanyTestSchemaAndData(repo *Repo,
	migrator *Migrator) func(t *testing.T) {
	return func(t *testing.T) {
		migrator.DropClass(context.Background(), kind.Thing, companyClass.Class)
	}
}

func testNumericalAggregationsWithGrouping(repo *Repo) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("single field, single aggregator", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:      kind.Thing,
				ClassName: schema.ClassName(companyClass.Class),
				GroupBy: &filters.Path{
					Class:    schema.ClassName(companyClass.Class),
					Property: schema.PropertyName("sector"),
				},
				Properties: []traverser.AggregateProperty{
					traverser.AggregateProperty{
						Name:        schema.PropertyName("dividendYield"),
						Aggregators: []traverser.Aggregator{traverser.MeanAggregator},
					},
				},
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						Count: 6,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"sector"},
							Value: "Food",
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean": 2.06667,
								},
							},
						},
					},
					aggregation.Group{
						Count: 3,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"sector"},
							Value: "Financials",
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean": 2.2,
								},
							},
						},
					},
				},
			}

			assert.ElementsMatch(t, expectedResult.Groups, res.Groups)
		})

		t.Run("grouping by a non-numerical, non-string prop", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:      kind.Thing,
				ClassName: schema.ClassName(companyClass.Class),
				GroupBy: &filters.Path{
					Class:    schema.ClassName(companyClass.Class),
					Property: schema.PropertyName("listedInIndex"),
				},
				Properties: []traverser.AggregateProperty{
					traverser.AggregateProperty{
						Name:        schema.PropertyName("dividendYield"),
						Aggregators: []traverser.Aggregator{traverser.MeanAggregator},
					},
				},
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						Count: 8,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"listedInIndex"},
							Value: 1.0,
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean": 2.375,
								},
							},
						},
					},
					aggregation.Group{
						Count: 1,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"listedInIndex"},
							Value: 0.0,
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean": 0.0,
								},
							},
						},
					},
				},
			}

			assert.ElementsMatch(t, expectedResult.Groups, res.Groups)
		})

		t.Run("multiple fields, multiple aggregators, grouped by string", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:      kind.Thing,
				ClassName: schema.ClassName(companyClass.Class),
				GroupBy: &filters.Path{
					Class:    schema.ClassName(companyClass.Class),
					Property: schema.PropertyName("sector"),
				},
				Properties: []traverser.AggregateProperty{
					traverser.AggregateProperty{
						Name: schema.PropertyName("dividendYield"),
						Aggregators: []traverser.Aggregator{
							traverser.MeanAggregator,
							traverser.MaximumAggregator,
							traverser.MinimumAggregator,
							traverser.SumAggregator,
							traverser.ModeAggregator,
							traverser.MedianAggregator,
							traverser.CountAggregator,
							traverser.TypeAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("price"),
						Aggregators: []traverser.Aggregator{
							traverser.TypeAggregator,
							traverser.MeanAggregator,
							traverser.MaximumAggregator,
							traverser.MinimumAggregator,
							traverser.SumAggregator,
							traverser.ModeAggregator,
							traverser.MedianAggregator,
							traverser.CountAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("listedInIndex"),
						Aggregators: []traverser.Aggregator{
							traverser.TypeAggregator,
							traverser.PercentageTrueAggregator,
							traverser.PercentageFalseAggregator,
							traverser.TotalTrueAggregator,
							traverser.TotalFalseAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("location"),
						Aggregators: []traverser.Aggregator{
							traverser.TypeAggregator,
							traverser.TopOccurrencesAggregator,
						},
					},
				},
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						Count: 6,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"sector"},
							Value: "Food",
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    2.06667,
									"maximum": 8.0,
									"minimum": 0.0,
									"sum":     12.4,
									"mode":    0,
									"median":  1.2,
									"count":   6,
								},
							},
							"price": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    218.33333,
									"maximum": 800,
									"minimum": 10,
									"sum":     1310,
									"mode":    70,
									"median":  115,
									"count":   6,
								},
							},
							"listedInIndex": aggregation.Property{
								Type: aggregation.PropertyTypeBoolean,
								BooleanAggregation: aggregation.Boolean{
									TotalTrue:       5,
									TotalFalse:      1,
									PercentageTrue:  0.83333,
									PercentageFalse: 0.16667,
									Count:           6,
								},
							},
							"location": aggregation.Property{
								Type: aggregation.PropertyTypeText,
								TextAggregation: aggregation.Text{
									aggregation.TextOccurrence{
										Value:  "Atlanta",
										Occurs: 2,
									},
									aggregation.TextOccurrence{
										Value:  "Detroit",
										Occurs: 1,
									},
									aggregation.TextOccurrence{
										Value:  "Los Angeles",
										Occurs: 1,
									},
									aggregation.TextOccurrence{
										Value:  "New York",
										Occurs: 1,
									},
									aggregation.TextOccurrence{
										Value:  "San Francisco",
										Occurs: 1,
									},
								},
							},
						},
					},
					aggregation.Group{
						Count: 3,
						GroupedBy: &aggregation.GroupedBy{
							Path:  []string{"sector"},
							Value: "Financials",
						},
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    2.2,
									"maximum": 4.0,
									"minimum": 1.3,
									"sum":     6.6,
									"mode":    1.3,
									"median":  1.3,
									"count":   3,
								},
							},
							"price": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    265.66667,
									"maximum": 600,
									"minimum": 47,
									"sum":     797,
									"mode":    47,
									"median":  150,
									"count":   3,
								},
							},
							"listedInIndex": aggregation.Property{
								Type: aggregation.PropertyTypeBoolean,
								BooleanAggregation: aggregation.Boolean{
									TotalTrue:       3,
									TotalFalse:      0,
									PercentageTrue:  1,
									PercentageFalse: 0,
									Count:           3,
								},
							},
							"location": aggregation.Property{
								Type: aggregation.PropertyTypeText,
								TextAggregation: aggregation.Text{
									aggregation.TextOccurrence{
										Value:  "New York",
										Occurs: 2,
									},
									aggregation.TextOccurrence{
										Value:  "San Francisco",
										Occurs: 1,
									},
								},
							},
						},
					},
				},
			}

			// assert.ElementsMatch(t, expectedResult.Groups, res.Groups)
			assert.Equal(t, expectedResult.Groups, res.Groups)
		})
	}
}

func testNumericalAggregationsWithoutGrouping(repo *Repo) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("only meta count, no other aggregations", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:             kind.Thing,
				ClassName:        schema.ClassName(companyClass.Class),
				IncludeMetaCount: true,
				GroupBy:          nil, // explicitly set to nil
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						GroupedBy: nil,
						Count:     9,
					},
				},
			}

			assert.Equal(t, expectedResult.Groups, res.Groups)
		})
		t.Run("single field, single aggregator", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:      kind.Thing,
				ClassName: schema.ClassName(companyClass.Class),
				GroupBy:   nil, // explicitly set to nil
				Properties: []traverser.AggregateProperty{
					traverser.AggregateProperty{
						Name:        schema.PropertyName("dividendYield"),
						Aggregators: []traverser.Aggregator{traverser.MeanAggregator},
					},
				},
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						GroupedBy: nil,
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean": 2.11111,
								},
							},
						},
					},
				},
			}

			assert.Equal(t, expectedResult.Groups, res.Groups)
		})

		t.Run("multiple fields, multiple aggregators", func(t *testing.T) {
			params := traverser.AggregateParams{
				Kind:             kind.Thing,
				ClassName:        schema.ClassName(companyClass.Class),
				GroupBy:          nil, // explicitly set to nil,
				IncludeMetaCount: true,
				Properties: []traverser.AggregateProperty{
					traverser.AggregateProperty{
						Name: schema.PropertyName("dividendYield"),
						Aggregators: []traverser.Aggregator{
							traverser.MeanAggregator,
							traverser.MaximumAggregator,
							traverser.MinimumAggregator,
							traverser.SumAggregator,
							traverser.ModeAggregator,
							traverser.MedianAggregator,
							traverser.CountAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("price"),
						Aggregators: []traverser.Aggregator{
							traverser.MeanAggregator,
							traverser.MaximumAggregator,
							traverser.MinimumAggregator,
							traverser.SumAggregator,
							traverser.ModeAggregator,
							traverser.MedianAggregator,
							traverser.CountAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("listedInIndex"),
						Aggregators: []traverser.Aggregator{
							traverser.PercentageTrueAggregator,
							traverser.PercentageFalseAggregator,
							traverser.TotalTrueAggregator,
							traverser.TotalFalseAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("location"),
						Aggregators: []traverser.Aggregator{
							traverser.TopOccurrencesAggregator,
						},
					},
					traverser.AggregateProperty{
						Name: schema.PropertyName("makesProduct"),
						Aggregators: []traverser.Aggregator{
							traverser.PointingToAggregator,
							traverser.TypeAggregator,
						},
					},
				},
			}

			res, err := repo.Aggregate(context.Background(), params)
			require.Nil(t, err)

			expectedResult := &aggregation.Result{
				Groups: []aggregation.Group{
					aggregation.Group{
						Count: 9, // because includeMetaCount was set
						Properties: map[string]aggregation.Property{
							"dividendYield": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    2.11111,
									"maximum": 8.0,
									"minimum": 0.0,
									"sum":     19,
									"mode":    1.3,
									"median":  1.3,
									"count":   9,
								},
							},
							"price": aggregation.Property{
								Type: aggregation.PropertyTypeNumerical,
								NumericalAggregations: map[string]float64{
									"mean":    234.11111,
									"maximum": 800,
									"minimum": 10,
									"sum":     2107,
									"mode":    70,
									"median":  150,
									"count":   9,
								},
							},
							"listedInIndex": aggregation.Property{
								Type: aggregation.PropertyTypeBoolean,
								BooleanAggregation: aggregation.Boolean{
									TotalTrue:       8,
									TotalFalse:      1,
									PercentageTrue:  0.88889,
									PercentageFalse: 0.11111,
									Count:           9,
								},
							},
							"location": aggregation.Property{
								Type: aggregation.PropertyTypeText,
								TextAggregation: aggregation.Text{
									aggregation.TextOccurrence{
										Value:  "New York",
										Occurs: 3,
									},
									aggregation.TextOccurrence{
										Value:  "Atlanta",
										Occurs: 2,
									},
									aggregation.TextOccurrence{
										Value:  "San Francisco",
										Occurs: 2,
									},
									aggregation.TextOccurrence{
										Value:  "Detroit",
										Occurs: 1,
									},
									aggregation.TextOccurrence{
										Value:  "Los Angeles",
										Occurs: 1,
									},
								},
							},
						},
					},
				},
			}

			assert.Equal(t, expectedResult.Groups, res.Groups)
		})
	}
}
