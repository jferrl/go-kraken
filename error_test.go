package kraken

import "testing"

func TestError_Error(t *testing.T) {
	type fields struct {
		errors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "error", fields: fields{errors: []string{"error"}}, want: "[error]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				errors: tt.fields.errors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
