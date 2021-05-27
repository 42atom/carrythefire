package service

import (
	_ "embed"
	"io/ioutil"
)

//go:embed templates/config.yaml
var defaultConfig []byte

func InitConfig(filename string) error {
	return ioutil.WriteFile(filename, defaultConfig, 0644)
}
