package kraken

import (
	"context"
	"net/http"
)

// Earn handles communication with the Earn related
// methods of the Kraken API.
type Earn service

// StrategiesOpts
type StrategiesOpts struct {
	Ascending bool             `json:"ascending,omitempty"`
	Asset     Asset            `json:"asset,omitempty"`
	Cursor    string           `json:"cursor,omitempty"`
	Limit     int              `json:"limit,omitempty"`
	LockType  StrategyLockType `json:"lock_type,omitempty"`
}

// Strategies lists earn strategies along with their parameters.
// Returns only strategies that are available to the user based on geographic region.
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
