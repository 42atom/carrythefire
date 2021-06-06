package service

import (
	"github.com/shirou/gopsutil/disk"
)

func DiskSize(path string) (uint64, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return 0, err
	}
	return stat.Free, nil
}

func DiskSizeGB(path string) (uint64, error) {
	size, err := DiskSize(path)
	if err != nil {
		return size, err
	}
	size = size / (1024 * 1024 * 1024)
	return size, nil
}

func DiskUsedPercent(path string) float64 {
	stat, err := disk.Usage(path)
	if err != nil {
		return 0
	}
	return float64(stat.Used) / float64(stat.Total)
}
