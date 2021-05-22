package app

import (
	"testing"
)

func TestStart(t *testing.T) {
	type args struct {
		src      string
		dst      string
		interval int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"../src", "../dst", 120}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
