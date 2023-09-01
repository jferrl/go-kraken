package kraken

import (
	"net/url"
	"testing"
)

func TestSignature_Sign(t *testing.T) {
	type fields struct {
		Secret string
	}
	type args struct {
		v    url.Values
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "crete a signature",
			fields: fields{Secret: "SECRET"},
			args: args{
				path: "/0/private/",
				v: url.Values{
					"TestKey": {"TestValue"},
				},
			},
			want: "Uog0MyIKZmXZ4/VFOh0g1u2U+A0ohuK8oCh0HFUiHLE2Csm23CuPCDaPquh/hpnAg/pSQLeXyBELpJejgOftCQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSigner(tt.fields.Secret)
			if got := s.Sign(tt.args.v, tt.args.path); got != tt.want {
				t.Errorf("Signature.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
