package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestWebsocketsAuth_WebsocketsToken(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WebsocketsToken
		wantErr bool
	}{
		{
			name:    "error creating request",
			fields:  fields{apiMock: createFakeServer(http.StatusOK, "")},
			args:    args{},
			wantErr: true,
		},
		{
			name:   "get websockets token",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "ws_auth.json")},
			args:   args{ctx: ctx},
			want: &WebsocketsToken{
				Token:   "1Dwc4lzSwNWOAwkMdqhssNNFhs1ed606d1WcF3XfEMw",
				Expires: 900,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client()).WithAuth(Secrets{})
			c.baseURL = baseURL

			got, err := c.WebsocketsAuth.WebsocketsToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebsocketsAuth.WebsocketsToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebsocketsAuth.WebsocketsToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
