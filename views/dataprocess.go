package views

import (
	"fmt"
	"os"
	"plotcarrier/app"
	"sort"
)

type process struct {
	ip       string
	fileName string
	percent  uint
}

func fetchProcess(plotsMap map[string]map[string]int64, machineCfgs []*app.MachineCfg) [][]string {
	res := [][]string{
		{"ip", "filename", "percent"},
	}
	localMap := map[string]string{}
	for _, cfg := range machineCfgs {
		localMap[cfg.IP] = cfg.Dst
	}
	processes := []*process{}
	for ip, plots := range plotsMap {
		for filename, totalSize := range plots {
			dst := localMap[ip]
			dstPath := fmt.Sprintf("%s/%s", dst, filename)
			finfo, err := os.Stat(dstPath)
			p := &process{}
			p.ip = ip
			p.fileName = filename
			if err != nil {
				p.fileName = fmt.Sprintf("error: %s, %s", p.fileName, err)
				p.percent = 0
			} else {
				percent := float64(finfo.Size()) / float64(totalSize)
				fmt.Println(percent)
				p.percent = uint(percent * 100)
			}

			processes = append(processes, p)
		}
	}
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].percent > processes[j].percent
	})

	for _, v := range processes {
		res = append(res, []string{
			v.ip, v.fileName, fmt.Sprintf("%v %%", v.percent),
		})
	}

	return res
}
