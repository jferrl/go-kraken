package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestTrading_AddOrder(t *testing.T) {
	ctx := ContextWithOtp(context.TODO(), "123456")

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
		want    *OrderCreation
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
			want: &OrderCreation{
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

			c := New(tt.fields.apiMock.Client())
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

func TestTrading_CancelOrder(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts CancelOrderOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderCancelation
		wantErr bool
	}{
		{
			name:   "cancel order",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "cancel_order.json")},
			args:   args{ctx: ctx, opts: CancelOrderOpts{TransactionID: "OUF4EM-FRGI2-MQMWZD"}},
			want:   &OrderCancelation{Count: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Trading.CancelOrder(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trading.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trading.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrading_CancelAllOrders(t *testing.T) {
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
		want    *OrderCancelation
		wantErr bool
	}{
		{
			name:   "cancel orders",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "cancel_order.json")},
			args:   args{ctx: ctx},
			want:   &OrderCancelation{Count: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Trading.CancelAllOrders(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trading.CancelAllOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trading.CancelAllOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrading_CancelAllOrdersAfter(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts CancelAllOrdersAfterOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TriggeredOrderCancellation
		wantErr bool
	}{
		{
			name:   "cancel all orders after 60 seconds",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "cancel_all_orders_after.json")},
			args: args{ctx: ctx, opts: CancelAllOrdersAfterOpts{
				Timeout: time.Minute,
			},
			},
			want: &TriggeredOrderCancellation{
				CurrentTime: "2023-03-24T17:41:56Z",
				TriggerTime: "2023-03-24T17:42:56Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Trading.CancelAllOrdersAfter(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Trading.CancelAllOrdersAfter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trading.CancelAllOrdersAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
