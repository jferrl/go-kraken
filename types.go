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

// TriggeredOrderCancellation defines the response from the CancelAllOrdersAfter method.
type TriggeredOrderCancellation struct {
	// Timestamp (RFC3339 format) at which the request was received.
	CurrentTime string `json:"currentTime"`
	// Timestamp (RFC3339 format) after which all orders will be cancelled, unless the timer is extended or disabled.
	TriggerTime string `json:"triggerTime"`
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

// TickerValues represents the values of a ticker.
type TickerValues []any

// Ticker represents a ticker.
type Ticker struct {
	Open   string
	High   string
	Low    string
	Close  string
	Vwap   string
	Volume string
	Count  int64
	Time   int64
}

// time returns the time of the tick.
func (t TickerValues) time() int64 {
	v := t[0]
	switch value := v.(type) {
	case float64:
		return int64(value)
	case int:
		return int64(value)
	default:
		return 0
	}
}

// open returns the open price of the tick.
func (t TickerValues) open() string {
	if v, ok := t[1].(string); ok {
		return v
	}
	return ""
}

// high returns the high price of the tick.
func (t TickerValues) high() string {
	if v, ok := t[2].(string); ok {
		return v
	}
	return ""
}

// low returns the low price of the tick.
func (t TickerValues) low() string {
	if v, ok := t[3].(string); ok {
		return v
	}
	return ""
}

// close returns the close price of the tick.
func (t TickerValues) close() string {
	if v, ok := t[4].(string); ok {
		return v
	}
	return ""
}

// vwap returns the vwap of the tick.
func (t TickerValues) vwap() string {
	if v, ok := t[5].(string); ok {
		return v
	}
	return ""
}

// volume returns the volume of the tick.
func (t TickerValues) volume() string {
	if v, ok := t[6].(string); ok {
		return v
	}
	return ""
}

// count returns the count of the tick.
func (t TickerValues) count() int64 {
	v := t[7]
	switch value := v.(type) {
	case float64:
		return int64(value)
	case int:
		return int64(value)
	default:
		return 0
	}
}

// Ticker returns the values of the tick.
func (t TickerValues) Ticker() Ticker {
	if len(t) != 8 {
		return Ticker{}
	}

	return Ticker{
		Open:   t.open(),
		High:   t.high(),
		Low:    t.low(),
		Close:  t.close(),
		Vwap:   t.vwap(),
		Volume: t.volume(),
		Count:  t.count(),
		Time:   t.time(),
	}
}

// OHCLTickers is a slice of Tick.
type OHCLTickers []TickerValues

// OHCL represents the OHCL data. It represents the "Open-high-low-close chart".
type OHCL struct {
	Last int64
	Pair OHCLTickers
}

// AssetTickerInfo defines the information about an asset ticker.
type AssetTickerInfo struct {
	Ask                        []string `json:"a"` // Ask price array(<price>, <whole lot volume>, <lot volume>)
	Bid                        []string `json:"b"` // Bid price array(<price>, <whole lot volume>, <lot volume>)
	Last                       []string `json:"c"` // Last trade closed array(<price>, <lot volume>)
	Volume                     []string `json:"v"` // Volume array(<today>, <last 24 hours>)
	VolumeWeightedAveragePrice []string `json:"p"` // Volume weighted average price array(<today>, <last 24 hours>)
	NumberOfTrades             []int    `json:"t"` // Number of trades array(<today>, <last 24 hours>)
	Low                        []string `json:"l"` // Low array(<today>, <last 24 hours>)
	High                       []string `json:"h"` // High array(<today>, <last 24 hours>)
	OpeningPrice               string   `json:"o"`
}

// Tickers defines a map of asset tickers.
type Tickers map[AssetPair]AssetTickerInfo

// Info returns the information about an asset pair.
func (t Tickers) Info(assetPair AssetPair) AssetTickerInfo {
	return t[assetPair]
}

// OrderBookEntries defines an order book entries.
type OrderBookEntries []any

// OrderBookEntry defines an order book entry.
type OrderBookEntry struct {
	Price     string
	Volume    string
	Timestamp int64
}

// price returns the price of the order book entries.
func (o OrderBookEntries) price() string {
	if v, ok := o[0].(string); ok {
		return v
	}
	return ""
}

// volume returns the volume of the order book entries.
func (o OrderBookEntries) volume() string {
	if v, ok := o[1].(string); ok {
		return v
	}
	return ""
}

// timestamp returns the timestamp of the order book entries.
func (o OrderBookEntries) timestamp() int64 {
	v := o[2]
	switch value := v.(type) {
	case float64:
		return int64(value)
	case int:
		return int64(value)
	default:
		return 0
	}
}

// OrderBookEntry returns the values of the order book entries.
func (o OrderBookEntries) OrderBookEntry() OrderBookEntry {
	if len(o) != 3 {
		return OrderBookEntry{}
	}

	return OrderBookEntry{
		Price:     o.price(),
		Volume:    o.volume(),
		Timestamp: o.timestamp(),
	}
}

// OrderBook defines an order book.
type OrderBook struct {
	Asks []OrderBookEntries `json:"asks"`
	Bids []OrderBookEntries `json:"bids"`
}

// TransferStatus defines the transfer status.
type TransferStatus string

const (
	Pending  TransferStatus = "pending"
	Complete TransferStatus = "complete"
)

// TransferResult defines the result of a transfer.
type TransferResult struct {
	TransferID string         `json:"transfer_id"`
	Status     TransferStatus `json:"status"`
}

// StrategyLockType .
type StrategyLockType string

const (
	Flex    StrategyLockType = "flex"
	Bounded StrategyLockType = "bounded"
	Timed   StrategyLockType = "timed"
	Instant StrategyLockType = "instant"
)

// LockType .
type LockType struct {
	Type            StrategyLockType `json:"type"`
	PayoutFrequency int              `json:"payout_frequency"`
}

// AprEstimate .
type AprEstimate struct {
	Low  string `json:"low"`
	High string `json:"high"`
}

// AutoCompound .
type AutoCompound struct {
	Default bool   `json:"default"`
	Type    string `json:"type"`
}

// YieldSource .
type YieldSource struct {
	Type string `json:"type"`
}

// Strategy represents a strategy.
type Strategy struct {
	ID                        string       `json:"id"`
	AllocationFee             any          `json:"allocation_fee"`
	AllocationRestrictionInfo []string     `json:"allocation_restriction_info"`
	AprEstimate               AprEstimate  `json:"apr_estimate"`
	Asset                     Asset        `json:"asset"`
	AutoCompound              AutoCompound `json:"auto_compound"`
	CanAllocate               bool         `json:"can_allocate"`
	CanDeallocate             bool         `json:"can_deallocate"`
	DeallocationFee           any          `json:"deallocation_fee"`
	LockType                  LockType     `json:"lock_type"`
	UserCap                   string       `json:"user_cap"`
	UserMinAllocation         string       `json:"user_min_allocation"`
	YieldSource               YieldSource  `json:"yield_source"`
}

// Strategies represents a list of strategies.
type Strategies struct {
	NextCursor string     `json:"next_cursor"`
	Items      []Strategy `json:"items"`
}

// StrategyOperationStatus .
type StrategyOperationStatus struct {
	Pending bool `json:"pending"`
}
