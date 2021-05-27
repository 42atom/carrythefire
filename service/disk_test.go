package service

import (
	"testing"
)

func TestDiskSize(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"../dst"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DiskSize(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= 0 {
				t.Errorf("DiskSize() = %v", got)
			}
		})
	}
}

func TestDiskSizeGB(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"../dst"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DiskSizeGB(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskSizeGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= 0 {
				t.Errorf("DiskSize() = %v", got)
			}
		})
	}
}
