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
