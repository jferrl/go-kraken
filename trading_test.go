package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestTrading_AddOrder(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts AddOrderOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AddOrderResponse
		wantErr bool
	}{
		{
			name: "add order",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "add_order.json"),
			},
			args: args{
				ctx: ctx,
				opts: AddOrderOpts{
					UserRef:        "test",
					OrderType:      Limit,
					Type:           Buy,
					Volume:         "2.1234",
					Pair:           "XXBTZUSD",
					Price:          "1000",
					Leverage:       "2:1",
					CloseOrderType: StopLossLimit,
					ClosePrice:     "900",
					ClosePrice2:    "800",
				},
			},
			want: &AddOrderResponse{
				Description: OrderDescription{
					Order: "buy 2.12340000 XBTUSD @ limit 25000.1 with 2:1 leverage",
					Close: "close position @ stop loss 22000.0 -> limit 21000.0",
				},
				Transaction: []TransactionID{
					"OUF4EM-FRGI2-MQMWZD",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client()).WithAuth(Secrets{})
			c.baseURL = baseURL

			got, err := c.Trading.AddOrder(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trading.AddOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trading.AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
