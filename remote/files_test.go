package remote

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPlots(t *testing.T) {

	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{src: "/home/vagrant/test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sclient, err := connectSSH("192.168.33.6", "22", "vagrant", "/Users/jimwang/.ssh/id_rsa")
			if err != nil {
				t.Fatalf(err.Error())
			}
			got, err := getPlots(sclient, tt.args.src)
			assert.NotNil(t, got)
			assert.Nil(t, err)
		})
	}
}

func TestReadHome(t *testing.T) {
	log.Println(os.UserHomeDir())
	_, err := ioutil.ReadFile("~/.ssh/id_rsa")
	assert.Nil(t, err)
}
