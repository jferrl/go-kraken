package kraken

import (
	"net/url"
	"testing"
)

func Test_newFormURLEncodedBody(t *testing.T) {
	type fakeJSON struct {
		Asset  string `json:"asset"`
		Cursor string `json:"cursor"`
		Limit  int    `json:"limit"`
	}

	type args struct {
		b url.Values
	}
	tests := []struct {
		name            string
		args            args
		wantString      string
		wantNonce       string
		wantContentType string
		wantValues      string
		wantErr         bool
	}{
		{
			name: "new json body as reference",
			args: args{
				b: url.Values{
					"asset":  []string{"XBT"},
					"cursor": []string{"cursor"},
					"limit":  []string{"10"},
				},
			},
			wantNonce:       "nonce",
			wantContentType: "application/x-www-form-urlencoded; charset=utf-8",
			wantValues:      "asset=XBT&cursor=cursor&limit=10&nonce=nonce&otp=otp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newFormURLEncodedBody(tt.args.b)

			got.Values.Set("nonce", "nonce")
			got.withOtp("otp")

			if got.nonce() != tt.wantNonce {
				t.Errorf("newJSONBody() nonce = %v, want %v", got.nonce(), tt.wantNonce)
				return
			}

			if got.contentType() != tt.wantContentType {
				t.Errorf("newJSONBody() contentType = %v, want %v", got.contentType(), tt.wantContentType)
				return
			}

			if got.string() != tt.wantValues {
				t.Errorf("newJSONBody() string = %v, want %v", got.string(), tt.wantValues)
			}
		})
	}
}

func Test_newJSONBody(t *testing.T) {
	type fakeJSON struct {
		Asset  string `json:"asset"`
		Cursor string `json:"cursor"`
		Limit  int    `json:"limit"`
	}

	type args struct {
		b any
	}
	tests := []struct {
		name            string
		args            args
		wantString      string
		wantNonce       string
		wantContentType string
		wantJSON        string
		wantErr         bool
	}{
		{
			name: "new json body as pointer",
			args: args{
				b: &fakeJSON{
					Asset:  "XBT",
					Cursor: "cursor",
					Limit:  10,
				},
			},
			wantNonce:       "nonce",
			wantContentType: "application/json",
			wantJSON:        `{"asset":"XBT","cursor":"cursor","limit":10,"nonce":"nonce","otp":"otp"}`,
		},
		{
			name: "new json body as reference",
			args: args{
				b: fakeJSON{
					Asset:  "XBT",
					Cursor: "cursor",
					Limit:  10,
				},
			},
			wantNonce:       "nonce",
			wantContentType: "application/json",
			wantJSON:        `{"asset":"XBT","cursor":"cursor","limit":10,"nonce":"nonce","otp":"otp"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newJSONBody(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("newJSONBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Add fake nonce to the got json body.
			got.jsonMessage["nonce"] = "nonce"
			got.withOtp("otp")

			if got.nonce() != tt.wantNonce {
				t.Errorf("newJSONBody() nonce = %v, want %v", got.nonce(), tt.wantNonce)
				return
			}

			if got.contentType() != tt.wantContentType {
				t.Errorf("newJSONBody() contentType = %v, want %v", got.contentType(), tt.wantContentType)
				return
			}

			if got.string() != tt.wantJSON {
				t.Errorf("newJSONBody() string = %v, want %v", got.string(), tt.wantJSON)
			}
		})
	}
}
