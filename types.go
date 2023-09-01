package kraken

// ServerTime represents the server time.
type ServerTime struct {
	UnixTime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}

// SystemStatus represents the system status.
type SystemStatus string

const (
	//Online means Kraken is operating normally. All order types may be submitted and trades can occur.
	Online SystemStatus = "online"
	// Maintenance means exchange is offline. No new orders or cancellations may be submitted.
	Maintenance SystemStatus = "maintenance"
	// CancelOnly means resting (open) orders can be cancelled but no new orders may be submitted. No trades will occur.
	CancelOnly SystemStatus = "cancel_only"
	// PostOnly means only post-only limit orders can be submitted. Existing orders may still be cancelled. No trades will occur.
	PostOnly SystemStatus = "post_only"
)

// CurrentSystemStatus represents the current system status.
type CurrentSystemStatus struct {
	Status    SystemStatus `json:"status"`
	Timestamp string       `json:"timestamp"` // Current timestamp (RFC3339)
}
