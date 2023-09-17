package kraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestSubaccounts_Create(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts CreateSubaccountOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "invalid opts",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "create_subaccount.json")},
			args: args{
				ctx:  ctx,
				opts: CreateSubaccountOpts{Username: "test"},
			},
			wantErr: true,
		},
		{
			name:   "create subaccount",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "create_subaccount.json")},
			args: args{
				ctx:  ctx,
				opts: CreateSubaccountOpts{Username: "test", Email: "test"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Subaccounts.Create(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subaccounts.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subaccounts.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubaccounts_Transfer(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		apiMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts TransferOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransferResult
		wantErr bool
	}{
		{
			name:   "invalid opts",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "transfer.json")},
			args: args{
				ctx:  ctx,
				opts: TransferOpts{},
			},
			wantErr: true,
		},
		{
			name:   "transfer funds",
			fields: fields{apiMock: createFakeServer(http.StatusOK, "transfer.json")},
			args: args{
				ctx: ctx,
				opts: TransferOpts{
					Asset:  "XBT",
					Amount: "0.1",
					From:   "master",
					To:     "subaccount",
				},
			},
			want: &TransferResult{
				TransferID: "TOH3AS2-LPCWR8-JDQGEU",
				Status:     Complete,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.apiMock.URL + "/")

			c := New(tt.fields.apiMock.Client())
			c.baseURL = baseURL

			got, err := c.Subaccounts.Transfer(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subaccounts.Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subaccounts.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
