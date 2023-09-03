package kraken

import (
	"context"
	"net/http"
	"net/url"
)

// Account handles communication with the account data related
// methods of the Kraken API.
type Account service

// Balance retrieves all cash balances, net of pending withdrawals.
func (a *Account) Balance(ctx context.Context) (AccountBalance, error) {
	req, err := a.client.newPrivateRequest(ctx, http.MethodPost, "Balance", url.Values{})
	if err != nil {
		return nil, err
	}

	var v AccountBalance
	if err := a.client.do(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}
