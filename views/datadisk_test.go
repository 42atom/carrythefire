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
			_, _, targets := parseConfig()
			if got := fetchDisk(targets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchDisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
