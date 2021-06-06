package views

import (
	"reflect"
	"testing"
)

func Test_fetchDisk(t *testing.T) {
	tests := []struct {
		name string
		want [][]string
	}{
		{"test", [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, machineCfgs := parseConfig()
			if got := fetchDisk(machineCfgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchDisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
