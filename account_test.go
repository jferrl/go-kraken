package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestAccount_Balance(t *testing.T) {
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
		want    AccountBalance
		wantErr bool
	}{
		{
			name: "error creating request",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, ""),
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "get account balance",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "account_balance.json"),
			},
			args: args{
				ctx: ctx,
			},
			want: AccountBalance{
				ZUSD: "171288.6158",
				XXBT: "0.0000000000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Account.Balance(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_ExtendedBalance(t *testing.T) {
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
		want    AccountExtendedBalance
		wantErr bool
	}{
		{
			name: "error creating request",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, ""),
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "get extended account balance",
			fields: fields{
				apiMock: createFakeServer(http.StatusOK, "extended_balance.json"),
			},
			args: args{
				ctx: ctx,
			},
			want: AccountExtendedBalance{
				ZUSD: {
					Balance:   "25435.21",
					HoldTrade: "8249.76",
				},
				XXBT: {
					Balance:   "1.2435",
					HoldTrade: "0.8423",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Account.ExtendedBalance(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.ExtendedBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.ExtendedBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
