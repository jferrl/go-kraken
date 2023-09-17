package kraken

import (
	"reflect"
	"testing"
)

func TestTickerValues_Ticker(t *testing.T) {
	tests := []struct {
		name       string
		tr         TickerValues
		wantTicker Ticker
	}{
		{
			name: "invalid ticker",
			tr:   TickerValues{},
		},
		{
			name: "tick with invalid values",
			tr:   TickerValues{"0", 1, 2, 3, 4, 5, 6, "7"},
		},
		{
			name: "tick with valid values",
			tr: TickerValues{
				1688671200,
				"30306.1",
				"30306.2",
				"30305.7",
				"30305.7",
				"30306.1",
				"3.39243896",
				23,
			},
			wantTicker: Ticker{
				Time:   1688671200,
				Open:   "30306.1",
				High:   "30306.2",
				Low:    "30305.7",
				Close:  "30305.7",
				Vwap:   "30306.1",
				Volume: "3.39243896",
				Count:  23,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Ticker(); !reflect.DeepEqual(got, tt.wantTicker) {
				t.Errorf("TickerValues.Ticker() = %v, want %v", got, tt.wantTicker)
			}
		})
	}
}

func TestOrderBookEntries_OrderBookEntry(t *testing.T) {
	tests := []struct {
		name string
		o    OrderBookEntries
		want OrderBookEntry
	}{
		{
			name: "invalid order book entry",
		},
		{
			name: "order book entry with invalid values",
			o: OrderBookEntries{
				1, 2, "fake",
			},
		},
		{
			name: "order book entry with valid values",
			o: OrderBookEntries{
				"30306.1", "30306.2", 3242,
			},
			want: OrderBookEntry{
				Price:     "30306.1",
				Volume:    "30306.2",
				Timestamp: 3242,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OrderBookEntry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBookEntries.OrderBookEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_Float64(t *testing.T) {
	tests := []struct {
		name string
		b    Balance
		want float64
	}{
		{
			name: "invalid balance",
			b:    "fake",
		},
		{
			name: "valid balance",
			b:    "1.2345678901",
			want: 1.2345678901,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Float64(); got != tt.want {
				t.Errorf("Balance.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
