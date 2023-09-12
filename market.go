package kraken

import (
	"context"
	"errors"
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

// String returns the query string representation of the AssetsOpts.
func (o AssetsOpts) String() string {
	v, _ := query.Values(o)
	return v.Encode()
}

// Assets gets information about the assets available for trading on Kraken.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getAssetInfo
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

// TradableAssetPairsOpts represents the parameters to get information about the asset pairs available for trading on Kraken.
type TradableAssetPairsOpts struct {
	Pairs []AssetPair `url:"pair,omitempty"`
	Info  PairInfo    `url:"info,omitempty"`
}

// IsZero returns true if the TradableAssetPairsOpts is empty.
func (o TradableAssetPairsOpts) IsZero() bool {
	return len(o.Pairs) == 0 && o.Info == ""
}

// String returns the query string representation of the TradableAssetPairsOpts.
func (o TradableAssetPairsOpts) String() string {
	v, _ := query.Values(o)
	return v.Encode()
}

// TradableAssetPairs gets information about the asset pairs available for trading on Kraken.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getTradableAssetPairs
func (m *MarketData) TradableAssetPairs(ctx context.Context, opts TradableAssetPairsOpts) (AssetPairs, error) {
	path := "AssetPairs"
	if !opts.IsZero() {
		path = fmt.Sprintf("%s?%s", path, opts.String())
	}

	req, err := m.client.newPublicRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var v AssetPairs
	if err := m.client.do(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}

type OHCLDataOpts struct {
	Pair     AssetPair `url:"pair,omitempty"`
	Interval int       `url:"interval,omitempty"`
	Since    int       `url:"since,omitempty"`
}

func (o OHCLDataOpts) String() string {
	v, _ := query.Values(o)
	return v.Encode()
}

type ohclData map[string]any

// OHCLData retrieves the last entry in the OHLC array is for the current,
// not-yet-committed frame and will always be present, regardless of the value of since.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getOHLCData
func (m *MarketData) OHCLData(ctx context.Context, opts OHCLDataOpts) (*OHCL, error) {
	if opts.Pair == "" {
		return nil, errors.New("pair is required")
	}

	path := fmt.Sprintf("OHLC?%s", opts.String())
	req, err := m.client.newPublicRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var v ohclData
	if err := m.client.do(req, &v); err != nil {
		return nil, err
	}

	var last int64
	if l, ok := v["last"].(float64); ok {
		last = int64(l)
	}

	var ticks Ticks
	if p, ok := v[string(opts.Pair)].([]any); ok {
		for _, t := range p {
			if pv, ok := t.([]any); ok {
				ticks = append(ticks, pv)
			}
		}
	}

	return &OHCL{
		Last: last,
		Pair: ticks,
	}, nil
}
