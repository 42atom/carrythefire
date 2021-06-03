package remote

import (
	"testing"
)

func TestStartSCP(t *testing.T) {
	type args struct {
		ip           string
		src          string
		dst          string
		hostUsername string
		hostKeypath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{
			ip:           "192.168.33.6",
			src:          "/home/vagrant/test",
			dst:          "../dst",
			hostUsername: "vagrant",
			hostKeypath:  "/Users/jimwang/.ssh/id_rsa",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartSCP(tt.args.ip, tt.args.src, tt.args.dst, tt.args.hostUsername, tt.args.hostKeypath); (err != nil) != tt.wantErr {
				t.Errorf("StartSCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
