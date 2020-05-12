package logutil

import "testing"

func Test_write(t *testing.T) {
	type args struct {
		m  string
		lt string
		s  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{m: "sys", lt: "error", s: "this is a test for error"}, false},
		{"test2", args{m: "sys", lt: "info", s: "this is a test for info"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := write(tt.args.m, tt.args.lt, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_write_setPath(t *testing.T) {

	SetLogPath("D:\\log\\")

	type args struct {
		m  string
		lt string
		s  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", args{m: "sys", lt: "error", s: "this is a test for error"}, false},
		{"test2", args{m: "sys", lt: "info", s: "this is a test for info"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := write(tt.args.m, tt.args.lt, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
