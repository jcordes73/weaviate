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

package test

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/semi-technologies/weaviate/client/schema"
	"github.com/semi-technologies/weaviate/client/things"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/test/acceptance/helper"
)

// this test prevents a regression on
// https://github.com/semi-technologies/weaviate/issues/981
func TestInvalidDataTypeInProperty(t *testing.T) {
	t.Parallel()
	className := "WrongPropertyClass"

	t.Run("asserting that this class does not exist yet", func(t *testing.T) {
		assert.NotContains(t, GetThingClassNames(t), className)
	})

	t.Run("trying to import empty string as data type", func(t *testing.T) {

		c := &models.Class{
			Class: className,
			Properties: []*models.Property{
				&models.Property{
					Name:     "someProperty",
					DataType: []string{""},
				},
			},
		}

		params := schema.NewSchemaThingsCreateParams().WithThingClass(c)
		resp, err := helper.Client(t).Schema.SchemaThingsCreate(params, nil)
		helper.AssertRequestFail(t, resp, err, func() {
			parsed, ok := err.(*schema.SchemaThingsCreateUnprocessableEntity)
			require.True(t, ok, "error should be unprocessable entity")
			assert.Equal(t, "property 'someProperty': invalid dataType: dataType cannot be an empty string",
				parsed.Payload.Error[0].Message)
		})

	})

}

func TestAddAndRemoveThingClass(t *testing.T) {
	t.Parallel()

	randomThingClassName := "YellowCars"

	// Ensure that this name is not in the schema yet.
	t.Log("Asserting that this class does not exist yet")
	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)

	tc := &models.Class{
		Class: randomThingClassName,
	}

	t.Log("Creating class")
	params := schema.NewSchemaThingsCreateParams().WithThingClass(tc)
	resp, err := helper.Client(t).Schema.SchemaThingsCreate(params, nil)
	helper.AssertRequestOk(t, resp, err, nil)

	t.Log("Asserting that this class is now created")
	assert.Contains(t, GetThingClassNames(t), randomThingClassName)

	// Now clean up this class.
	t.Log("Remove the class")
	delParams := schema.NewSchemaThingsDeleteParams().WithClassName(randomThingClassName)
	delResp, err := helper.Client(t).Schema.SchemaThingsDelete(delParams, nil)
	helper.AssertRequestOk(t, delResp, err, nil)

	// And verify that the class does not exist anymore.
	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)
}

// TODO: https://github.com/semi-technologies/weaviate/issues/973
// // This test prevents a regression on the fix for this bug:
// // https://github.com/semi-technologies/weaviate/issues/831
// func TestDeleteSingleProperties(t *testing.T) {
// 	t.Parallel()

// 	randomThingClassName := "RedShip"

// 	// Ensure that this name is not in the schema yet.
// 	t.Log("Asserting that this class does not exist yet")
// 	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)

// 	tc := &models.Class{
// 		Class: randomThingClassName,
// 		Properties: []*models.Property{
// 			&models.Property{
// 				DataType: []string{"string"},
// 				Name:     "name",
// 			},
// 			&models.Property{
// 				DataType: []string{"string"},
// 				Name:     "description",
// 			},
// 		},
// 	}

// 	t.Log("Creating class")
// 	params := schema.NewSchemaThingsCreateParams().WithThingClass(tc)
// 	resp, err := helper.Client(t).Schema.SchemaThingsCreate(params, nil)
// 	helper.AssertRequestOk(t, resp, err, nil)

// 	t.Log("Asserting that this class is now created")
// 	assert.Contains(t, GetThingClassNames(t), randomThingClassName)

// 	t.Log("adding an instance of this particular class that uses both properties")
// 	instanceParams := things.NewThingsCreateParams().WithBody(
// 		&models.Thing{
// 			Class: randomThingClassName,
// 			Schema: map[string]interface{}{
// 				"name":        "my name",
// 				"description": "my description",
// 			},
// 		})
// 	instanceRes, err := helper.Client(t).Things.ThingsCreate(instanceParams, nil)
// 	assert.Nil(t, err, "adding a class instance should not error")

// 	t.Log("delete a single property of the class")
// 	deleteParams := schema.NewSchemaThingsPropertiesDeleteParams().
// 		WithClassName(randomThingClassName).
// 		WithPropertyName("description")
// 	_, err = helper.Client(t).Schema.SchemaThingsPropertiesDelete(deleteParams, nil)
// 	assert.Nil(t, err, "deleting the property should not error")

// 	t.Log("retrieve the class and make sure the property is gone")
// 	thing := assertGetThingEventually(t, instanceRes.Payload.ID)
// 	expectedSchema := map[string]interface{}{
// 		"name": "my name",
// 	}
// 	assert.Equal(t, expectedSchema, thing.Schema)

// 	t.Log("verifying that we can still retrieve the thing through graphQL")
// 	result := gql.AssertGraphQL(t, helper.RootAuth, "{  Get { Things { RedShip { name } } } }")
// 	ships := result.Get("Get", "Things", "RedShip").AsSlice()
// 	expectedShip := map[string]interface{}{
// 		"name": "my name",
// 	}
// 	assert.Contains(t, ships, expectedShip)

// 	t.Log("verifying other GQL/REST queries still work")
// 	gql.AssertGraphQL(t, helper.RootAuth, "{  Meta { Things { RedShip { name { count } } } } }")
// 	gql.AssertGraphQL(t, helper.RootAuth, `{  Aggregate { Things { RedShip(groupBy: ["name"]) { name { count } } } } }`)
// 	_, err = helper.Client(t).Things.ThingsList(things.NewThingsListParams(), nil)
// 	assert.Nil(t, err, "listing things should not error")

// 	t.Log("verifying we could re-add the property with the same name")
// 	readdParams := schema.NewSchemaThingsPropertiesAddParams().
// 		WithClassName(randomThingClassName).
// 		WithBody(&models.Property{
// 			Name:     "description",
// 			DataType: []string{"string"},
// 		})

// 	_, err = helper.Client(t).Schema.SchemaThingsPropertiesAdd(readdParams, nil)
// 	assert.Nil(t, err, "adding the previously deleted property again should not error")

// 	// Now clean up this class.
// 	t.Log("Remove the class")
// 	delParams := schema.NewSchemaThingsDeleteParams().WithClassName(randomThingClassName)
// 	delResp, err := helper.Client(t).Schema.SchemaThingsDelete(delParams, nil)
// 	helper.AssertRequestOk(t, delResp, err, nil)

// 	// And verify that the class does not exist anymore.
// 	assert.NotContains(t, GetThingClassNames(t), randomThingClassName)
// }

func assertGetThingEventually(t *testing.T, uuid strfmt.UUID) *models.Thing {
	var (
		resp *things.ThingsGetOK
		err  error
	)

	checkThunk := func() interface{} {
		resp, err = helper.Client(t).Things.ThingsGet(things.NewThingsGetParams().WithID(uuid), nil)
		return err == nil
	}

	helper.AssertEventuallyEqual(t, true, checkThunk)

	var thing *models.Thing

	helper.AssertRequestOk(t, resp, err, func() {
		thing = resp.Payload
	})

	return thing
}
