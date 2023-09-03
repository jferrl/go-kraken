package kraken

import (
	"strconv"
)

// ServerTime represents the server time.
type ServerTime struct {
	UnixTime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}

// Status represents an status within Kraken.
type Status string

const (
	//Online means Kraken is operating normally. All order types may be submitted and trades can occur.
	Online Status = "online"
	// Maintenance means exchange is offline. No new orders or cancellations may be submitted.
	Maintenance Status = "maintenance"
	// CancelOnly means resting (open) orders can be cancelled but no new orders may be submitted. No trades will occur.
	CancelOnly Status = "cancel_only"
	// PostOnly means only post-only limit orders can be submitted. Existing orders may still be cancelled. No trades will occur.
	PostOnly Status = "post_only"
	// ReduceOnly .
	ReduceOnly Status = "reduce_only"
)

// SystemStatus represents the current system status.
type SystemStatus struct {
	Status    Status `json:"status"`
	Timestamp string `json:"timestamp"` // Current timestamp (RFC3339)
}

// Balance represents the user's balance.
type Balance string

// Float64 returns the balance as a float64.
func (b Balance) Float64() float64 {
	fValue, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0
	}
	return fValue
}

// AccountBalance represents the user's account balance.
type AccountBalance map[Asset]Balance

// OrderType represents the order type.
type OrderType string

const (
	Market          OrderType = "market"
	Limit           OrderType = "limit"
	StopLoss        OrderType = "stop-loss"
	TakeProfit      OrderType = "take-profit"
	StopLossLimit   OrderType = "stop-loss-limit"
	TakeProfitLimit OrderType = "take-profit-limit"
	SettlePosition  OrderType = "settle-position"
)

// OrderDirection defines the order direction.
type OrderDirection string

const (
	Buy  OrderDirection = "buy"
	Sell OrderDirection = "sell"
)

// OrderTrigger defines the order trigger.
type OrderTrigger string

const (
	Index OrderTrigger = "index"
	Last  OrderTrigger = "last"
)

// StopType defines the stop price type.
type StopType string

const (
	CancelNewest StopType = "cancel-newest"
	CancelOldest StopType = "cancel-oldest"
	CancelBoth   StopType = "cancel-both"
)

// TimeInForce defines the time in force.
type TimeInForce string

const (
	GoodTillCancelled TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"
	GoodTillDate      TimeInForce = "GTD"
)

// OrderDescription defines an orders description.
type OrderDescription struct {
	AssetPair      string `json:"pair"`
	Close          string `json:"close"`
	Leverage       string `json:"leverage"`
	Order          string `json:"order"`
	OrderType      string `json:"ordertype"`
	PrimaryPrice   string `json:"price"`
	SecondaryPrice string `json:"price2"`
	Type           string `json:"type"`
}

// TransactionID defines a transaction ID.
type TransactionID string

// OrderCreation defines the response from the AddOrder method.
type OrderCreation struct {
	Description OrderDescription `json:"descr"`
	Transaction []TransactionID  `json:"txid"`
}

// OrderCancelation defines the response from the CancelOrder method.
type OrderCancelation struct {
	Count int `json:"count"`
}

// PairInfo defines the information about an asset pair.
type PairInfo string

const (
	Info     PairInfo = "info"
	Leverage PairInfo = "leverage"
	Fees     PairInfo = "fees"
	Margin   PairInfo = "margin"
)
