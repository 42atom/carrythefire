package app

import (
	"fmt"
	"log"
	"plotcarrier/remote"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Target struct {
	BindAddress string        `yaml:"bindAddress"`
	MachineCfgs []*MachineCfg `mapstructure:"machines"`
}

type MachineCfg struct {
	IP  string `yaml:"ip"`
	Src string `yaml:"remote_src"`
	Dst string `yaml:"local_dst"`
}

func RemoteStart() error {
	hostName := viper.GetString("host.username")
	keyPath := viper.GetString("host.keypath")
	if hostName == "" || keyPath == "" {
		return fmt.Errorf("HostName or ssh priavate key path is empty")
	}
	targets := []*Target{}
	err := viper.UnmarshalKey("targets", &targets)
	if err != nil {
		return err
	}
	interval := viper.GetInt("interval")

	var wg sync.WaitGroup
	carrierWorker := viper.GetInt("worker")
	for _, v := range targets {
		wg.Add(1)
		go distributeByBindAddress(v.MachineCfgs, v.BindAddress, hostName, keyPath, interval, carrierWorker)
	}
	wg.Wait()

	return nil
}

func distributeByBindAddress(cfgs []*MachineCfg, bindAddress, hostName, keyPath string, interval, carrierWorker int) {
	//Maximum 8 workers for each bindAddress
	if carrierWorker > 8 {
		carrierWorker = 8
	}

	count := 0
	mtotal := len(cfgs)
	machine := make(chan *MachineCfg, carrierWorker)

	//Start worker
	workerMap := map[string]bool{}
	mutex := &sync.RWMutex{}
	for i := 0; i < carrierWorker; i++ {
		go worker(i, bindAddress, hostName, keyPath, interval, machine, carrierWorker, workerMap, mutex)
	}

	for {
		if count > mtotal-1 {
			count = 0
		}
		machine <- cfgs[count]
		count++
	}
}

func worker(id int, bindAddress, hostname, keypath string, interval int, machine <-chan *MachineCfg, carrierWorker int, workerMap map[string]bool, mutex *sync.RWMutex) {
	log.Printf("BindAddress %s, start worker: %d\n", bindAddress, id)
	for m := range machine {
		if v, ok := workerMap[m.IP]; ok && v {
			log.Printf("%s_%d, already exists, skip ip: %s, src: %s, dst: %s\n", bindAddress, id, m.IP, m.Src, m.Dst)
			time.Sleep(time.Duration(interval) * time.Second)
			continue
		}
		log.Printf("%s_%d, start job. ip: %s, src: %s, dst: %s\n", bindAddress, id, m.IP, m.Src, m.Dst)
		mutex.RLock()
		workerMap[m.IP] = true
		mutex.RUnlock()
		// r := rand.Intn(20)
		// time.Sleep(time.Duration(r) * time.Second)
		err := remote.StartSCPSimple(m.IP, bindAddress, m.Src, m.Dst, hostname, keypath, carrierWorker)
		if err != nil {
			log.Printf("%s_%d, Move file error: %s", bindAddress, id, err)
		}
		log.Printf("%s_%d, finish job. ip: %s, src: %s, dst: %s, sleep %d second\n", bindAddress, id, m.IP, m.Src, m.Dst, interval)
		time.Sleep(time.Duration(interval) * time.Second)
		mutex.Lock()
		delete(workerMap, m.IP)
		mutex.Unlock()
	}
}
