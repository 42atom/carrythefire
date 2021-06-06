package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchRemotePlots() []string {
	return fetchFromHttp("remote-plots")
}

func fetchFromHttp(uri string) []string {
	res := []string{}
	url := fmt.Sprintf("%s/%s", "http://localhost:5678", uri)
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
