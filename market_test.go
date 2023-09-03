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

			got, err := c.Market.Time(tt.args.ctx)
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

			got, err := c.Market.SystemStatus(tt.args.ctx)
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

func TestMarketData_Assets(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts AssetsOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		asset   Asset
		want    AssetInfo
		wantErr bool
	}{
		{
			name: "get assets",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "asset_info.json"),
			},
			args: args{
				ctx: ctx,
				opts: AssetsOpts{
					Class: Currency,
				},
			},
			asset: ETH2,
			want: AssetInfo{
				Altname:         "ETH2",
				AssetClass:      "currency",
				Decimals:        10,
				DisplayDecimals: 5,
				Status:          "enabled",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Market.Assets(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.Assets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			a := got.Info(tt.asset)

			if !reflect.DeepEqual(a, tt.want) {
				t.Errorf("MarketData.Assets() = %v, want %v", a, tt.want)
			}
		})
	}
}

func Test_TradableAssetPairs(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts TradableAssetPairsOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pair    AssetPair
		want    AssetPairInfo
		wantErr bool
	}{
		{
			name: "get tradable asset pairs",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "asset_pairs.json"),
			},
			args: args{
				ctx: ctx,
				opts: TradableAssetPairsOpts{
					Info: Info,
				},
			},
			pair: ETHUSDC,
			want: AssetPairInfo{
				Altname:            "ETHUSDC",
				WSName:             "ETH/USDC",
				AClassBase:         "currency",
				Base:               "XETH",
				AClassQuote:        "currency",
				Quote:              "USDC",
				PairDecimals:       2,
				CostDecimals:       6,
				LotDecimals:        8,
				LotMultiplier:      1,
				LeverageBuy:        []int{2, 3, 4},
				LeverageSell:       []int{2, 3, 4},
				Fees:               []FeeTuple{{0, 0.26}},
				FeesMaker:          []FeeTuple{{0, 0.16}},
				FeeVolumeCurrency:  "ZUSD",
				MarginCall:         80,
				MarginStop:         40,
				OrderMin:           "0.01",
				CostMin:            "0.5",
				TickSize:           "0.01",
				Status:             "online",
				LongPositionLimit:  170,
				ShortPositionLimit: 125,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Market.TradableAssetPairs(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.TradableAssetPairs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			a := got.Info(tt.pair)

			if !reflect.DeepEqual(a, tt.want) {
				t.Errorf("MarketData.TradableAssetPairs() = %v, want %v", a, tt.want)
			}
		})
	}
}
