package service

import (
	"log"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var kmap = map[string]int{
	"32": 101,
	"33": 209,
	"34": 430,
	"35": 884,
}

type ChiaPlotJob struct {
	PID    int32
	JobID  string
	Kvalue string
}

func CurrentPlotSize() (uint64, error) {
	jobs, err := getPlotJobs()
	if err != nil {
		return 0, err
	}

	size := uint64(0)
	for _, v := range jobs {
		ksize := kmap[v.Kvalue]
		size += uint64(ksize)
	}
	return size, nil
}

func getPlotJobs() ([]*ChiaPlotJob, error) {
	jobs := []*ChiaPlotJob{}
	list, err := process.Processes()
	if err != nil {
		return jobs, err
	}
	for _, v := range list {
		cmd, err := v.Cmdline()
		if err != nil {
			log.Printf("Fetch process erorr: %s\n", err.Error())
			continue
		}
		cmdArr := strings.Split(cmd, " ")
		if isPlottingCmd(cmdArr) {
			job := parseCmd(cmdArr)
			job.PID = v.Pid
			jobs = append(jobs, job)
		}
	}
	return jobs, nil
}

// an example as of 1.0.5
// {
//     'size': 32,
//     'num_threads': 4,
//     'buckets': 128,
//     'buffer': 6000,
//     'tmp_dir': '/farm/yards/901',
//     'final_dir': '/farm/wagons/801',
//     'override_k': False,
//     'num': 1,
//     'alt_fingerprint': None,
//     'pool_contract_address': None,
//     'farmer_public_key': None,
//     'pool_public_key': None,
//     'tmp2_dir': None,
//     'plotid': None,
//     'memo': None,
//     'nobitfield': False,
//     'exclude_final_dir': False,
// }
func parseCmd(cmdArr []string) *ChiaPlotJob {
	cmdArr = cmdArr[4:]
	return &ChiaPlotJob{
		Kvalue: cmdArr[1],
	}
}

func isPlottingCmd(cmdArr []string) bool {
	cmd := strings.ToLower(cmdArr[0])
	pycmd := strings.Contains(cmd, "python")
	if !pycmd {
		return false
	}
	return len(cmdArr) >= 4 && strings.HasSuffix(cmdArr[1], "chia") && cmdArr[2] == "plots" && cmdArr[3] == "create"
}
