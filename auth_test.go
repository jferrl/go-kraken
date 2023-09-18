package kraken

import (
	"context"
	"reflect"
	"testing"
)

func TestOtpFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Otp
	}{
		{
			name: "nil context",
			args: args{ctx: nil},
			want: "",
		},
		{
			name: "empty context",
			args: args{ctx: context.Background()},
			want: "",
		},
		{
			name: "context with otp",
			args: args{ctx: ContextWithOtp(context.Background(), "123456")},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OtpFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("OtpFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContextWithOtp(t *testing.T) {
	type args struct {
		ctx context.Context
		otp Otp
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "nil context",
			args: args{otp: "123456"},
			want: context.WithValue(context.Background(), otpKey("otp"), Otp("123456")),
		},
		{
			name: "empty context",
			args: args{ctx: context.Background(), otp: "123456"},
			want: context.WithValue(context.Background(), otpKey("otp"), Otp("123456")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextWithOtp(tt.args.ctx, tt.args.otp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextWithOtp() = %v, want %v", got, tt.want)
			}
		})
	}
}
