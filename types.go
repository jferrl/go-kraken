package kraken

// ServerTime represents the server time.
type ServerTime struct {
	UnixTime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}

// ServerStatus represents the system status.
type ServerStatus string

const (
	//Online means Kraken is operating normally. All order types may be submitted and trades can occur.
	Online ServerStatus = "online"
	// Maintenance means exchange is offline. No new orders or cancellations may be submitted.
	Maintenance ServerStatus = "maintenance"
	// CancelOnly means resting (open) orders can be cancelled but no new orders may be submitted. No trades will occur.
	CancelOnly ServerStatus = "cancel_only"
	// PostOnly means only post-only limit orders can be submitted. Existing orders may still be cancelled. No trades will occur.
	PostOnly ServerStatus = "post_only"
)

// SystemStatus represents the current system status.
type SystemStatus struct {
	Status    ServerStatus `json:"status"`
	Timestamp string       `json:"timestamp"` // Current timestamp (RFC3339)
}

// AccountBalance represents the user's account balance.
type AccountBalance struct {
	ADA   float64 `json:"ADA,string"`
	AAVE  float64 `json:"AAVE,string"`
	BCH   float64 `json:"BCH,string"`
	DASH  float64 `json:"DASH,string"`
	EOS   float64 `json:"EOS,string"`
	GNO   float64 `json:"GNO,string"`
	QTUM  float64 `json:"QTUM,string"`
	KFEE  float64 `json:"KFEE,string"`
	LINK  float64 `json:"LINK,string"`
	USDC  float64 `json:"USDC,string"`
	USDT  float64 `json:"USDT,string"`
	XDAO  float64 `json:"XDAO,string"`
	XETC  float64 `json:"XETC,string"`
	XETH  float64 `json:"XETH,string"`
	XICN  float64 `json:"XICN,string"`
	XLTC  float64 `json:"XLTC,string"`
	XMLN  float64 `json:"XMLN,string"`
	XNMC  float64 `json:"XNMC,string"`
	XREP  float64 `json:"XREP,string"`
	XXBT  float64 `json:"XXBT,string"`
	XXDG  float64 `json:"XXDG,string"`
	XXLM  float64 `json:"XXLM,string"`
	XXMR  float64 `json:"XXMR,string"`
	XXRP  float64 `json:"XXRP,string"`
	XTZ   float64 `json:"XTZ,string"`
	XXVN  float64 `json:"XXVN,string"`
	XZEC  float64 `json:"XZEC,string"`
	ZCAD  float64 `json:"ZCAD,string"`
	ZEUR  float64 `json:"ZEUR,string"`
	ZGBP  float64 `json:"ZGBP,string"`
	ZJPY  float64 `json:"ZJPY,string"`
	ZKRW  float64 `json:"ZKRW,string"`
	ZUSD  float64 `json:"ZUSD,string"`
	TRX   float64 `json:"TRX,string"`
	DAI   float64 `json:"DAI,string"`
	DOT   float64 `json:"DOT,string"`
	ETH2S float64 `json:"ETH2.S,string"`
	ETH2  float64 `json:"ETH2,string"`
	USDM  float64 `json:"USD.M,string"`
}

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

// AddOrderResponse defines the response from the AddOrder method.
type AddOrderResponse struct {
	Description OrderDescription `json:"descr"`
	Transaction []TransactionID  `json:"txid"`
}
