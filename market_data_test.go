package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestMarketData_Time(t *testing.T) {
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
		want    *ServerTime
		wantErr bool
	}{
		{
			name: "error getting server time",
			fields: fields{
				apiMock: createFakeServer(http.StatusBadRequest, "error_response.json"),
			},
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
		{
			name: "get server time",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "server_time.json"),
			},
			args: args{
				ctx: ctx,
			},
			want: &ServerTime{
				UnixTime: 1688669448,
				Rfc1123:  "Thu, 06 Jul 23 18:50:48 +0000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.MarketData.Time(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketData.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketData_SystemStatus(t *testing.T) {
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
		want    *SystemStatus
		wantErr bool
	}{
		{
			name: "get server time",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "system_status.json"),
			},
			args: args{
				ctx: ctx,
			},
			want: &SystemStatus{
				Status:    "online",
				Timestamp: "2023-07-06T18:52:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.MarketData.SystemStatus(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.SystemStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketData.SystemStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
