package kraken

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// MarketData handles communication with the market data related
// methods of the Kraken API.
type MarketData service

// Time gets the server time.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getServerTime
func (m *MarketData) Time(ctx context.Context) (*ServerTime, error) {
	req, err := m.client.newPublicRequest(ctx, http.MethodGet, "Time", nil)
	if err != nil {
		return nil, err
	}

	var v ServerTime
	if err := m.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// SystemStatus gets the current system status or trading mode.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getSystemStatus
func (m *MarketData) SystemStatus(ctx context.Context) (*SystemStatus, error) {
	req, err := m.client.newPublicRequest(ctx, http.MethodGet, "SystemStatus", nil)
	if err != nil {
		return nil, err
	}

	var v SystemStatus
	if err := m.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// AssetsOpts represents the parameters to get information about the assets available for trading on Kraken.
type AssetsOpts struct {
	Assets []Asset    `url:"assets,omitempty"`
	Class  AssetClass `url:"aclass,omitempty"`
}

// IsZero returns true if the AssetsOpts is empty.
func (o AssetsOpts) IsZero() bool {
	return len(o.Assets) == 0 && o.Class == ""
}

func (o AssetsOpts) String() string {
	v, _ := query.Values(o)
	return v.Encode()
}

// Assets gets information about the assets available for trading on Kraken.
func (m *MarketData) Assets(ctx context.Context, opts AssetsOpts) (Assets, error) {
	path := "Assets"
	if !opts.IsZero() {
		path = fmt.Sprintf("%s?%s", path, opts.String())
	}

	req, err := m.client.newPublicRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var v Assets
	if err := m.client.do(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}
