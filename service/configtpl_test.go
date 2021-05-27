package service

import (
	_ "embed"
	"testing"
)

func TestInitConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"config.yaml"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
