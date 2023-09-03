package kraken

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Trading handles communication with the trading related
// methods of the Kraken API.
type Trading service

// AddOrderOpts represents the parameters to create an Order.
type AddOrderOpts struct {
	UserRef        string         `url:"userref,omitempty"`
	OrderType      OrderType      `url:"ordertype,omitempty"`
	Type           OrderDirection `url:"type,omitempty"`
	Volume         string         `url:"volume,omitempty"`
	DisplayVol     string         `url:"displayvol,omitempty"`
	Pair           string         `url:"pair,omitempty"`
	Price          string         `url:"price,omitempty"`
	Price2         string         `url:"price2,omitempty"`
	Trigger        OrderTrigger   `url:"trigger,omitempty"`
	Leverage       string         `url:"leverage,omitempty"`
	ReduceOnly     bool           `url:"reduce_only,omitempty"`
	StopType       StopType       `url:"stptype,omitempty"`
	OrderFlags     string         `url:"oflags,omitempty"`
	TimeInForce    TimeInForce    `url:"timeinforce,omitempty"`
	Starttm        string         `url:"starttm,omitempty"`
	Expiretm       string         `url:"expiretm,omitempty"`
	CloseOrderType OrderType      `url:"close[ordertype],omitempty"`
	ClosePrice     string         `url:"close[price],omitempty"`
	ClosePrice2    string         `url:"close[price2],omitempty"`
	Deadline       string         `url:"deadline,omitempty"`
	Validate       bool           `url:"validate,omitempty"`
}

// AddOrder places a new order.
// Docs: https://docs.kraken.com/rest/#tag/Trading/operation/addOrder
func (t *Trading) AddOrder(ctx context.Context, opts AddOrderOpts) (*OrderCreation, error) {
	body, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	req, err := t.client.newPrivateRequest(ctx, http.MethodPost, "AddOrder", body)
	if err != nil {
		return nil, err
	}

	var v OrderCreation
	if err := t.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// CancelOrderOpts represents the parameters to cancel an Order.
type CancelOrderOpts struct {
	TransactionID string `url:"txid,omitempty"`
}

// CancelOrder cancels an order.
// Docs: https://docs.kraken.com/rest/#tag/Trading/operation/cancelOrder
func (t *Trading) CancelOrder(ctx context.Context, opts CancelOrderOpts) (*OrderCancelation, error) {
	body, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	req, err := t.client.newPrivateRequest(ctx, http.MethodPost, "CancelOrder", body)
	if err != nil {
		return nil, err
	}

	var v OrderCancelation
	if err := t.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
