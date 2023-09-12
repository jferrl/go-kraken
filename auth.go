package kraken

import "context"

// Secret represents a Kraken API secret.
type Secret []byte

// APIKey represents a Kraken API key.
type APIKey string

// Otp defines a single-use password.
// It is used as a second authentication factor.
type Otp string

type otpKey string

// OtpFromContext returns the one-time password from the context.
func OtpFromContext(ctx context.Context) Otp {
	otp, _ := ctx.Value(otpKey("otp")).(Otp)
	return otp
}

// ContextWithOtp adds a one-time password to the context.
func ContextWithOtp(ctx context.Context, otp Otp) context.Context {
	return context.WithValue(ctx, otpKey("otp"), otp)
}
