package views

import (
	"fmt"
	"plotcarrier/app"
	"plotcarrier/service"
)

func fetchDisk(machineCfgs []*app.MachineCfg) [][]string {
	res := [][]string{{
		"Disk", "Used", "Free",
	}}
	for _, cfg := range machineCfgs {
		size, err := service.DiskSizeGB(cfg.Dst)
		percent := uint(service.DiskUsedPercent(cfg.Dst) * 100)
		if err != nil {
			res = append(res, []string{
				fmt.Sprintf("%s:%s", cfg.Dst, err.Error()), "error", "error",
			})
		} else {
			res = append(res, []string{
				cfg.Dst, fmt.Sprintf("%v %%", percent), fmt.Sprintf("%v GB", size),
			})
		}
	}
	return res
}
