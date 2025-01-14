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

package rest

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/auth/authorization/errors"
	"github.com/semi-technologies/weaviate/usecases/kinds"
	"github.com/semi-technologies/weaviate/usecases/telemetry"
)

type batchKindHandlers struct {
	manager     *kinds.BatchManager
	requestsLog *telemetry.RequestsLog
}

func (h *batchKindHandlers) addThings(params operations.BatchingThingsCreateParams,
	principal *models.Principal) middleware.Responder {
	things, err := h.manager.AddThings(params.HTTPRequest.Context(), principal,
		params.Body.Things, params.Body.Fields)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return operations.NewBatchingThingsCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return operations.NewBatchingThingsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return operations.NewBatchingThingsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	for range params.Body.Things {
		h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalAdd)
	}
	return operations.NewBatchingThingsCreateOK().
		WithPayload(h.thingsResponse(things))
}

func (h *batchKindHandlers) thingsResponse(input kinds.BatchThings) []*models.ThingsGetResponse {
	response := make([]*models.ThingsGetResponse, len(input), len(input))
	for i, thing := range input {
		var errorResponse *models.ErrorResponse
		if thing.Err != nil {
			errorResponse = errPayloadFromSingleErr(thing.Err)
		}

		thing.Thing.ID = thing.UUID
		response[i] = &models.ThingsGetResponse{
			Thing: *thing.Thing,
			Result: &models.ThingsGetResponseAO1Result{
				Errors: errorResponse,
			},
		}
	}

	return response
}

func (h *batchKindHandlers) addActions(params operations.BatchingActionsCreateParams,
	principal *models.Principal) middleware.Responder {
	actions, err := h.manager.AddActions(params.HTTPRequest.Context(), principal,
		params.Body.Actions, params.Body.Fields)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return operations.NewBatchingActionsCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return operations.NewBatchingActionsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return operations.NewBatchingActionsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	for range params.Body.Actions {
		h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalAdd)
	}
	return operations.NewBatchingActionsCreateOK().
		WithPayload(h.actionsResponse(actions))
}

func (h *batchKindHandlers) actionsResponse(input kinds.BatchActions) []*models.ActionsGetResponse {
	response := make([]*models.ActionsGetResponse, len(input), len(input))
	for i, action := range input {
		var errorResponse *models.ErrorResponse
		if action.Err != nil {
			errorResponse = errPayloadFromSingleErr(action.Err)
		}

		action.Action.ID = action.UUID
		response[i] = &models.ActionsGetResponse{
			Action: *action.Action,
			Result: &models.ActionsGetResponseAO1Result{
				Errors: errorResponse,
			},
		}
	}

	return response
}

func (h *batchKindHandlers) addReferences(params operations.BatchingReferencesCreateParams,
	principal *models.Principal) middleware.Responder {
	references, err := h.manager.AddReferences(params.HTTPRequest.Context(), principal, params.Body)
	if err != nil {
		switch err.(type) {
		case errors.Forbidden:
			return operations.NewBatchingReferencesCreateForbidden().
				WithPayload(errPayloadFromSingleErr(err))
		case kinds.ErrInvalidUserInput:
			return operations.NewBatchingReferencesCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return operations.NewBatchingReferencesCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	for range params.Body {
		h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	}
	return operations.NewBatchingReferencesCreateOK().
		WithPayload(h.referencesResponse(references))
}

func (h *batchKindHandlers) referencesResponse(input kinds.BatchReferences) []*models.BatchReferenceResponse {
	response := make([]*models.BatchReferenceResponse, len(input), len(input))
	for i, ref := range input {
		var errorResponse *models.ErrorResponse
		var reference models.BatchReference

		status := models.BatchReferenceResponseAO1ResultStatusSUCCESS
		if ref.Err != nil {
			errorResponse = errPayloadFromSingleErr(ref.Err)
			status = models.BatchReferenceResponseAO1ResultStatusFAILED
		} else {
			reference.From = strfmt.URI(ref.From.String())
			reference.To = strfmt.URI(ref.To.String())
		}

		response[i] = &models.BatchReferenceResponse{
			BatchReference: reference,
			Result: &models.BatchReferenceResponseAO1Result{
				Errors: errorResponse,
				Status: &status,
			},
		}
	}

	return response
}

func setupKindBatchHandlers(api *operations.WeaviateAPI, requestsLog *telemetry.RequestsLog, manager *kinds.BatchManager) {
	h := &batchKindHandlers{manager, requestsLog}

	api.BatchingThingsCreateHandler = operations.
		BatchingThingsCreateHandlerFunc(h.addThings)
	api.BatchingActionsCreateHandler = operations.
		BatchingActionsCreateHandlerFunc(h.addActions)
	api.BatchingReferencesCreateHandler = operations.
		BatchingReferencesCreateHandlerFunc(h.addReferences)
}

func (h *batchKindHandlers) telemetryLogAsync(requestType, identifier string) {
	go func() {
		h.requestsLog.Register(requestType, identifier)
	}()
}
