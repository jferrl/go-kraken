package kraken

import (
	"context"
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
