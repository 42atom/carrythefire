package service

import (
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("email.host", "smtp.test.com")
	viper.SetDefault("email.port", 547)
	viper.SetDefault("email.username", "username")
	viper.SetDefault("email.password", "password")
	viper.SetDefault("email.to", []string{"abc@abc.com", "cde@abc.com"})
}

func TestSMTPMailTo(t *testing.T) {
	type args struct {
		subject string
		msg     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{"Alert", "message"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SMTPMailTo(tt.args.subject, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SMTPMailTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
