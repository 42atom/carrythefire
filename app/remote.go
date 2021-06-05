package app

import (
	"fmt"
	"log"
	"plotcarrier/remote"
	"time"

	"github.com/spf13/viper"
)

type MachineCfg struct {
	IP          string `yaml:"ip"`
	BindAddress string `yaml:"bindAddress"`
	Src         string `yaml:"remote_src"`
	Dst         string `yaml:"local_dst"`
}

func RemoteStart() error {
	hostName := viper.GetString("host.username")
	keyPath := viper.GetString("host.keypath")
	if hostName == "" || keyPath == "" {
		return fmt.Errorf("HostName or ssh priavate key path is empty")
	}
	machineCfgs := []*MachineCfg{}
	err := viper.UnmarshalKey("machines", &machineCfgs)
	if err != nil {
		return err
	}
	interval := viper.GetInt("interval")
	count := 0
	mtotal := len(machineCfgs)
	workerNum := viper.GetInt("worker")
	machine := make(chan *MachineCfg, workerNum)

	//Start worker
	for i := 0; i < workerNum; i++ {
		go worker(i, hostName, keyPath, machine)
	}

	for {
		if count > mtotal-1 {
			log.Printf("Already finish a round, sleep %d second", interval)
			time.Sleep(time.Duration(interval) * time.Second)
			count = 0
		}
		machine <- machineCfgs[count]
		count++
	}
}

func worker(id int, hostname, keypath string, machine <-chan *MachineCfg) {
	log.Printf("Start worker_%d\n", id)
	for m := range machine {
		log.Printf("Worker_%d, start job. ip: %s, src: %s, dst: %s\n", id, m.IP, m.Src, m.Dst)
		//err := remote.StartSCP(m.IP, m.Src, m.Dst, hostname, keypath)
		err := remote.StartSCPSimple(m.IP, m.BindAddress, m.Src, m.Dst, hostname, keypath)
		if err != nil {
			log.Printf("Worker_%d, Move file error: %s", id, err)
		}
		log.Printf("Worker_%d, finish job. ip: %s, src: %s, dst: %s\n", id, m.IP, m.Src, m.Dst)
	}
	log.Printf("End worker_%d\n", id)
}
