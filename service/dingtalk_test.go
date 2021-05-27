package service

import (
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("dingtalk.url", "")
}

func TestDingTalkClient_SendMessage(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewDingTalk().SendMessage(DingTalkMessage{Type: "text", Message: "I'm crazy!"})
			if (err != nil) != tt.wantErr {
				t.Errorf("DingTalkClient.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
