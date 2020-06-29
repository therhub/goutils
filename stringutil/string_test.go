package stringutil

import "testing"

func TestFormatField(t *testing.T) {
	type args struct {
		field   string
		formats []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "11", args: args{field: "1", formats: []string{"01", "02"}}, want: "`01`"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatField(tt.args.field, tt.args.formats); got != tt.want {
				t.Errorf("FormatField() = %v, want %v", got, tt.want)
			}
		})
	}
}
