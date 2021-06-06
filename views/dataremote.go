package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"plotcarrier/app"
	"plotcarrier/remote"
	"strconv"

	"github.com/spf13/viper"
)

func fetchRemotePlots(hostName, keyPath string, machineCfgs []*app.MachineCfg, plotsMap map[string]map[string]int64) [][]string {
	res := [][]string{
		{"ip", "src", "count"},
	}

	for _, mcfg := range machineCfgs {
		//Connect host
		sshClient, err := remote.ConnectSSH(mcfg.IP, "22", hostName, keyPath)
		if err != nil {
			res = append(res, []string{err.Error()})
			return res
		}

		//Fetch plots
		plots, err := remote.GetPlots(sshClient, mcfg.Src)
		if err != nil {
			res = append(res, []string{err.Error()})
			return res
		}
		plotsMap[mcfg.IP] = plots
		res = append(res, []string{mcfg.IP, mcfg.Src, strconv.FormatInt(int64(len(plots)), 10)})
	}

	return res
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
