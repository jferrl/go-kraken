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

// ExtendedBalance represents the user's extended balance.
type ExtendedBalance struct {
	Balance    Balance `json:"balance"`
	Credit     string  `json:"credit"`
	CreditUsed string  `json:"credit_used"`
	HoldTrade  string  `json:"hold_trade"`
}

type (
	// AccountBalance represents the user's account balance.
	AccountBalance         map[Asset]Balance
	AccountExtendedBalance map[Asset]ExtendedBalance
)

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

// WebsocketsToken defines the response from the GetWebSocketsToken method.
type WebsocketsToken struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

// Tick represents a tick.
type Tick []any

// TickValues represents the values of a tick.
type TickValues struct {
	Open   string
	High   string
	Low    string
	Close  string
	Vwap   string
	Volume string
	Count  int64
	Time   int64
}

// Valid returns true if the tick is valid.
func (t Tick) Valid() bool {
	return len(t) == 8
}

// Time returns the time of the tick.
func (t Tick) Time() int64 {
	v := t[0]
	switch value := v.(type) {
	case float64:
		return int64(value)
	case int64:
		return value
	case int:
		return int64(value)
	default:
		return 0
	}
}

// Open returns the open price of the tick.
func (t Tick) Open() string {
	if v, ok := t[1].(string); ok {
		return v
	}
	return ""
}

// High returns the high price of the tick.
func (t Tick) High() string {
	if v, ok := t[2].(string); ok {
		return v
	}
	return ""
}

// Low returns the low price of the tick.
func (t Tick) Low() string {
	if v, ok := t[3].(string); ok {
		return v
	}
	return ""
}

// Close returns the close price of the tick.
func (t Tick) Close() string {
	if v, ok := t[4].(string); ok {
		return v
	}
	return ""
}

// Vwap returns the vwap of the tick.
func (t Tick) Vwap() string {
	if v, ok := t[5].(string); ok {
		return v
	}
	return ""
}

// Volume returns the volume of the tick.
func (t Tick) Volume() string {
	if v, ok := t[6].(string); ok {
		return v
	}
	return ""
}

// Count returns the count of the tick.
func (t Tick) Count() int64 {
	v := t[7]
	switch value := v.(type) {
	case float64:
		return int64(value)
	case int64:
		return value
	case int:
		return int64(value)
	default:
		return 0
	}
}

// Values returns the values of the tick.
func (t Tick) Values() TickValues {
	if !t.Valid() {
		return TickValues{}
	}

	return TickValues{
		Open:   t.Open(),
		High:   t.High(),
		Low:    t.Low(),
		Close:  t.Close(),
		Vwap:   t.Vwap(),
		Volume: t.Volume(),
		Count:  t.Count(),
		Time:   t.Time(),
	}
}

// Ticks is a slice of Tick.
type Ticks []Tick

// OHCL represents the OHCL data. It represents the "Open-high-low-close chart".
type OHCL struct {
	Last int64
	Pair Ticks
}
