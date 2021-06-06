package views

import (
	"fmt"
	"testing"
	"time"
)

func TestView_Update(t *testing.T) {
	hostName, keyPath, machineCfgs := parseConfig()
	plotsMap := map[string]map[string]int64{}

	//Fetch remote plots
	plotsData := fetchRemotePlots(hostName, keyPath, machineCfgs, plotsMap)
	//Fetch disk usage
	DiskData := fetchDisk(machineCfgs)
	processData := fetchProcess(plotsMap, machineCfgs)

	fmt.Println(plotsData)
	fmt.Println(DiskData)
	fmt.Println(processData)

	//remoteUpdateInterval := 3 * time.Minute
	remoteUpdateInterval := 10 * time.Second
	go func() {
		for range time.NewTicker(remoteUpdateInterval).C {
			//Fetch remote plots
			plotsData = fetchRemotePlots(hostName, keyPath, machineCfgs, plotsMap)
			//Fetch disk usage
			DiskData = fetchDisk(machineCfgs)
		}
	}()

	pInterval := 1 * time.Second
	go func() {
		for range time.NewTicker(pInterval).C {
			//Fetch process
			processData = fetchProcess(plotsMap, machineCfgs)
			fmt.Println(processData)
		}
	}()
	for {
		time.Sleep(5 * time.Second)
	}
}
