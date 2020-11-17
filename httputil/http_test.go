package httputil

import (
	"io"
	"reflect"
	"testing"

	"golang.org/x/text/encoding"
)

func TestGetEncoding(t *testing.T) {
	type args struct {
		r io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want encoding.Encoding
	}{
		// {name: "test1", args: &io.ReadCloser{}}, {want: encoding.ASCIISub},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEncoding(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
