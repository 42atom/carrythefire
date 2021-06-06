package views

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../config.yaml")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Test_fetchRemotePlots(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"test", []string{"hello world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hostName, keyPath, machineCfgs := parseConfig()
			if got := fetchRemotePlots(hostName, keyPath, machineCfgs, map[string]map[string]int64{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchRemotePlots() = %v, want %v", got, tt.want)
			}
		})
	}
}
