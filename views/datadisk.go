package views

import (
	"fmt"
	"plotcarrier/app"
	"plotcarrier/service"
)

func fetchDisk(targets []*app.Target) [][]string {
	res := [][]string{{
		"Disk", "Used", "Free",
	}}
	machineCfgs := []*app.MachineCfg{}
	for _, v := range targets {
		machineCfgs = append(machineCfgs, v.MachineCfgs...)
	}
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
