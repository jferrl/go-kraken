package kraken

import (
	"context"
	"net/http"
)

// MarketData handles communication with the market data related
// methods of the Kraken API.
type MarketData service

// Time gets the server time.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getServerTime
func (md *MarketData) Time(ctx context.Context) (*ServerTime, error) {
	req, err := md.client.newPublicRequest(ctx, http.MethodGet, "Time", nil)
	if err != nil {
		return nil, err
	}

	var v ServerTime
	if err := md.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// SystemStatus gets the current system status or trading mode.
// Docs: https://docs.kraken.com/rest/#tag/Market-Data/operation/getSystemStatus
func (md *MarketData) SystemStatus(ctx context.Context) (*SystemStatus, error) {
	req, err := md.client.newPublicRequest(ctx, http.MethodGet, "SystemStatus", nil)
	if err != nil {
		return nil, err
	}

	var v SystemStatus
	if err := md.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
