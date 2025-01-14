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

package kinds

import (
	"context"
	"fmt"
	"math"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

type getRepo interface {
	GetThing(context.Context, strfmt.UUID, *models.Thing) error
	GetAction(context.Context, strfmt.UUID, *models.Action) error

	ListThings(ctx context.Context, limit int, thingsResponse *models.ThingsListResponse) error
	ListActions(ctx context.Context, limit int, actionsResponse *models.ActionsListResponse) error
}

// GetThing Class from the connected DB
func (m *Manager) GetThing(ctx context.Context, principal *models.Principal,
	id strfmt.UUID, meta bool) (*models.Thing, error) {
	err := m.authorizer.Authorize(principal, "get", fmt.Sprintf("things/%s", id.String()))
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, NewErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	res, err := m.getThingFromRepo(ctx, id, meta)
	if err != nil {
		return nil, err
	}

	return res.Thing(), nil
}

// GetThings Class from the connected DB
func (m *Manager) GetThings(ctx context.Context, principal *models.Principal, limit *int64) ([]*models.Thing, error) {
	err := m.authorizer.Authorize(principal, "list", "things")
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, NewErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getThingsFromRepo(ctx, limit)
}

// GetAction Class from connected DB
func (m *Manager) GetAction(ctx context.Context, principal *models.Principal, id strfmt.UUID,
	meta bool) (*models.Action, error) {
	err := m.authorizer.Authorize(principal, "get", fmt.Sprintf("actions/%s", id.String()))
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, NewErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	action, err := m.getActionFromRepo(ctx, id, meta)
	if err != nil {
		return nil, err
	}

	return action.Action(), nil
}

// GetActions Class from connected DB
func (m *Manager) GetActions(ctx context.Context, principal *models.Principal, limit *int64) ([]*models.Action, error) {
	err := m.authorizer.Authorize(principal, "list", "actions")
	if err != nil {
		return nil, err
	}

	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, NewErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getActionsFromRepo(ctx, limit)
}

func (m *Manager) getThingFromRepo(ctx context.Context, id strfmt.UUID, meta bool) (*search.Result, error) {
	res, err := m.vectorRepo.ThingByID(ctx, id, traverser.SelectProperties{}, meta)
	if err != nil {
		return nil, NewErrInternal("repo: thing by id: %v", err)
	}

	if res == nil {
		return nil, NewErrNotFound("no thing with id '%s'", id)
	}

	return res, nil
}

func (m *Manager) getThingsFromRepo(ctx context.Context, limit *int64) ([]*models.Thing, error) {
	smartLimit := m.localLimitOrGlobalLimit(limit)

	res, err := m.vectorRepo.ThingSearch(ctx, smartLimit, nil)
	if err != nil {
		return nil, NewErrInternal("list things: %v", err)
	}

	return res.Things(), nil
}

func (m *Manager) getActionFromRepo(ctx context.Context, id strfmt.UUID, meta bool) (*search.Result, error) {
	res, err := m.vectorRepo.ActionByID(ctx, id, traverser.SelectProperties{}, meta)
	if err != nil {
		return nil, NewErrInternal("repo: action by id: %v", err)
	}

	if res == nil {
		return nil, NewErrNotFound("no action with id '%s'", id)
	}

	return res, nil
}

func (m *Manager) getActionsFromRepo(ctx context.Context, limit *int64) ([]*models.Action, error) {
	smartLimit := m.localLimitOrGlobalLimit(limit)

	res, err := m.vectorRepo.ActionSearch(ctx, smartLimit, nil)
	if err != nil {
		return nil, NewErrInternal("list actions: %v", err)
	}

	return res.Actions(), nil
}

func (m *Manager) localLimitOrGlobalLimit(paramMaxResults *int64) int {
	maxResults := m.config.Config.QueryDefaults.Limit
	// Get the max results from params, if exists
	if paramMaxResults != nil {
		maxResults = *paramMaxResults
	}

	// Max results form URL, otherwise max = config.Limit.
	return int(math.Min(float64(maxResults), float64(m.config.Config.QueryDefaults.Limit)))
}
