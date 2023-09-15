package kraken

import (
	"net/url"
	"testing"
)

func TestSignature_Sign(t *testing.T) {
	payload := url.Values{}
	payload.Add("pair", "XBTUSD")
	payload.Add("type", "buy")
	payload.Add("ordertype", "limit")
	payload.Add("price", "37500")
	payload.Add("volume", "1.25")
	payload.Add("nonce", "1616492376594")

	type fields struct {
		Secret string
	}
	type args struct {
		v    reqBody
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "crete a signature",
			fields: fields{
				Secret: "kQH5HW/8p1uGOVjbgWA7FunAmGO8lsSUXNsu3eow76sz84Q18fWxnyRzBHCd3pd5nE9qa99HAZtuZuj6F1huXg==",
			},
			args: args{
				path: "/0/private/AddOrder",
				v:    formURLEncodedBody{payload},
			},
			want: "4/dpxb3iT4tp/ZCVEwSnEsLxx0bqyhLpdfOpc6fn7OR8+UClSV5n9E6aSS8MPtnRfp32bAb0nmbRn6H8ndwLUQ==",
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
