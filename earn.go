package kraken

import (
	"context"
	"errors"
	"net/http"
)

// Earn handles communication with the Earn related
// methods of the Kraken API.
type Earn service

// StatusOpts represents the parameters to get the status of the last allocation or deallocation request.
type StatusOpts struct {
	StrategyID string `json:"strategy_id,omitempty"`
}

// AllocationStatus gets the status of the last allocation request.
// Docs: https://docs.kraken.com/rest/#tag/Earn/operation/getAllocateStrategyStatus
func (e *Earn) AllocationStatus(ctx context.Context, opts StatusOpts) (*StrategyOperationStatus, error) {
	if opts.StrategyID == "" {
		return nil, errors.New("strategy_id is required")
	}

	body, err := newJSONBody(opts)
	if err != nil {
		return nil, err
	}

	req, err := e.client.newPrivateRequest(ctx, http.MethodPost, "Earn/AllocationStatus", body)
	if err != nil {
		return nil, err
	}

	var v StrategyOperationStatus
	if err := e.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// DeallocationStatus gets the status of the last deallocation request.
// Docs: https://docs.kraken.com/rest/#tag/Earn/operation/getDeallocateStrategyStatus
func (e *Earn) DeallocationStatus(ctx context.Context, opts StatusOpts) (*StrategyOperationStatus, error) {
	if opts.StrategyID == "" {
		return nil, errors.New("strategy_id is required")
	}

	body, err := newJSONBody(opts)
	if err != nil {
		return nil, err
	}

	req, err := e.client.newPrivateRequest(ctx, http.MethodPost, "Earn/DeallocationStatus", body)
	if err != nil {
		return nil, err
	}

	var v StrategyOperationStatus
	if err := e.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// StrategiesOpts represents the parameters to list earn strategies.
type StrategiesOpts struct {
	Ascending bool             `json:"ascending,omitempty"`
	Asset     Asset            `json:"asset,omitempty"`
	Cursor    string           `json:"cursor,omitempty"`
	Limit     int              `json:"limit,omitempty"`
	LockType  StrategyLockType `json:"lock_type,omitempty"`
}

// Strategies lists earn strategies along with their parameters.
// Returns only strategies that are available to the user based on geographic region.
// Docs: https://docs.kraken.com/rest/#tag/Earn/operation/listStrategies
func (e *Earn) Strategies(ctx context.Context, opts StrategiesOpts) (*Strategies, error) {
	body, err := newJSONBody(opts)
	if err != nil {
		return nil, err
	}

	req, err := e.client.newPrivateRequest(ctx, http.MethodPost, "Earn/Strategies", body)
	if err != nil {
		return nil, err
	}

	var v Strategies
	if err := e.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// AllocationsOpts represents the parameters to get allocations.
type AllocationsOpts struct {
	Ascending           bool  `json:"ascending,omitempty"`
	ConvertedAsset      Asset `json:"converted_asset,omitempty"`
	HideZeroAllocations bool  `json:"hide_zero_allocations,omitempty"`
}

// Allocations list all allocations for the user.
// Docs: https://docs.kraken.com/rest/#tag/Earn/operation/listAllocations
func (e *Earn) Allocations(ctx context.Context, opts AllocationsOpts) (*Allocations, error) {
	body, err := newJSONBody(opts)
	if err != nil {
		return nil, err
	}

	req, err := e.client.newPrivateRequest(ctx, http.MethodPost, "Earn/Allocations", body)
	if err != nil {
		return nil, err
	}

	var v Allocations
	if err := e.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
