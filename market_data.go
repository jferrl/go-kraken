package kraken

import (
	"context"
	"net/http"
)

// MarketData handles communication with the market data related
// methods of the Kraken API.
type MarketData service

// Time returns the server time.
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
