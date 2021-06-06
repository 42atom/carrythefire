package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func fetchRemotePlots() []string {
	return fetchFromHttp("remote-plots")
}

func fetchFromHttp(uri string) []string {
	res := []string{}
	host := viper.GetString("status.host")
	port := viper.GetString("status.port")
	url := fmt.Sprintf("http://%s:%s/%s", host, port, uri)
	response, err := http.Get(url)
	if err != nil {
		return []string{"Not found"}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []string{"Not found"}
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return []string{"Not found"}
	}
	return res
}
