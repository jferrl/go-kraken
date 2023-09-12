package kraken

import (
	"reflect"
	"testing"
)

func TestTick_Values(t *testing.T) {
	tests := []struct {
		name     string
		tr       Tick
		wantTick TickValues
	}{
		{
			name: "invalid tick",
			tr:   Tick{},
		},
		{
			name: "tick with invalid values",
			tr:   Tick{"0", 1, 2, 3, 4, 5, 6, "7"},
		},
		{
			name: "tick with valid values",
			tr: Tick{
				1688671200,
				"30306.1",
				"30306.2",
				"30305.7",
				"30305.7",
				"30306.1",
				"3.39243896",
				23,
			},
			wantTick: TickValues{
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
			if got := tt.tr.Values(); !reflect.DeepEqual(got, tt.wantTick) {
				t.Errorf("tick = %v, want %v", got, tt.wantTick)
			}
		})
	}
}
