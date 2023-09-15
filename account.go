package kraken

import (
	"context"
	"net/http"
)

// Account handles communication with the account data related
// methods of the Kraken API.
type Account service

// Balance retrieves all cash balances, net of pending withdrawals.
// Docs: https://docs.kraken.com/rest/#tag/Account-Data/operation/getAccountBalance.
func (a *Account) Balance(ctx context.Context) (AccountBalance, error) {
	req, err := a.client.newPrivateRequest(ctx, http.MethodPost, "Balance", newFormURLEncodedBody(nil))
	if err != nil {
		return nil, err
	}

	var v AccountBalance
	if err := a.client.do(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}

// ExtendedBalance retrievs all extended account balances, including credits and held amounts.
// Balance available for trading is calculated as: available balance = balance + credit - credit_used - hold_trade.
// Docs: https://docs.kraken.com/rest/#tag/Account-Data/operation/getExtendedBalance.
func (a *Account) ExtendedBalance(ctx context.Context) (AccountExtendedBalance, error) {
	req, err := a.client.newPrivateRequest(ctx, http.MethodPost, "BalanceEx", newFormURLEncodedBody(nil))
	if err != nil {
		return nil, err
	}

	var v AccountExtendedBalance
	if err := a.client.do(req, &v); err != nil {
		return nil, err
	}

	return v, nil
}
