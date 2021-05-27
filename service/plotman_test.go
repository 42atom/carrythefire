package service

import (
	"strings"
	"testing"
)

func Test_isPlottingCmd(t *testing.T) {
	type args struct {
		cmdLine []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{strings.Split("/home/chiadmin/chia-blockchain/venv/bin/python /home/chiadmin/chia-blockchain/venv/bin/chia plots create -k 32 -r 2 -u 128 -b 3200 -t /mnt/nvme -d /mnt/ssd -f 88caca544064b891818846c7030cff003cdd0604c012734d7e6b2fe43e495924c946c98b75c1d3c9045e688e06d9c4b0 -p b75c653652dad89571ff578d49f796798299b7a4293c9e8907a6d8f33b8c33afbd81bb41bd32ef2621703af83bda3c3d -2 /mnt/ssd", " ")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPlottingCmd(tt.args.cmdLine); got != tt.want {
				t.Errorf("isPlottingCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentPlotSize(t *testing.T) {
	tests := []struct {
		name    string
		want    uint64
		wantErr bool
	}{
		{"test", 101, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CurrentPlotSize()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentPlotSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CurrentPlotSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
