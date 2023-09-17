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
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
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
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
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
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
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
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
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

func TestMarketData_OHCLData(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts OHCLDataOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OHCL
		wantErr bool
	}{
		{
			name: "pair is required",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "ohcl_data.json"),
			},
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
		{
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
				opts: OHCLDataOpts{
					Pair: XXBTZUSD,
				},
			},
			wantErr: true,
		},
		{
			name: "get OHCL data",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "ohcl_data.json"),
			},
			args: args{
				ctx: ctx,
				opts: OHCLDataOpts{
					Pair: XXBTZUSD,
				},
			},
			want: &OHCL{
				Last: 1688672160,
				Pair: OHCLTickers{
					{
						float64(1688671200),
						"30306.1",
						"30306.2",
						"30305.7",
						"30305.7",
						"30306.1",
						"3.39243896",
						float64(23),
					},
					{
						float64(1688671260),
						"30304.5",
						"30304.5",
						"30300.0",
						"30300.0",
						"30300.0",
						"4.42996871",
						float64(18),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Market.OHCLData(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.OHCLData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketData.OHCLData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarketData_TickerInformation(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts TickerInformationOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pair    AssetPair
		want    AssetTickerInfo
		wantErr bool
	}{
		{
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
			},
			wantErr: true,
		},
		{
			name: "get ticker information",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "ticker.json"),
			},
			args: args{
				ctx: ctx,
				opts: TickerInformationOpts{
					Pairs: []AssetPair{XXBTZUSD},
				},
			},
			pair: XXBTZUSD,
			want: AssetTickerInfo{
				Ask:                        []string{"30300.10000", "1", "1.000"},
				Bid:                        []string{"30300.00000", "1", "1.000"},
				Last:                       []string{"30303.20000", "0.00067643"},
				Volume:                     []string{"4083.67001100", "4412.73601799"},
				VolumeWeightedAveragePrice: []string{"30706.77771", "30689.13205"},
				NumberOfTrades:             []int{34619, 38907},
				Low:                        []string{"29868.30000", "29868.30000"},
				High:                       []string{"31631.00000", "31631.00000"},
				OpeningPrice:               "30502.80000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Market.TickerInformation(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.TickerInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			a := got.Info(tt.pair)

			if !reflect.DeepEqual(a, tt.want) {
				t.Errorf("MarketData.TickerInformation() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestMarketData_OrderBook(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts OrderBookOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderBook
		wantErr bool
	}{
		{
			name: "pair is required",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "order_book.json"),
			},
			args: args{
				ctx:  ctx,
				opts: OrderBookOpts{},
			},
			wantErr: true,
		},
		{
			name: "error building request",
			fields: fields{
				apiMock: createFakeServer(http.StatusFound, ""),
			},
			args: args{
				ctx: nil,
				opts: OrderBookOpts{
					Pair: XXBTZUSD,
				},
			},
			wantErr: true,
		},
		{
			name: "count must be between 1 and 500",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "order_book.json"),
			},
			args: args{
				ctx: ctx,
				opts: OrderBookOpts{
					Pair:  XXBTZUSD,
					Count: 501,
				},
			},
			wantErr: true,
		},
		{
			name: "get order book",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "order_book.json"),
			},
			args: args{
				ctx: ctx,
				opts: OrderBookOpts{
					Pair: XXBTZUSD,
				},
			},
			want: &OrderBook{
				Asks: []OrderBookEntries{
					{
						"30384.10000",
						"2.059",
						float64(1688671659),
					},
					{
						"30387.90000",
						"1.500",
						float64(1688671380),
					},
				},
				Bids: []OrderBookEntries{
					{
						"30297.00000",
						"1.115",
						float64(1688671636),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Market.OrderBook(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarketData.OrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarketData.OrderBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
